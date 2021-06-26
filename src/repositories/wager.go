package repositories

import "github.com/jinzhu/gorm"

type Wager struct {
	gorm.Model
	TotalWagerValue     uint    `json:"total_wager_value"`
	Odds                uint    `json:"odds"`
	SellingPercentage   int     `json:"selling_percentage"`
	SellingPrice        float64 `json:"selling_price"`
	CurrentSellingPrice float64 `json:"current_selling_price"`
	PercentageSold      int     `json:"percentage_sold"`
	AmountSold          int     `json:"amount_sold"`
}

type WagerReposioty struct {
	DB *gorm.DB
}

func ProvideWagerReposioty(DB *gorm.DB) WagerReposioty {
	return WagerReposioty{DB: DB}
}

func (w *WagerReposioty) FindAll(page int, limit int) []Wager {
	var wagers []Wager
	if page > 0 {
		// Offset start from 0
		page = page - 1
	}
	// w.DB.Find(&wagers)
	w.DB.Offset(page).Limit(limit).Find(&wagers)

	return wagers
}

func (w *WagerReposioty) FindByID(id uint) Wager {
	var wager Wager
	w.DB.First(&wager, id)

	return wager
}

func (w *WagerReposioty) Save(wager Wager) Wager {
	w.DB.Save(&wager)

	return wager
}
