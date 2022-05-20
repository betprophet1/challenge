package services

import (
	"betprophet1.com/wagers/internal/domains"
	"betprophet1.com/wagers/internal/repositories"
)

type WagerService interface {
	PlaceWager(wager *domains.Wager) (*domains.Wager, error)
	BuyWager(wagerId uint) (*domains.Wager, error)
	GetWagers(page uint, limit uint) ([]*domains.Wager, error)
}

type WagerServiceImpl struct {
	wagerRepository repositories.WagerRepository
}

func (w *WagerServiceImpl) PlaceWager(wager *domains.Wager) (*domains.Wager, error) {
	return w.wagerRepository.Create(wager)
}

func (w *WagerServiceImpl) BuyWager(wagerId uint) (*domains.Wager, error)  {
	return nil, nil
}

func (w *WagerServiceImpl) GetWagers(page uint, limit uint) ([]*domains.Wager, error) {
	return nil, nil
}

func NewWagerServiceImpl(wagerRepository repositories.WagerRepository) *WagerServiceImpl {
	return &WagerServiceImpl{wagerRepository: wagerRepository}
}
