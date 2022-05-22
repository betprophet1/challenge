package dtos

import (
	"reflect"
	"time"
)

type WagerRequestDto struct {
	TotalWagerValue   float32 `json:"total_wager_value"`
	Odds              float32 `json:"odds"`
	SellingPercentage float32 `json:"selling_percentage"`
	SellingPrice      float32 `json:"selling_price"`
}

type WagerResponseDto struct {
	Id                  uint       `json:"id"`
	TotalWagerValue     float32    `json:"total_wager_value"`
	Odds                float32    `json:"odds"`
	SellingPercentage   float32    `json:"selling_percentage"`
	SellingPrice        float32    `json:"selling_price"`
	CurrentSellingPrice float32    `json:"current_selling_price"`
	PercentageSold      float32    `json:"percentage_sold"`
	AmountSold          float32    `json:"amount_sold"`
	PlacedAt            *time.Time `json:"placed_at"`
}

type WagerErrorResponse struct {
	Error string `json:"error"`
}

func (w *WagerRequestDto) IsEmpty() bool {
	return reflect.DeepEqual(w, &WagerRequestDto{})
}