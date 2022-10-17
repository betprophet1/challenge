package database

import "project/common/database/orm"

func GetOrm() orm.Orm {
	return orm.GetGormOrm()
}

func Transaction(callback orm.TransactionWrapperCallback) error {
	return orm.Transaction(callback)
}
