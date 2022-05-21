package repositories

import (
	"betprophet1.com/wagers/internal/domains"
	"gorm.io/gorm"
)

type IPurchaseRepository interface {
	Buy(purchase *domains.Purchase) (*domains.Purchase, error)
	Count() (int64, error)
}

type PurchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (p *PurchaseRepository) Buy(purchase *domains.Purchase) (*domains.Purchase, error)  {
	err := p.db.Model(&domains.Purchase{}).Create(purchase).Error
	return purchase, err
}

func (p *PurchaseRepository) Count() (int64, error) {
	var count int64
	err := p.db.Model(&domains.Purchase{}).Count(&count).Error
	return count, err
}