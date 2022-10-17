package wager

import (
	"net/http"

	"project/common/failure"
)

var (
	selling_price_under_threshold   = "selling_price_under_threshold"
	buying_price_over_selling_price = "buying_price_over_selling_price"
	wager_out_of_item               = "wager_out_of_item"
	wager_not_found                 = "wager_not_found"
)

var ErrorCodeMap = map[string]int{
	selling_price_under_threshold:   http.StatusBadRequest,
	buying_price_over_selling_price: http.StatusBadRequest,
	wager_out_of_item:               http.StatusBadRequest,
	wager_not_found:                 http.StatusNotFound,
}

var (
	SellingPriceUnderThreshold failure.ErrorWraper = failure.NewFailure(
		selling_price_under_threshold,
		ErrorCodeMap[selling_price_under_threshold],
	)

	BuyingPriceOverSellingPrice failure.ErrorWraper = failure.NewFailure(
		buying_price_over_selling_price,
		ErrorCodeMap[buying_price_over_selling_price],
	)

	WagerOutOfItem failure.ErrorWraper = failure.NewFailure(
		wager_out_of_item,
		ErrorCodeMap[wager_out_of_item],
	)

	WagerNotFound failure.ErrorWraper = failure.NewFailure(
		wager_not_found,
		ErrorCodeMap[wager_not_found],
	)
)
