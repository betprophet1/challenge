package handlers

import (
	"betprophet1.com/wagers/internal/dtos"
	"betprophet1.com/wagers/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
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
	wager := &dtos.WagerRequestDto{}
	if err := json.NewDecoder(req.Body).Decode(wager); err != nil {
		e, _ := json.Marshal(&dtos.WagerErrorResponse{Error: err.Error()})
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

	converter := &dtos.WagerResponseDto{
		Id:                  placeWager.BaseModel.ID,
		TotalWagerValue:     placeWager.TotalWagerValue,
		Odds:                float32(placeWager.Odds),
		SellingPercentage:   placeWager.SellingPercentage,
		SellingPrice:        placeWager.SellingPrice,
		CurrentSellingPrice: placeWager.CurrentSellingPrice,
		PercentageSold:      placeWager.PercentageSold,
		AmountSold:          placeWager.AmountSold,
		PlacedAt:            placeWager.BaseModel.PlacedAt,
	}
	response, err := json.Marshal(converter)
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

func (w *WagerHandlerImpl) BuyWager(res http.ResponseWriter, req *http.Request) {

}

func (w *WagerHandlerImpl) ListWager(res http.ResponseWriter, req *http.Request) {
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
