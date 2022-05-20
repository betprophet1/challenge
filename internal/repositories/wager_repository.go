package repositories

import (
	"betprophet1.com/wagers/internal/domains"
	"gorm.io/gorm"
)

type WagerRepository interface {
	Create(wager *domains.Wager) (*domains.Wager, error)
	Update(wager *domains.Wager) (*domains.Wager, error)
	GetAll() ([]*domains.Wager, error)
}

type WagerRepositoryImpl struct {
	db gorm.DB
}

func (w *WagerRepositoryImpl) Create(wager *domains.Wager) (*domains.Wager, error)  {
	err := w.db.Create(wager).Error
	return wager, err
}

func (w *WagerRepositoryImpl) Update(wager *domains.Wager) (*domains.Wager, error)  {
	err := w.db.Model(&domains.Wager{}).Where("id = ?", wager.BaseModel.ID).Updates(wager).Error
	return wager, err
}

func (w *WagerRepositoryImpl) GetAll() ([]*domains.Wager, error) {
	var wagers []*domains.Wager
	err := w.db.Model(&domains.Wager{}).Find(&wagers).Error
	return wagers, err
}
func NewWagerRepositoryImpl(db gorm.DB) *WagerRepositoryImpl {
	return &WagerRepositoryImpl{db: db}
}
