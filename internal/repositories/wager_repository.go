package repositories

import (
	"betprophet1.com/wagers/internal/domains"
	"betprophet1.com/wagers/pkg"
	"gorm.io/gorm"
)

type IWagerRepository interface {
	Create(wager *domains.Wager) (*domains.Wager, error)
	Update(wager *domains.Wager) (*domains.Wager, error)
	GetAll(pagination *pkg.Pagination) (*pkg.Pagination, error)
}

type WagerRepository struct {
	db *gorm.DB
}

func NewWagerRepository(db *gorm.DB) *WagerRepository {
	return &WagerRepository{db: db}
}

func (w *WagerRepository) Create(wager *domains.Wager) (*domains.Wager, error)  {
	err := w.db.Create(wager).Error
	return wager, err
}

func (w *WagerRepository) Update(wager *domains.Wager) (*domains.Wager, error)  {
	err := w.db.Model(&domains.Wager{}).Where("id = ?", wager.BaseModel.ID).Updates(wager).Error
	return wager, err
}

func (w *WagerRepository) GetAll(pagination *pkg.Pagination) (*pkg.Pagination, error) {
	var wagers []*domains.Wager
	w.db.Scopes(pagination.Paginate(wagers, w.db)).Find(&wagers)
	pagination.Rows = wagers
	return pagination, nil
}
