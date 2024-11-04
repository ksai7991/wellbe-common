package query

import (
	entity "wellbe-common/adapter/repository/entity"
	querySql "wellbe-common/adapter/repository/sql"
	"wellbe-common/domain/model"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func (q Query) QueryCurrencyForPayment(ctx *context.Context, language_cd string) ([]*entity.CurrencyForPayment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cs []*entity.CurrencyForPayment
    rows, err := q.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(querySql.GetCurrencyForPayment, language_cd)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return rows, nil
    })
    rowsv, _ := rows.(*sql.Rows)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    defer rowsv.Close()
    for rowsv.Next() {
        currency := &model.CCurrency{}
        c := &entity.CurrencyForPayment{}
        err := rowsv.Scan(&currency.CurrencyCd, &currency.LanguageCd, &currency.CurrencyName, &currency.CurrencyCdIso)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        c.CCurrency = *currency
        cs = append(cs, c)
    }

    return cs, nil
}