package handlers

import (
	"errors"
	"repositories"

	"github.com/gin-gonic/gin"
)

func PrepareWagerModel(wagerDto CreateWagerDto) repositories.Wager {
	wager := repositories.Wager{}
	wager.TotalWagerValue = wagerDto.TotalWagerValue
	wager.Odds = wagerDto.Odds
	wager.SellingPercentage = wagerDto.SellingPercentage
	wager.SellingPrice = wagerDto.SellingPrice
	wager.CurrentSellingPrice = wagerDto.SellingPrice
	wager.PercentageSold = 0
	wager.AmountSold = 0

	return wager
}

func PrepareWagerDTO(wager repositories.Wager) WagerDto {
	wagerDto := WagerDto{}
	wagerDto.Id = wager.ID
	wagerDto.TotalWagerValue = wager.TotalWagerValue
	wagerDto.Odds = wager.Odds
	wagerDto.SellingPercentage = wager.SellingPercentage
	wagerDto.SellingPrice = wager.SellingPrice
	wagerDto.CurrentSellingPrice = wager.CurrentSellingPrice
	wagerDto.PercentageSold = wager.PercentageSold
	wagerDto.AmountSold = wager.AmountSold
	wagerDto.PlacedAt = wager.CreatedAt
	return wagerDto
}

func PrepareBuyModel(wagerId uint, BuyingPrice float64) repositories.Buy {
	buy := repositories.Buy{}
	buy.WagerId = wagerId
	buy.BuyingPrice = BuyingPrice
	return buy
}

func PrepareBuyDTO(buy repositories.Buy) BuyWagerResponseDto {
	buyDto := BuyWagerResponseDto{}
	buyDto.Id = buy.ID
	buyDto.WagerId = buy.WagerId
	buyDto.BuyingPrice = buy.BuyingPrice
	buyDto.BoughtAt = buy.CreatedAt
	return buyDto
}

func GetErrorResponse(message string) interface{} {
	return gin.H{"error": message}
}

func ValidateCreateWagerDTO(dto CreateWagerDto) error {
	if dto.SellingPrice <= float64(dto.TotalWagerValue)*(float64(dto.SellingPercentage)/100) {
		return errors.New("selling price invalid")
	}
	return nil
}
