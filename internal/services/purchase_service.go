package services

import (
	"betprophet1.com/wagers/internal/domains"
	"betprophet1.com/wagers/internal/dtos"
	"betprophet1.com/wagers/internal/repositories"
	"errors"
)

type IPurchaseService interface {
	Buy(purchase *dtos.PurchaseRequestDto) (*dtos.PurchaseResponseDto, error)
}

type PurchaseService struct {
	purchaseRepository repositories.IPurchaseRepository
	wagerRepository    repositories.IWagerRepository
}

func NewPurchaseService(purchaseRepository repositories.IPurchaseRepository, wagerRepository repositories.IWagerRepository) *PurchaseService {
	return &PurchaseService{purchaseRepository: purchaseRepository, wagerRepository: wagerRepository}
}

func (p *PurchaseService) Buy(purchase *dtos.PurchaseRequestDto) (*dtos.PurchaseResponseDto, error) {
	wager, err := p.wagerRepository.GetById(purchase.WagerId)
	if err != nil {
		return nil, errors.New("Wagger doesn't exist to buy")
	}

	currentSellingPrice := wager.CurrentSellingPrice
	if purchase.BuyingPrice > currentSellingPrice {
		return nil, errors.New("Buying Price must be lesser or equal current selling price")
	}

	buyPurchase, err := p.purchaseRepository.Buy(&domains.Purchase{
		WagerId:     purchase.WagerId,
		BuyingPrice: purchase.BuyingPrice,
	})
	if err != nil {
		return nil, errors.New("Buying is failed")
	}

	if currentSellingPrice >= buyPurchase.BuyingPrice {
		currentSellingPrice = buyPurchase.BuyingPrice
	}
	wager.CurrentSellingPrice = currentSellingPrice
	wager.PercentageSold      = (wager.CurrentSellingPrice / wager.TotalWagerValue) * 100
	amountSold, _ := p.purchaseRepository.Count()
	wager.AmountSold = float32(amountSold)

	_, _ = p.wagerRepository.Update(wager)

	return &dtos.PurchaseResponseDto{
		Id:          buyPurchase.BaseModel.ID,
		WagerId:     buyPurchase.WagerId,
		BuyingPrice: buyPurchase.BuyingPrice,
		BoughtAt:    buyPurchase.BoughtAt,
	}, err
}
