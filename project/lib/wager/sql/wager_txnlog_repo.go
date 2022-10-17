package sql

import (
	"context"
	"fmt"
	"strings"

	"project/common/database"
	"project/common/database/orm"
	"project/project/dbmodels"
)

func AddWagerTxnLog(ctx context.Context, db orm.Orm, wagerTxnlogs dbmodels.WagerTxnLog) (_ dbmodels.WagerTxnLog, err error) {
	fields := database.
		Fields(dbmodels.WagerTxnLog{})
	fieldStr := strings.Join(fields, ",")
	query := fmt.Sprintf(`
		insert into wager_txn_logs(%s)
			values(?, ?, ?, ?, ?, ?)
		returning *
	`, fieldStr)
	if err = db.Raw(query,
		wagerTxnlogs.WagerID,
		wagerTxnlogs.UserID,
		wagerTxnlogs.Action,
		wagerTxnlogs.Amount,
		wagerTxnlogs.PostSellingPrice,
		wagerTxnlogs.BoughtAt,
	).Scan(&wagerTxnlogs); err != nil {
		err = fmt.Errorf("`AddWagerTxnLog` failed | err=%s", err.Error())
	}
	return wagerTxnlogs, err
}
