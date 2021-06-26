package repositories

import "github.com/jinzhu/gorm"

type Buy struct {
	gorm.Model
	WagerId     uint    `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price"`
}

type BuyReposioty struct {
	DB *gorm.DB
}

func ProvideBuyReposioty(DB *gorm.DB) BuyReposioty {
	return BuyReposioty{DB: DB}
}

func (w *BuyReposioty) Save(buy Buy) Buy {
	w.DB.Save(&buy)

	return buy
}
