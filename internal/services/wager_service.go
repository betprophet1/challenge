package services

import (
	"betprophet1.com/wagers/internal/domains"
	"betprophet1.com/wagers/internal/dtos"
	"betprophet1.com/wagers/internal/repositories"
	"betprophet1.com/wagers/pkg"
)

type IWagerService interface {
	PlaceWager(wager *dtos.WagerRequestDto) (*domains.Wager, error)
	GetWagers(page int, limit int) (*pkg.Pagination, error)
}

type WagerService struct {
	wagerRepository repositories.IWagerRepository
}

func NewWagerService(wagerRepository repositories.IWagerRepository) *WagerService {
	return &WagerService{wagerRepository: wagerRepository}
}

func (w *WagerService) PlaceWager(wager *dtos.WagerRequestDto) (*domains.Wager, error) {
	wagerDomain := &domains.Wager{
		TotalWagerValue:   wager.TotalWagerValue,
		Odds:              wager.Odds,
		SellingPercentage: wager.SellingPercentage,
		SellingPrice:      wager.SellingPrice,
	}
	return w.wagerRepository.Create(wagerDomain)
}

func (w *WagerService) GetWagers(page int, limit int) (*pkg.Pagination, error) {
	pagination := &pkg.Pagination{
		Limit:      limit,
		Page:       page,
	}
	wagers, err := w.wagerRepository.GetAll(pagination)
	if err != nil {
		return nil, err
	}

	var wagerDtos []*dtos.WagerResponseDto
	for _, value := range wagers.Rows.([]*domains.Wager) {
		wagerDtos = append(wagerDtos, value.ConvertToDto())
	}
	wagers.Rows = wagerDtos
	return wagers, nil
}
