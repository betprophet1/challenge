package orm

type (
	// General ORM interface
	Orm interface {
		Raw(sql string, values ...interface{}) RawResult
		Exec(sql string, values ...interface{}) error
		Upsert(value interface{}, batchSize int, updateFields []string) error
		Transaction(callback TransactionCallback) error
	}

	RawResult interface {
		Scan(data interface{}) error
	}
)
