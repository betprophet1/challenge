package dbmodels

import "github.com/shopspring/decimal"

type Wager struct {
	ID                  uint64          `gorm:"column:id;primaryKey"`
	TotalWagerValue     uint            `gorm:"column:total_wager_value"`
	Odds                uint            `gorm:"column:odds"`
	SellingPercentage   uint            `gorm:"column:selling_percentage"`
	SellingPrice        decimal.Decimal `gorm:"column:selling_price"`
	CurrentSellingPrice decimal.Decimal `gorm:"column:current_selling_price"`
	PercentageSold      *uint           `gorm:"column:percentage_sold"`
	AmountSold          *uint           `gorm:"column:amount_sold"`
	PlacedAt            int64           `gorm:"column:placed_at"`
}

func (w Wager) TableName() string {
	return "wagers"
}

func (w Wager) Fields() []string {
	return []string{
		"total_wager_value",
		"odds",
		"selling_percentage",
		"selling_price",
		"current_selling_price",
		"percentage_sold",
		"amount_sold",
		"placed_at",
	}
}
