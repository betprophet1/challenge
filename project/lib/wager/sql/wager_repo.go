package sql

import (
	"context"
	"fmt"
	"strings"

	"project/common/database"
	"project/common/database/orm"
	"project/project/dbmodels"
)

func AddWager(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (_ dbmodels.Wager, err error) {
	fields := database.
		Fields(dbmodels.Wager{})
	fieldStr := strings.Join(fields, ",")
	query := fmt.Sprintf(`
		insert into wagers(%s)
			values(?, ?, ?, ?, ?, ?, ?, ?)
		returning *
	`, fieldStr)
	if err = db.Raw(query,
		wager.TotalWagerValue,
		wager.Odds,
		wager.SellingPercentage,
		wager.SellingPrice,
		wager.CurrentSellingPrice,
		wager.PercentageSold,
		wager.AmountSold,
		wager.PlacedAt,
	).Scan(&wager); err != nil {
		err = fmt.Errorf("`AddWagerTxnLog` failed | err=%s", err.Error())
	}
	return wager, err
}

func FetchWagerForUpdate(ctx context.Context, db orm.Orm, wagerId uint64) (wager dbmodels.Wager, err error) {
	query := `
		select * 
		from wagers 
			where id = ? for update
	`
	if err = db.Raw(query, wagerId).Scan(&wager); err != nil {
		err = fmt.Errorf("`FetchWagerForUpdate` failed | err=%s", err.Error())
	}
	return
}

func UpdateWagerSellingPrice(ctx context.Context, db orm.Orm, wager dbmodels.Wager) (err error) {
	query := `
		update wagers 
			set current_selling_price = ?, percentage_sold = ?, amount_sold = ?
		where id = ?
	`
	if err := db.Exec(query,
		wager.CurrentSellingPrice,
		wager.PercentageSold,
		wager.AmountSold,
		wager.ID); err != nil {
		err = fmt.Errorf("`UpdateWagerSellingPrice` failed | err=%s", err.Error())
	}
	return
}

func FetchWagers(ctx context.Context, db orm.Orm, limit, offset uint) (wagers []dbmodels.Wager, err error) {
	if limit == 0 {
		limit = 100
	}
	query := `
		select *
		from wagers
		limit ?
		offset ? 
	`
	if err = db.Raw(query, limit, offset).Scan(&wagers); err != nil {
		err = fmt.Errorf("`FetchWagers` failed | err=%s", err.Error())
	}
	return
}
