package orm

import (
	"project/common/log"
)

type (
	TransactionWrapper struct {
		ext                 Orm
		log                 *log.Logger
		onCommitCallbacks   []TransactionOnCallback
		onRollbackCallbacks []TransactionOnCallback
	}

	TransactionWrapperCallback = func(txnWrapper *TransactionWrapper) TransactionCallback
	TransactionCallback        = func(innerDB Orm) error
	TransactionOnCallback      = func() error
)

// Call transaction
func Transaction(callback TransactionWrapperCallback) error {
	txnWrapper := newTransactionWrapper()
	return txnWrapper.execute(callback(txnWrapper))
}

func newTransactionWrapper() *TransactionWrapper {
	return &TransactionWrapper{
		ext:                 GetGormOrm(),
		log:                 &log.Logger{},
		onCommitCallbacks:   make([]TransactionOnCallback, 0),
		onRollbackCallbacks: make([]TransactionOnCallback, 0),
	}
}

func (txn *TransactionWrapper) execute(callback TransactionCallback) error {
	if err := txn.ext.Transaction(callback); err != nil {
		if len(txn.onRollbackCallbacks) > 0 {
			for _, onRollbackCallback := range txn.onRollbackCallbacks {
				if cError := onRollbackCallback(); cError != nil {
					// commit callbacks would be considered for further warning handling
					txn.log.Error("run callback failed | err=%s", cError.Error())
				}
			}
		}
		return err
	}

	if len(txn.onCommitCallbacks) > 0 {
		for _, onCommitCallback := range txn.onCommitCallbacks {
			if err := onCommitCallback(); err != nil {
				// commit callbacks would be considered for further warning handling
				txn.log.Error("run callback failed | err=%s", err.Error())
			}
		}
	}
	return nil
}

// Register caller for successful commited transaction
func (txn *TransactionWrapper) RegisterOnCommitCallback(callback TransactionOnCallback) {
	txn.onCommitCallbacks = append(txn.onCommitCallbacks, callback)
}

func (txn *TransactionWrapper) RegisterOnRollbackCallback(callback TransactionOnCallback) {
	txn.onRollbackCallbacks = append(txn.onRollbackCallbacks, callback)
}
