package orm

import (
	"os"
	"strconv"
	"time"

	"project/common"
	"project/common/rootcloser"
	"project/common/singleton"
	"project/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var gormOrm Orm

func InitGorm() {
	singleGorm := singleton.New(func() interface{} {
		db, err := gorm.Open(postgres.Open(os.Getenv(config.SERVICE_POSTGRES_DNS)), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            false,
		})
		common.PanicOnError(err)

		sqlDB, err := db.DB()
		common.PanicOnError(err)

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		isDebug, _ := strconv.ParseBool(os.Getenv(config.DEBUG))
		if isDebug {
			db = db.Debug()
		}
		instance := &GormOrm{db}
		rootcloser.Register(func() {
			sqlDB.Close()
		})
		return instance
	})
	gormOrm = singleGorm.Get().(*GormOrm)
}

type GormOrm struct {
	db *gorm.DB
}

func (g GormOrm) Scan(data interface{}) error {
	return g.db.Scan(data).Error
}

func (g GormOrm) Exec(sql string, values ...interface{}) error {
	return g.db.Exec(sql, values...).Error
}

func (g *GormOrm) Raw(sql string, values ...interface{}) RawResult {
	g.db = g.db.Raw(sql, values...)
	return g
}

func (g *GormOrm) Transaction(callback TransactionCallback) error {
	return g.db.Transaction(func(tx *gorm.DB) error {
		txn := &GormOrm{db: tx}
		if err := callback(txn); err != nil {
			return err
		}
		return nil
	})
}

func (g GormOrm) Upsert(values interface{}, batchSize int, updateFields []string) error {
	return g.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{clause.PrimaryColumn},
		DoUpdates: clause.AssignmentColumns(updateFields),
	}).CreateInBatches(values, batchSize).Error
}

func GetGormOrm() Orm {
	return gormOrm
}
