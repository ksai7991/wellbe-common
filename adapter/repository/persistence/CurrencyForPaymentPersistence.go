package persistence

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    log "wellbe-common/share/log"
    constants "wellbe-common/share/commonsettings/constants"

    _ "github.com/lib/pq"
    "database/sql"
    "context"
)

func (lp persistence) CreateCurrencyForPayment(ctx *context.Context, currencyForPayment *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.currency_for_payment(currency_cd, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            currencyForPayment.CurrencyCd,
                            currencyForPayment.CreateDatetime,
                            currencyForPayment.CreateFunction,
                            currencyForPayment.UpdateDatetime,
                            currencyForPayment.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return currencyForPayment, nil
}

func (lp persistence) UpdateCurrencyForPayment(ctx *context.Context, currencyForPayment *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.currency_for_payment "
    script = script + "SET create_datetime = $2, create_function = $3, update_datetime = $4, update_function = $5 "
    script = script + "WHERE currency_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            currencyForPayment.CurrencyCd,
                            currencyForPayment.CreateDatetime,
                            currencyForPayment.CreateFunction,
                            currencyForPayment.UpdateDatetime,
                            currencyForPayment.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return currencyForPayment, nil
}

func (lp persistence) DeleteCurrencyForPayment(ctx *context.Context, currencyCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.currency_for_payment WHERE currency_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, currencyCd); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return err
    }
    return nil
}

func (lp persistence) GetCurrencyForPaymentWithKey(ctx *context.Context, currencyCd int) ([]*model.CurrencyForPayment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var currencyForPayments []*model.CurrencyForPayment
    script := "SELECT currency_cd, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.currency_for_payment WHERE currency_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, currencyCd)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return rows, nil
    })
    rowsv, _ := rows.(*sql.Rows)
    if err != nil {
        logger.Error(err.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    defer rowsv.Close()
    for rowsv.Next() {
        currencyForPayment := &model.CurrencyForPayment{}
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&currencyForPayment.CurrencyCd, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if createDatetime.Valid {
            currencyForPayment.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            currencyForPayment.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            currencyForPayment.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            currencyForPayment.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        currencyForPayments = append(currencyForPayments, currencyForPayment)
    }

    return currencyForPayments, nil
}
