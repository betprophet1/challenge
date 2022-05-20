package handlers

import (
	"betprophet1.com/wagers/internal/services"
	"fmt"
	"net/http"
	"strconv"
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
	//vars := mux.Vars(req)
	//page  := vars["page"]
	//limit := vars["limit"]
	query := req.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))

	fmt.Fprintf(res, "You've requested the book: %s on page %s\n", page, limit)
}
