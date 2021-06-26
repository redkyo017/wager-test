package handlers

import (
	"net/http"
	"repositories"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var wagerRepository *repositories.WagerReposioty
var buyRepository *repositories.BuyReposioty

type WagerDto struct {
	Id                  uint      `json:"id"`
	TotalWagerValue     uint      `json:"total_wager_value"`
	Odds                uint      `json:"odds"`
	SellingPercentage   int       `json:"selling_percentage"`
	SellingPrice        float64   `json:"selling_price"`
	CurrentSellingPrice float64   `json:"current_selling_price"`
	PercentageSold      int       `json:"percentage_sold"`
	AmountSold          int       `json:"amount_sold"`
	PlacedAt            time.Time `json:"placed_at"`
}

type CreateWagerDto struct {
	TotalWagerValue   uint    `json:"total_wager_value" validate:"gt=0"`
	Odds              uint    `json:"odds" validate:"gt=0"`
	SellingPercentage int     `json:"selling_percentage" validate:"gt=0,lte=100"`
	SellingPrice      float64 `json:"selling_price" validate:"gt=0"`
}

type BuyWagerResponseDto struct {
	Id          uint      `json:"id"`
	WagerId     uint      `json:"wager_id"`
	BuyingPrice float64   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}

type BuyWagerDto struct {
	BuyingPrice int `json:"buying_price"`
}

func GetWagers(c *gin.Context) {
	pageQuery := c.DefaultQuery("page", "1")
	limitQuery := c.DefaultQuery("limit", "50")
	page, err := strconv.Atoi(pageQuery)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, GetErrorResponse("page query params invalid"))
	}
	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, GetErrorResponse("page query params invalid"))
	}
	wagers := []WagerDto{}
	result := wagerRepository.FindAll(page, limit)
	for _, wager := range result {
		wagers = append(wagers, PrepareWagerDTO(wager))
	}
	c.JSON(http.StatusOK, wagers)
}

func CreateWager(c *gin.Context) {
	var payload CreateWagerDto
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, GetErrorResponse("invalid create wager playload"))
		return
	}
	err = ValidateCreateWagerDTO(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, GetErrorResponse(err.Error()))
		return
	}
	wager := PrepareWagerModel(payload)

	result := wagerRepository.Save(wager)

	c.JSON(http.StatusOK, PrepareWagerDTO(result))
}

func BuyWager(c *gin.Context) {
	wagerId, err := strconv.Atoi(c.Param("wager_id"))
	if err != nil || wagerId <= 0 {
		c.JSON(http.StatusBadRequest, GetErrorResponse("invalid wager id"))
		return
	}
	var buyWagerPayload BuyWagerDto
	err = c.ShouldBindJSON(&buyWagerPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, GetErrorResponse("invalid buy wager payload"))
		return
	}
	wager := wagerRepository.FindByID(uint(wagerId))
	if wager.ID <= 0 {
		c.JSON(http.StatusNotFound, GetErrorResponse("wager not found"))
		return
	}
	if buyWagerPayload.BuyingPrice > int(wager.CurrentSellingPrice) {
		c.JSON(http.StatusBadRequest, GetErrorResponse("buying price over current selling price"))
		return
	}
	buy := PrepareBuyModel(uint(wagerId), float64(buyWagerPayload.BuyingPrice))
	buyResult := buyRepository.Save(buy)
	wager.AmountSold += buyWagerPayload.BuyingPrice
	wager.CurrentSellingPrice -= float64(buyWagerPayload.BuyingPrice)
	wager.PercentageSold = int((float64(wager.AmountSold) / float64(wager.TotalWagerValue)) * 100)
	wager = wagerRepository.Save(wager)
	c.JSON(http.StatusCreated, PrepareBuyDTO(buyResult))
}

func InitRepo(DB *gorm.DB) {
	wagerRepoProvider := repositories.ProvideWagerReposioty(DB)
	wagerRepository = &wagerRepoProvider
	buyRepoProvider := repositories.ProvideBuyReposioty(DB)
	buyRepository = &buyRepoProvider
}
