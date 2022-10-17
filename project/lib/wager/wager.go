package wager

import (
	"context"
	"fmt"
	"time"

	"project/common/cache"
	"project/common/database"
	"project/common/database/orm"
	"project/common/failure"
	"project/project/dbmodels"

	"github.com/shopspring/decimal"
)

func PlaceOne(ctx context.Context, wager dbmodels.Wager) (_ dbmodels.Wager, err error) {
	validator := WagerValidator(wager)
	if err = validator.ValidateSellingPrice(); err != nil {
		return
	}

	wager.CurrentSellingPrice = wager.SellingPrice
	wager.PercentageSold = nil
	wager.AmountSold = nil
	wager.PlacedAt = time.Now().Unix()

	db := database.GetOrm()
	wager, err = AddWagerSql(ctx, db, wager)
	if err != nil {
		err = failure.InternalServerError(err)
	}
	return wager, err
}

/*
	Note on percentage_sold:
		(selling price / selling price) * 100 = 100% -> sold out
		current_selling_price -> decreases from selling_price to 0
*/
func BuyOneOrPart(ctx context.Context, wagerID uint64, userID string, buyingPrice decimal.Decimal) (wagerTxnLog dbmodels.WagerTxnLog, err error) {
	err = database.Transaction(func(txnWrapper *orm.TransactionWrapper) orm.TransactionCallback {
		return func(innerDB orm.Orm) (err error) {
			// lock on wager<id>
			wager, err := FetchWagerForUpdateSql(ctx, innerDB, wagerID)
			if err != nil {
				err = failure.InternalServerError(err)
				return
			}
			if wager.ID == 0 {
				err = WagerNotFound(fmt.Errorf("wager_id=%d", wagerID))
				return
			}
			if buyingPrice.
				GreaterThan(wager.CurrentSellingPrice) {
				err = BuyingPriceOverSellingPrice(
					fmt.Errorf("buying_price=%v/current_selling_price=%v", buyingPrice, wager.CurrentSellingPrice))
				return
			}
			if wager.CurrentSellingPrice.IsZero() {
				err = WagerOutOfItem(nil)
				return
			}

			wagerTxnLog = dbmodels.WagerTxnLog{
				WagerID:  wager.ID,
				UserID:   userID,
				Action:   dbmodels.WagerTxnBuyActionType,
				BoughtAt: time.Now().Unix(),
			}
			wagerTxnLog.Amount = buyingPrice.Neg()
			wagerTxnLog.PostSellingPrice = wager.CurrentSellingPrice.Sub(buyingPrice)
			wagerTxnLog, err = AddWagerTxnLogSql(ctx, innerDB, wagerTxnLog)
			if err != nil {
				err = failure.InternalServerError(err)
				return
			}

			counter := NewSoldCouter(wagerID, cache.GetClient())
			if err = counter.Incr(ctx); err != nil {
				err = failure.InternalServerError(err)
				return
			}
			txnWrapper.RegisterOnRollbackCallback(func() error {
				return counter.Decr(ctx)
			})
			soldCount, err := counter.GetCount(ctx)
			if err != nil {
				err = failure.InternalServerError(err)
				return
			}

			wager.CurrentSellingPrice = wagerTxnLog.PostSellingPrice
			percentageSold := uint(wager.SellingPrice.
				Sub(wager.CurrentSellingPrice).
				Div(wager.SellingPrice).
				Mul(decimal.NewFromInt(100)).
				BigInt().Uint64())
			wager.PercentageSold = &percentageSold
			wager.AmountSold = &soldCount
			if err = UpdateWagerSellingPriceSql(ctx, innerDB, wager); err != nil {
				err = failure.InternalServerError(err)
			}
			return
		}
	})
	return
}

func List(ctx context.Context, limit, offset uint) (wagers []dbmodels.Wager, err error) {
	db := orm.GetGormOrm()
	wagers, err = FetchWagersSql(ctx, db, limit, offset)
	if err != nil {
		err = failure.InternalServerError(err)
	}
	return
}
