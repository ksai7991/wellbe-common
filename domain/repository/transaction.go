package repository

import (
	"context"
	"database/sql"
	errordef "wellbe-common/share/errordef"
)

type Transaction interface {
	DoInTx(*context.Context, func(tx *sql.Tx) (interface{}, *errordef.LogicError)) (interface{}, *errordef.LogicError)
	Rollback(*context.Context)
	Commit(*context.Context) *errordef.LogicError
}