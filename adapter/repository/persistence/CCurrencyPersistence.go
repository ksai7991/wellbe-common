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

func (lp persistence) CreateCCurrency(ctx *context.Context, cCurrency *model.CCurrency) (*model.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_currency(currency_cd, language_cd, currency_name, currency_cd_iso, significant_digit, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCurrency.CurrencyCd,
                            cCurrency.LanguageCd,
                            cCurrency.CurrencyName,
                            cCurrency.CurrencyCdIso,
                            cCurrency.SignificantDigit,
                            cCurrency.CreateDatetime,
                            cCurrency.CreateFunction,
                            cCurrency.UpdateDatetime,
                            cCurrency.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCurrency, nil
}

func (lp persistence) UpdateCCurrency(ctx *context.Context, cCurrency *model.CCurrency) (*model.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_currency "
    script = script + "SET currency_name = $3, currency_cd_iso = $4, significant_digit = $5, create_datetime = $6, create_function = $7, update_datetime = $8, update_function = $9 "
    script = script + "WHERE currency_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCurrency.CurrencyCd,
                            cCurrency.LanguageCd,
                            cCurrency.CurrencyName,
                            cCurrency.CurrencyCdIso,
                            cCurrency.SignificantDigit,
                            cCurrency.CreateDatetime,
                            cCurrency.CreateFunction,
                            cCurrency.UpdateDatetime,
                            cCurrency.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCurrency, nil
}

func (lp persistence) DeleteCCurrency(ctx *context.Context, currencyCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_currency WHERE currency_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, currencyCd, languageCd); err != nil {
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

func (lp persistence) GetCCurrencyWithKey(ctx *context.Context, currencyCd int,languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCurrencys []*model.CCurrency
    script := "SELECT currency_cd, language_cd, currency_name, currency_cd_iso, significant_digit, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_currency WHERE currency_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, currencyCd,languageCd)
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
        cCurrency := &model.CCurrency{}
        var currencyName sql.NullString
        var currencyCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCurrency.CurrencyCd, &cCurrency.LanguageCd, &currencyName, &currencyCdIso, &cCurrency.SignificantDigit, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if currencyName.Valid {
            cCurrency.CurrencyName = currencyName.String
        }
        if currencyCdIso.Valid {
            cCurrency.CurrencyCdIso = currencyCdIso.String
        }
        if createDatetime.Valid {
            cCurrency.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCurrency.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCurrency.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCurrency.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCurrencys = append(cCurrencys, cCurrency)
    }

    return cCurrencys, nil
}

func (lp persistence) GetCCurrencyWithCurrencyCdIso(ctx *context.Context, currencyCdIso string,significantDigit int) ([]*model.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCurrencys []*model.CCurrency
    script := "SELECT currency_cd, language_cd, currency_name, currency_cd_iso, significant_digit, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_currency WHERE currency_cd_iso = $1 and significant_digit = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, currencyCdIso,significantDigit)
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
        cCurrency := &model.CCurrency{}
        var currencyName sql.NullString
        var currencyCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCurrency.CurrencyCd, &cCurrency.LanguageCd, &currencyName, &currencyCdIso, &cCurrency.SignificantDigit, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if currencyName.Valid {
            cCurrency.CurrencyName = currencyName.String
        }
        if currencyCdIso.Valid {
            cCurrency.CurrencyCdIso = currencyCdIso.String
        }
        if createDatetime.Valid {
            cCurrency.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCurrency.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCurrency.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCurrency.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCurrencys = append(cCurrencys, cCurrency)
    }

    return cCurrencys, nil
}

func (lp persistence) GetCCurrencyWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCurrencys []*model.CCurrency
    script := "SELECT currency_cd, language_cd, currency_name, currency_cd_iso, significant_digit, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_currency WHERE language_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCd)
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
        cCurrency := &model.CCurrency{}
        var currencyName sql.NullString
        var currencyCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCurrency.CurrencyCd, &cCurrency.LanguageCd, &currencyName, &currencyCdIso, &cCurrency.SignificantDigit, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if currencyName.Valid {
            cCurrency.CurrencyName = currencyName.String
        }
        if currencyCdIso.Valid {
            cCurrency.CurrencyCdIso = currencyCdIso.String
        }
        if createDatetime.Valid {
            cCurrency.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCurrency.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCurrency.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCurrency.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCurrencys = append(cCurrencys, cCurrency)
    }

    return cCurrencys, nil
}
