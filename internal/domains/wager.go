package domains

import (
	"gorm.io/gorm"
	"time"
)

type Wager struct {
	BaseModel           BaseModel `gorm:"embedded"`
	TotalWagerValue     float32   `gorm:"total_wager_value"`
	Odds                uint32    `gorm:"odds"`
	SellingPercentage   float32   `gorm:"selling_percentage"`
	SellingPrice        float32   `gorm:"selling_price"`
	CurrentSellingPrice float32   `gorm:"current_selling_price"`
	PercentageSold      float32   `gorm:"percentage_sold"`
	AmountSold          float32   `gorm:"amount_sold"`
}

type BaseModel struct {
	ID        uint            `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt *time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time      `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy string          `gorm:"column:created_by;default:SYSTEM"`
	UpdatedBy string          `gorm:"column:updated_by;default:SYSTEM"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`
	PlacedAt  *time.Time      `gorm:"column:placed_at;autoCreateTime"`
}