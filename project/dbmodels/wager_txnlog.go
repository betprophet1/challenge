package dbmodels

import "github.com/shopspring/decimal"

type WagerTxnActionType uint8

const (
	WagerTxnBuyActionType WagerTxnActionType = iota + 1
)

type WagerTxnLog struct {
	ID      uint64             `gorm:"column:id;primaryKey"`
	WagerID uint64             `gorm:"column:wager_id"`
	UserID  string             `gorm:"column:user_id"`
	Action  WagerTxnActionType `gorm:"column:action"`
	Amount  decimal.Decimal    `gorm:"column:amount"`
	// Wager current selling price after an successful action
	PostSellingPrice decimal.Decimal `gorm:"column:post_selling_price"`
	BoughtAt         int64           `gorm:"column:bought_at"`
}

func (w WagerTxnLog) TableName() string {
	return "wager_txn_logs"
}

func (w WagerTxnLog) Fields() []string {
	return []string{
		"wager_id",
		"user_id",
		"action",
		"amount",
		"post_selling_price",
		"bought_at",
	}
}
