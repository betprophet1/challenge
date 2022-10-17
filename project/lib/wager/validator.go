package wager

import (
	"fmt"

	"project/project/dbmodels"

	"github.com/shopspring/decimal"
)

type WagerValidator dbmodels.Wager

/*
	ValidateSellingPrice returns error if
	selling_price <= total_wager_value * (selling_percentage/100)
*/
func (wv WagerValidator) ValidateSellingPrice() (err error) {
	sellingTheshold := (float32(wv.SellingPercentage) / 100.0) * float32(wv.TotalWagerValue)
	ok := wv.SellingPrice.LessThanOrEqual(
		decimal.NewFromFloat32(sellingTheshold))
	if ok {
		err = SellingPriceUnderThreshold(
			fmt.Errorf("selling_price=%v/selling_thershold=%v", wv.SellingPrice, sellingTheshold))
	}
	return
}
