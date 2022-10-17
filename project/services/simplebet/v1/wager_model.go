package v1

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	PlaceOneWagerRequest struct {
		TotalWagerValue   uint            `json:"total_wager_value" binding:"required,min=1"`
		Odds              uint            `json:"odds" binding:"required,min=1"`
		SellingPercentage uint            `json:"selling_percentage" binding:"required,min=1,max=100"`
		SellingPrice      decimal.Decimal `json:"selling_price" binding:"required,min=1"`
	}
	PlaceOneWagerResponse struct {
		ID                  uint64          `json:"id"`
		TotalWagerValue     uint            `json:"total_wager_value"`
		Odds                uint            `json:"odds"`
		SellingPercentage   uint            `json:"selling_percentage"`
		SellingPrice        decimal.Decimal `json:"selling_price"`
		CurrentSellingPrice decimal.Decimal `json:"current_selling_price"`
		PercentageSold      *uint           `json:"percentage_sold"`
		AmountSold          *uint           `json:"amount_sold"`
		PlacedAt            time.Time       `json:"placed_at"`
	}
)

type (
	WagerID struct {
		Value uint64 `uri:"wager_id" binding:"required,min=1"`
	}
	BuyFullOrPartOneWagerRequest struct {
		WagerID     WagerID         `binding:"-"`
		UserID      string          `json:"user_id" binding:"required"`
		BuyingPrice decimal.Decimal `json:"buying_price" binding:"required,min=1"`
	}
	BuyFullOrPartOneWagerResponse struct {
		ID          uint64          `json:"id"`
		WagerID     uint64          `json:"wager_id"`
		BuyingPrice decimal.Decimal `json:"buying_price"`
		BoughtAt    time.Time       `json:"bought_at"`
	}
)

type (
	Paging struct {
		Limit  uint `form:"limit"`
		Offset uint `form:"offset"`
	}
	ListWagersResponse struct {
		ID                  uint64          `json:"id"`
		TotalWagerValue     uint            `json:"total_wager_value"`
		Odds                uint            `json:"odds"`
		SellingPercentage   uint            `json:"selling_percentage"`
		SellingPrice        decimal.Decimal `json:"selling_price"`
		CurrentSellingPrice decimal.Decimal `json:"current_selling_price"`
		PercentageSold      *uint           `json:"percentage_sold"`
		AmountSold          *uint           `json:"amount_sold"`
		PlacedAt            time.Time       `json:"placed_at"`
	}
)
