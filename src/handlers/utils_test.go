package handlers_test

import (
	"handlers"
	"repositories"
	"testing"
)

func TestPrepareWagerModel(t *testing.T) {
	dto := handlers.CreateWagerDto{
		TotalWagerValue:   11,
		Odds:              1,
		SellingPercentage: 100,
		SellingPrice:      11,
	}
	wager := handlers.PrepareWagerModel(dto)
	if wager.CurrentSellingPrice != wager.SellingPrice {
		t.Errorf("PrepareWagerModel failed, expect %f, got %f", wager.SellingPrice, wager.CurrentSellingPrice)
	} else {
		t.Logf("PrepareWagerModel passed, expect %f, got %f", wager.SellingPrice, wager.CurrentSellingPrice)
	}
}

func TestPrepareWagerDTO(t *testing.T) {
	wager := repositories.Wager{
		TotalWagerValue:     11,
		Odds:                1,
		SellingPercentage:   100,
		SellingPrice:        12,
		CurrentSellingPrice: 12,
		PercentageSold:      0,
		AmountSold:          0,
	}
	wagerDTO := handlers.PrepareWagerDTO(wager)
	var isError bool = false
	if wagerDTO.TotalWagerValue != wager.TotalWagerValue {
		t.Errorf("PrepareWagerDTO failed, expect %d, got %d", wager.TotalWagerValue, wagerDTO.TotalWagerValue)
		isError = true
	}
	if wagerDTO.SellingPercentage != wager.SellingPercentage {
		t.Errorf("PrepareWagerDTO failed, expect %d, got %d", wager.SellingPercentage, wagerDTO.SellingPercentage)
		isError = true
	}
	if wagerDTO.CurrentSellingPrice != wager.CurrentSellingPrice {
		t.Errorf("PrepareWagerDTO failed, expect %f, got %f", wager.SellingPrice, wagerDTO.SellingPrice)
		isError = true
	}
	if wagerDTO.PercentageSold != wager.PercentageSold {
		t.Errorf("PrepareWagerDTO failed, expect %d, got %d", wager.TotalWagerValue, wagerDTO.TotalWagerValue)
		isError = true
	}
	if isError {
		return
	}

	t.Logf("PrepareWagerDTO passed")
}

func TestPrepareBuyModel(t *testing.T) {
	var wagerId uint = 1
	var BuyingPrice float64 = 10
	var isError bool = false
	buyModel := handlers.PrepareBuyModel(wagerId, BuyingPrice)
	if buyModel.WagerId != wagerId {
		t.Errorf("PrepareBuyModel failed, expect %d, got %d", wagerId, buyModel.WagerId)
		isError = true
	}
	if buyModel.BuyingPrice != BuyingPrice {
		t.Errorf("PrepareBuyModel failed, expect %f, got %f", BuyingPrice, buyModel.BuyingPrice)
		isError = true
	}
	if isError {
		return
	}
	t.Logf("PrepareBuyModel passed")
}

func TestPrepareBuyDTO(t *testing.T) {
	buy := repositories.Buy{
		WagerId:     1,
		BuyingPrice: 10,
	}
	var isError bool = false
	buyDTO := handlers.PrepareBuyDTO(buy)
	if buyDTO.WagerId != buy.WagerId {
		t.Errorf("PrepareBuyDTO failed, expect %d, got %d", buy.WagerId, buyDTO.WagerId)
		isError = true
	}
	if buyDTO.BuyingPrice != buy.BuyingPrice {
		t.Errorf("PrepareBuyDTO failed, expect %f, got %f", buy.BuyingPrice, buyDTO.BuyingPrice)
		isError = true
	}
	if isError {
		return
	}
	t.Logf("PrepareBuyDTO passed")
}

func TestValidateCreateWagerDTO(t *testing.T) {
	invalidDTO := handlers.CreateWagerDto{
		TotalWagerValue:   11,
		SellingPercentage: 100,
		SellingPrice:      10,
	}
	validDTO := handlers.CreateWagerDto{
		TotalWagerValue:   11,
		SellingPercentage: 100,
		SellingPrice:      12,
	}
	var isError bool = false
	error := handlers.ValidateCreateWagerDTO(invalidDTO)
	if error == nil {
		t.Errorf("PrepareBuyDTO failed, expect %v, got %v", error, nil)
		isError = true
	}
	error = handlers.ValidateCreateWagerDTO(validDTO)
	if error != nil {
		t.Errorf("PrepareBuyDTO failed, expect %v, got %v", nil, error)
		isError = true
	}
	if isError {
		return
	}
	t.Logf("ValidateCreateWagerDTO passed")
}
