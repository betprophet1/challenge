package handlers

import (
	"betprophet1.com/wagers/internal/dtos"
	"betprophet1.com/wagers/internal/services"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type IWagerHandler interface {
	PlaceWager(res http.ResponseWriter, req *http.Request)
	BuyWager(res http.ResponseWriter, req *http.Request)
	ListWager(res http.ResponseWriter, req *http.Request)
}

type WagerHandler struct {
	wagerService    services.IWagerService
	purchaseService services.IPurchaseService
}

func NewWagerHandler(wagerService services.IWagerService, purchaseService services.IPurchaseService) *WagerHandler {
	return &WagerHandler{wagerService: wagerService, purchaseService: purchaseService}
}



func (w *WagerHandler) PlaceWager(res http.ResponseWriter, req *http.Request) {
	wager := &dtos.WagerRequestDto{}
	if err := json.NewDecoder(req.Body).Decode(wager); err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: errors.New("Body must not be null or empty").Error()})
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(e)
		return
	}

	placeWager, err := w.wagerService.PlaceWager(wager)
	if err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: err.Error()})
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(e)
		return
	}

	response, err := json.Marshal(placeWager.ConvertToDto())
	if err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: err.Error()})
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(e)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(response)
}

func (w *WagerHandler) BuyWager(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	wagerId, _ := strconv.Atoi(vars["wager_id"])
	purchaseReq := &dtos.PurchaseRequestDto{
		WagerId:     uint(wagerId),
	}

	if err := json.NewDecoder(req.Body).Decode(purchaseReq); err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: err.Error()})
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(e)
		return
	}
	purchase, err := w.purchaseService.Buy(purchaseReq)
	if err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: err.Error()})
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(e)
		return
	}

	response, _ := json.Marshal(purchase)
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}

func (w *WagerHandler) ListWager(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	page, _  := strconv.Atoi(vars["page"])
	limit, _ := strconv.Atoi(vars["limit"])

	wagers, err := w.wagerService.GetWagers(page, limit)
	if err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: err.Error()})
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(e)
		return
	}

	response, _ := json.Marshal(wagers)
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(response)
}
