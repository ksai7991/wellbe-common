package repository

import (
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
	constants "wellbe-common/share/commonsettings/constants"
	repository "wellbe-common/domain/repository"
	"context"
	"database/sql"
)

var txKey = struct{}{}

type transaction struct{
	db *sql.DB
} 

func NewTransaction(db *sql.DB) repository.Transaction {
	return &transaction{db: db}
}

func (tr transaction) DoInTx(ctx *context.Context, f func(tx *sql.Tx) (interface{}, *errordef.LogicError)) (interface{}, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	tx, ok := tr.getTx(ctx)
	if !ok {
		txBegin, err := tr.db.Begin()
		if  err != nil {
			logger.Error(err.Error())
			return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
		}
		tx = txBegin
	}
	*ctx = context.WithValue(*ctx, txKey, tx)
	v, err := f(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}

func (tr transaction) Rollback(ctx *context.Context) {
	tx, ok := tr.getTx(ctx)
	if !ok {
		return
	}
	tx.Rollback()
}

func (tr transaction) Commit(ctx *context.Context) *errordef.LogicError {
	logger := log.GetLogger()
	defer logger.Sync()
	tx, ok := tr.getTx(ctx)
	if !ok {
		return nil
	}
	if err := tx.Commit(); err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}
	return nil
}

func (tr transaction) getTx(ctx *context.Context) (*sql.Tx, bool) {
	logger := log.GetLogger()
	defer logger.Sync()
	tx, ok := (*ctx).Value(txKey).(*sql.Tx)
	return tx, ok
}