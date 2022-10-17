package wager

import (
	"context"
	"fmt"
	"testing"

	"project/common/cache/client"
	"project/common/database"
	"project/common/database/orm"
	"project/common/failure"
	"project/project/dbmodels"

	"github.com/shopspring/decimal"
)

type testCase struct {
	name        string
	ctx         context.Context
	mockFun     func()
	input       interface{}
	expectedErr error
}

func TestPlaceOne(t *testing.T) {
	testCases := []testCase{
		{
			name: "happy case",
			ctx:  context.Background(),
			mockFun: func() {
				AddWagerSql = func(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (_ dbmodels.Wager, err error) {
					return dbmodels.Wager{
						TotalWagerValue:   100,
						Odds:              1,
						SellingPercentage: 35,
						SellingPrice:      decimal.NewFromInt32(250),
					}, nil
				}
			},
			input: dbmodels.Wager{
				TotalWagerValue:   100,
				Odds:              1,
				SellingPercentage: 35,
				SellingPrice:      decimal.NewFromInt32(250),
			},
			expectedErr: nil,
		},
		{
			name: "selling price under threshold",
			ctx:  context.Background(),
			mockFun: func() {
				AddWagerSql = func(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (_ dbmodels.Wager, err error) {
					return dbmodels.Wager{
						TotalWagerValue:   100,
						Odds:              1,
						SellingPercentage: 35,
						SellingPrice:      decimal.NewFromInt32(250),
					}, nil
				}
			},
			input: dbmodels.Wager{
				TotalWagerValue:   100,
				Odds:              1,
				SellingPercentage: 35,
				SellingPrice:      decimal.NewFromInt32(10),
			},
			expectedErr: SellingPriceUnderThreshold(nil),
		},
		{
			name: "create wager fail",
			ctx:  context.Background(),
			mockFun: func() {
				AddWagerSql = func(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (_ dbmodels.Wager, err error) {
					return dbmodels.Wager{}, fmt.Errorf("unexpected error")
				}
			},
			input: dbmodels.Wager{
				TotalWagerValue:   100,
				Odds:              1,
				SellingPercentage: 35,
				SellingPrice:      decimal.NewFromInt32(1000),
			},
			expectedErr: failure.InternalServerError(nil),
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.name, func(it *testing.T) {
			tcase.mockFun()
			_, err := PlaceOne(tcase.ctx, tcase.input.(dbmodels.Wager))
			if err != nil {
				errCode := err.(failure.Failure).ErrorCode()
				expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
				if errCode != expectedErrCode {
					it.Errorf("case failed | expected_err=%s/actual_err=%s", expectedErrCode, errCode)
				}
			}
		})
	}
}

type buy_one_or_part_input struct {
	wagerid     uint64
	userid      string
	buyingprice decimal.Decimal
}
type mockSoldCouter struct{}

func (m mockSoldCouter) Incr(ctx context.Context) error {
	return nil
}
func (m mockSoldCouter) Decr(ctx context.Context) error {
	return nil
}
func (m mockSoldCouter) GetCount(ctx context.Context) (uint, error) {
	return 1, nil
}
func TestBuyOneOrPart(t *testing.T) {
	testCases := []testCase{
		{
			name: "happy case",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagerForUpdateSql = func(ctx context.Context, db orm.Orm, wagerId uint64) (wager dbmodels.Wager, err error) {
					return dbmodels.Wager{
						ID:                  1,
						CurrentSellingPrice: decimal.NewFromFloat32(100),
						SellingPrice:        decimal.NewFromFloat32(100),
					}, nil
				}
				AddWagerTxnLogSql = func(ctx context.Context, db orm.Orm, wagerTxnlogs dbmodels.WagerTxnLog) (_ dbmodels.WagerTxnLog, err error) {
					return dbmodels.WagerTxnLog{}, nil
				}
				UpdateWagerSellingPriceSql = func(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (err error) {
					return nil
				}
				NewSoldCouter = func(wagerid uint64, memdb client.Client) SoldCouter {
					return &mockSoldCouter{}
				}
				database.OrmTransaction = func(callback orm.TransactionWrapperCallback) error {
					return nil
				}
			},
			input: buy_one_or_part_input{
				wagerid:     1,
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromFloat32(10),
			},
			expectedErr: nil,
		},
		{
			name: "wager not found",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagerForUpdateSql = func(ctx context.Context, db orm.Orm, wagerId uint64) (wager dbmodels.Wager, err error) {
					return dbmodels.Wager{
						ID: 0,
					}, nil
				}
			},
			input: buy_one_or_part_input{
				wagerid:     1,
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromFloat32(10),
			},
			expectedErr: WagerNotFound(nil),
		},
		{
			name: "buying_price larger than current_selling_price",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagerForUpdateSql = func(ctx context.Context, db orm.Orm, wagerId uint64) (wager dbmodels.Wager, err error) {
					return dbmodels.Wager{
						ID:                  1,
						CurrentSellingPrice: decimal.NewFromFloat32(100),
					}, nil
				}
			},
			input: buy_one_or_part_input{
				wagerid:     1,
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromFloat32(1000000),
			},
			expectedErr: BuyingPriceOverSellingPrice(nil),
		},
		{
			name: "wager out of item",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagerForUpdateSql = func(ctx context.Context, db orm.Orm, wagerId uint64) (wager dbmodels.Wager, err error) {
					return dbmodels.Wager{
						ID:                  1,
						CurrentSellingPrice: decimal.NewFromInt(0),
					}, nil
				}
			},
			input: buy_one_or_part_input{
				wagerid:     1,
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromInt(0),
			},
			expectedErr: WagerOutOfItem(nil),
		},
		{
			name: "unexpected error",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagerForUpdateSql = func(ctx context.Context, db orm.Orm, wagerId uint64) (wager dbmodels.Wager, err error) {
					return dbmodels.Wager{
						ID:                  1,
						CurrentSellingPrice: decimal.NewFromFloat32(100),
						SellingPrice:        decimal.NewFromFloat32(100),
					}, nil
				}
				AddWagerTxnLogSql = func(ctx context.Context, db orm.Orm, wagerTxnlogs dbmodels.WagerTxnLog) (_ dbmodels.WagerTxnLog, err error) {
					return dbmodels.WagerTxnLog{}, nil
				}
				UpdateWagerSellingPriceSql = func(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (err error) {
					return fmt.Errorf("unexpected error")
				}
				NewSoldCouter = func(wagerid uint64, memdb client.Client) SoldCouter {
					return &mockSoldCouter{}
				}
			},
			input: buy_one_or_part_input{
				wagerid:     1,
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromInt(100),
			},
			expectedErr: failure.InternalServerError(nil),
		},
	}
	for _, tcase := range testCases {
		t.Run(tcase.name, func(it *testing.T) {
			tcase.mockFun()
			input := tcase.input.(buy_one_or_part_input)
			_, err := BuyOneOrPart(tcase.ctx, input.wagerid, input.userid, input.buyingprice)
			if err != nil {
				errCode := err.(failure.Failure).ErrorCode()
				expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
				if errCode != expectedErrCode {
					it.Errorf("case failed | expected_err=%s/actual_err=%s", expectedErrCode, errCode)
				}
			}
		})
	}
}

