package services

import (
	"betprophet1.com/wagers/internal/domains"
	"betprophet1.com/wagers/internal/dtos"
	"betprophet1.com/wagers/internal/repositories"
)

type IPurchaseService interface {
	Buy(purchase *dtos.PurchaseRequestDto) (*dtos.PurchaseResponseDto, error)
}

type PurchaseService struct {
	purchaseRepository repositories.IPurchaseRepository
}

func (p *PurchaseService) Buy(purchase *dtos.PurchaseRequestDto) (*dtos.PurchaseResponseDto, error) {
	buyPurchase, err := p.purchaseRepository.Buy(&domains.Purchase{
		WagerId:     purchase.WagerId,
		BuyingPrice: purchase.BuyingPrice,
	})
	return &dtos.PurchaseResponseDto{
		Id:          buyPurchase.BaseModel.ID,
		WagerId:     buyPurchase.WagerId,
		BuyingPrice: buyPurchase.BuyingPrice,
		BoughtAt:    buyPurchase.BoughtAt,
	}, err
}

func NewPurchaseService(purchaseRepository repositories.IPurchaseRepository) *PurchaseService {
	return &PurchaseService{purchaseRepository: purchaseRepository}
}
