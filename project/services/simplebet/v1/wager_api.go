package v1

import (
	"net/http"
	"time"

	"project/common/failure"
	"project/project/dbmodels"
	"project/project/lib/wager"

	"github.com/gin-gonic/gin"
)

func PlaceOneWager(ctx *gin.Context) {
	var reqModel PlaceOneWagerRequest
	if err := ctx.ShouldBind(&reqModel); err != nil {
		ctx.Error(failure.BadRequestError(err))
		return
	}
	mwager := dbmodels.Wager{
		TotalWagerValue:   reqModel.TotalWagerValue,
		Odds:              reqModel.Odds,
		SellingPercentage: reqModel.SellingPercentage,
		SellingPrice:      reqModel.SellingPrice,
	}
	mwager, err := wager.PlaceOne(ctx, mwager)
	if err != nil {
		ctx.Error(failure.BadRequestError(err))
		return
	}
	ctx.JSON(http.StatusOK, PlaceOneWagerResponse{
		ID:                  mwager.ID,
		TotalWagerValue:     mwager.TotalWagerValue,
		Odds:                mwager.Odds,
		SellingPercentage:   mwager.SellingPercentage,
		SellingPrice:        mwager.SellingPrice,
		CurrentSellingPrice: mwager.CurrentSellingPrice,
		PercentageSold:      mwager.PercentageSold,
		AmountSold:          mwager.AmountSold,
		PlacedAt:            time.Unix(mwager.PlacedAt, 0),
	})
}

func BuyFullOrPartOneWager(ctx *gin.Context) {
	var reqModel BuyFullOrPartOneWagerRequest
	if err := ctx.ShouldBindUri(&reqModel.WagerID); err != nil {
		ctx.Error(failure.BadRequestError(err))
		return
	}
	if err := ctx.ShouldBind(&reqModel); err != nil {
		ctx.Error(failure.BadRequestError(err))
		return
	}
	mwagerTxnlog, err := wager.BuyOneOrPart(ctx, reqModel.WagerID.Value, reqModel.UserID, reqModel.BuyingPrice)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, BuyFullOrPartOneWagerResponse{
		ID:          mwagerTxnlog.ID,
		WagerID:     mwagerTxnlog.WagerID,
		BuyingPrice: reqModel.BuyingPrice,
		BoughtAt:    time.Unix(mwagerTxnlog.BoughtAt, 0),
	})
}

func ListWagers(ctx *gin.Context) {
	var reqModel Paging
	if err := ctx.ShouldBindQuery(&reqModel); err != nil {
		ctx.Error(failure.BadRequestError(err))
		return
	}
	wagers, err := wager.List(ctx, reqModel.Limit, reqModel.Offset)
	if err != nil {
		ctx.Error(err)
		return
	}
	listWagersResponse := make([]ListWagersResponse, 0, len(wagers))
	for _, wager := range wagers {
		listWagersResponse = append(listWagersResponse, ListWagersResponse{
			ID:                  wager.ID,
			TotalWagerValue:     wager.TotalWagerValue,
			Odds:                wager.Odds,
			SellingPercentage:   wager.SellingPercentage,
			SellingPrice:        wager.SellingPrice,
			CurrentSellingPrice: wager.CurrentSellingPrice,
			PercentageSold:      wager.PercentageSold,
			AmountSold:          wager.AmountSold,
			PlacedAt:            time.Unix(wager.PlacedAt, 0),
		})
	}
	ctx.JSON(http.StatusOK, listWagersResponse)
}