type paging_input struct {
	limit  uint
	offset uint
}

func TestListWagers(t *testing.T) {
	testCases := []testCase{
		{
			name: "happy case",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagersSql = func(ctx context.Context, db orm.Orm, limit, offset uint) (wagers []dbmodels.Wager, err error) {
					return []dbmodels.Wager{
						{},
					}, nil
				}
			},
			input: paging_input{
				limit:  1,
				offset: 0,
			},
			expectedErr: nil,
		},
		{
			name: "unexpected error",
			ctx:  context.Background(),
			mockFun: func() {
				FetchWagersSql = func(ctx context.Context, db orm.Orm, limit, offset uint) (wagers []dbmodels.Wager, err error) {
					return nil, fmt.Errorf("unexpected error")
				}
			},
			input: paging_input{
				limit:  1,
				offset: 0,
			},
			expectedErr: failure.InternalServerError(nil),
		},
	}
	for _, tcase := range testCases {
		t.Run(tcase.name, func(it *testing.T) {
			tcase.mockFun()
			input := tcase.input.(paging_input)
			_, err := List(tcase.ctx, input.limit, input.offset)
			if err != nil {
				errCode := err.(failure.Failure).ErrorCode()
				expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
				if errCode != expectedErrCode {
					it.Errorf("case failed | expected_err=%s/actual_err=%s", expectedErrCode, errCode)
				}
			}
		})
	}
}
