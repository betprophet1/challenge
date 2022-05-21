package repositories

import (
	"betprophet1.com/wagers/internal/domains"
	"betprophet1.com/wagers/pkg"
	"gorm.io/gorm"
)

type WagerRepository interface {
	Create(wager *domains.Wager) (*domains.Wager, error)
	Update(wager *domains.Wager) (*domains.Wager, error)
	GetAll(pagination *pkg.Pagination) (*pkg.Pagination, error)
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

func (w *WagerRepositoryImpl) GetAll(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var wagers []*domains.Wager
	w.db.Scopes(pagination.Paginate(wagers, &w.db)).Find(&wagers)
	pagination.Rows = wagers
	return pagination, nil
}

func NewWagerRepositoryImpl(db gorm.DB) *WagerRepositoryImpl {
	return &WagerRepositoryImpl{db: db}
}
