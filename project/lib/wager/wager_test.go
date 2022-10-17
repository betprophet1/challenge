package wager

import (
	"context"
	"testing"

	"project/common"
	"project/common/cache/client"
	"project/common/database/orm"
	"project/common/failure"
	"project/project/dbmodels"

	"github.com/shopspring/decimal"
)

func prerunTest() {
	orm.InitGorm()
	client.InitRedis()
}

func prepareWager() uint64 {
	wager, err := PlaceOne(context.Background(), dbmodels.Wager{
		TotalWagerValue:   100,
		Odds:              1,
		SellingPercentage: 35,
		SellingPrice:      decimal.NewFromInt32(250),
	})
	common.PanicOnError(err)
	return wager.ID
}

type testCase struct {
	name        string
	ctx         func() context.Context
	input       interface{}
	expectedErr error
}

func TestPlaceOne(t *testing.T) {
	prerunTest()
	testCases := []testCase{
		{
			name: "happy case",
			ctx: func() context.Context {
				return context.Background()
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
			ctx: func() context.Context {
				return context.Background()
			},
			input: dbmodels.Wager{
				TotalWagerValue:   100,
				Odds:              1,
				SellingPercentage: 35,
				SellingPrice:      decimal.NewFromInt32(10),
			},
			expectedErr: SellingPriceUnderThreshold(nil),
		},
	}

	for _, tcase := range testCases {
		t.Logf("Running `%s` case...", tcase.name)
		_, err := PlaceOne(tcase.ctx(), tcase.input.(dbmodels.Wager))
		if err != nil {
			errCode := err.(failure.Failure).ErrorCode()
			expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
			if errCode != expectedErrCode {
				t.Errorf("case failed | expected_err=%s/actual_err=%s", expectedErrCode, errCode)
				continue
			}
		} else {
			if tcase.expectedErr != nil {
				expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
				t.Errorf("case failed | expected_err=%s/actual_err=%v", expectedErrCode, err)
				continue
			}
		}
		t.Logf("case passed | output: %v", err)
	}
}

type buy_one_or_part_input struct {
	setWagerID  func(id uint64) uint64
	userid      string
	buyingprice decimal.Decimal
}

func TestBuyOneOrPart(t *testing.T) {
	prerunTest()
	testCases := []testCase{
		{
			name: "happy case",
			ctx: func() context.Context {
				return context.Background()
			},
			input: buy_one_or_part_input{
				setWagerID: func(id uint64) uint64 {
					return id
				},
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromFloat32(10),
			},
			expectedErr: nil,
		},
		{
			name: "wager not found",
			ctx: func() context.Context {
				return context.Background()
			},
			input: buy_one_or_part_input{
				setWagerID: func(id uint64) uint64 {
					return id + 999999
				},
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromFloat32(10),
			},
			expectedErr: WagerNotFound(nil),
		},
		{
			name: "buying_price larger than current_selling_price",
			ctx: func() context.Context {
				return context.Background()
			},
			input: buy_one_or_part_input{
				setWagerID: func(id uint64) uint64 {
					return id
				},
				userid:      "thanhdinh",
				buyingprice: decimal.NewFromFloat32(1000000),
			},
			expectedErr: BuyingPriceOverSellingPrice(nil),
		},
	}
	for _, tcase := range testCases {
		preparedWagerId := prepareWager()
		input := tcase.input.(buy_one_or_part_input)
		testWagerId := input.setWagerID(preparedWagerId)
		t.Logf("Running `%s` case...", tcase.name)
		_, err := BuyOneOrPart(tcase.ctx(), testWagerId, input.userid, input.buyingprice)
		if err != nil {
			errCode := err.(failure.Failure).ErrorCode()
			expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
			if errCode != expectedErrCode {
				t.Errorf("case failed | expected_err=%s/actual_err=%s", expectedErrCode, errCode)
				continue
			}
		} else {
			if tcase.expectedErr != nil {
				expectedErrCode := tcase.expectedErr.(failure.Failure).ErrorCode()
				t.Errorf("case failed | expected_err=%s/actual_err=%v", expectedErrCode, err)
				continue
			}
		}
		t.Logf("case passed | output: %v", err)
	}
}

type paging_input struct {
	limit  uint
	offset uint
}

func TestListWagers(t *testing.T) {
	prerunTest()
	testCases := []testCase{
		{
			name: "happy case",
			ctx: func() context.Context {
				return context.Background()
			},
			input: paging_input{
				limit:  1,
				offset: 0,
			},
			expectedErr: nil,
		},
	}
	for _, tcase := range testCases {
		input := tcase.input.(paging_input)
		t.Logf("Running `%s` case...", tcase.name)
		wagers, err := List(tcase.ctx(), input.limit, input.offset)
		if err != nil {
			errCode := err.(failure.Failure).ErrorCode()
			t.Errorf("case failed | expected_err=%v/actual_err=%s", nil, errCode)
			continue
		}
		if len(wagers) != int(input.limit) {
			t.Errorf("case failed | expected_limit=%d/actual_limit=%d", input.limit, len(wagers))
			continue
		}
		t.Logf("case passed | output: %v", wagers)
	}
}
