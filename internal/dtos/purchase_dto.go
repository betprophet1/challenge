package dtos

import (
	"time"
)

type PurchaseRequestDto struct {
	BuyingPrice float32 `json:"buying_price"`
	WagerId     uint    `json:"wager_id"`
}

type PurchaseResponseDto struct {
	Id          uint       `json:"id"`
	WagerId     uint       `json:"wager_id"`
	BuyingPrice float32    `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}