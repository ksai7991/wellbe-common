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

func (lp persistence) CreateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.currency_exchange_rate(base_currency_cd, target_currency_cd, paire_name, rate, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            currencyExchangeRate.BaseCurrencyCd,
                            currencyExchangeRate.TargetCurrencyCd,
                            currencyExchangeRate.PaireName,
                            currencyExchangeRate.Rate,
                            currencyExchangeRate.CreateDatetime,
                            currencyExchangeRate.CreateFunction,
                            currencyExchangeRate.UpdateDatetime,
                            currencyExchangeRate.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return currencyExchangeRate, nil
}

func (lp persistence) UpdateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.currency_exchange_rate "
    script = script + "SET paire_name = $3, rate = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE base_currency_cd = $1 and target_currency_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            currencyExchangeRate.BaseCurrencyCd,
                            currencyExchangeRate.TargetCurrencyCd,
                            currencyExchangeRate.PaireName,
                            currencyExchangeRate.Rate,
                            currencyExchangeRate.CreateDatetime,
                            currencyExchangeRate.CreateFunction,
                            currencyExchangeRate.UpdateDatetime,
                            currencyExchangeRate.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return currencyExchangeRate, nil
}

func (lp persistence) DeleteCurrencyExchangeRate(ctx *context.Context, baseCurrencyCd int, targetCurrencyCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.currency_exchange_rate WHERE base_currency_cd = $1 and target_currency_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, baseCurrencyCd, targetCurrencyCd); err != nil {
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

func (lp persistence) GetCurrencyExchangeRateWithKey(ctx *context.Context, baseCurrencyCd int,targetCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var currencyExchangeRates []*model.CurrencyExchangeRate
    script := "SELECT base_currency_cd, target_currency_cd, paire_name, rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.currency_exchange_rate WHERE base_currency_cd = $1 and target_currency_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, baseCurrencyCd,targetCurrencyCd)
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
        currencyExchangeRate := &model.CurrencyExchangeRate{}
        var paireName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&currencyExchangeRate.BaseCurrencyCd, &currencyExchangeRate.TargetCurrencyCd, &paireName, &currencyExchangeRate.Rate, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if paireName.Valid {
            currencyExchangeRate.PaireName = paireName.String
        }
        if createDatetime.Valid {
            currencyExchangeRate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            currencyExchangeRate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            currencyExchangeRate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            currencyExchangeRate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        currencyExchangeRates = append(currencyExchangeRates, currencyExchangeRate)
    }

    return currencyExchangeRates, nil
}

func (lp persistence) GetCurrencyExchangeRateWithPaireName(ctx *context.Context, paireName string) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var currencyExchangeRates []*model.CurrencyExchangeRate
    script := "SELECT base_currency_cd, target_currency_cd, paire_name, rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.currency_exchange_rate WHERE paire_name = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, paireName)
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
        currencyExchangeRate := &model.CurrencyExchangeRate{}
        var paireName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&currencyExchangeRate.BaseCurrencyCd, &currencyExchangeRate.TargetCurrencyCd, &paireName, &currencyExchangeRate.Rate, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if paireName.Valid {
            currencyExchangeRate.PaireName = paireName.String
        }
        if createDatetime.Valid {
            currencyExchangeRate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            currencyExchangeRate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            currencyExchangeRate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            currencyExchangeRate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        currencyExchangeRates = append(currencyExchangeRates, currencyExchangeRate)
    }

    return currencyExchangeRates, nil
}

func (lp persistence) GetCurrencyExchangeRateWithBase(ctx *context.Context, baseCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var currencyExchangeRates []*model.CurrencyExchangeRate
    script := "SELECT base_currency_cd, target_currency_cd, paire_name, rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.currency_exchange_rate WHERE base_currency_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, baseCurrencyCd)
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
        currencyExchangeRate := &model.CurrencyExchangeRate{}
        var paireName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&currencyExchangeRate.BaseCurrencyCd, &currencyExchangeRate.TargetCurrencyCd, &paireName, &currencyExchangeRate.Rate, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if paireName.Valid {
            currencyExchangeRate.PaireName = paireName.String
        }
        if createDatetime.Valid {
            currencyExchangeRate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            currencyExchangeRate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            currencyExchangeRate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            currencyExchangeRate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        currencyExchangeRates = append(currencyExchangeRates, currencyExchangeRate)
    }

    return currencyExchangeRates, nil
}
