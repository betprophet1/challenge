package handlers

import (
	"betprophet1.com/wagers/internal/services"
	"net/http"
)

type WagerHandler interface {
	PlaceWager(res http.ResponseWriter, req *http.Request)
	BuyWager(res http.ResponseWriter, req *http.Request)
	ListWager(res http.ResponseWriter, req *http.Request)
}

type WagerHandlerImpl struct {
	wagerService services.WagerService
}

func NewWagerHandlerImpl(wagerService services.WagerService) *WagerHandlerImpl {
	return &WagerHandlerImpl{wagerService: wagerService}
}

func (w *WagerHandlerImpl) PlaceWager(res http.ResponseWriter, req *http.Request) {

}

func (w *WagerHandlerImpl) BuyWager(res http.ResponseWriter, req *http.Request) {

}

func (w *WagerHandlerImpl) ListWager(res http.ResponseWriter, req *http.Request) {

}
