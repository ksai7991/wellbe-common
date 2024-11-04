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

func (lp persistence) CreateCCheckoutTiming(ctx *context.Context, cCheckoutTiming *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_checkout_timing(checkout_timing_cd, language_cd, checkout_timing_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCheckoutTiming.CheckoutTimingCd,
                            cCheckoutTiming.LanguageCd,
                            cCheckoutTiming.CheckoutTimingName,
                            cCheckoutTiming.CreateDatetime,
                            cCheckoutTiming.CreateFunction,
                            cCheckoutTiming.UpdateDatetime,
                            cCheckoutTiming.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCheckoutTiming, nil
}

func (lp persistence) UpdateCCheckoutTiming(ctx *context.Context, cCheckoutTiming *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_checkout_timing "
    script = script + "SET checkout_timing_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE checkout_timing_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCheckoutTiming.CheckoutTimingCd,
                            cCheckoutTiming.LanguageCd,
                            cCheckoutTiming.CheckoutTimingName,
                            cCheckoutTiming.CreateDatetime,
                            cCheckoutTiming.CreateFunction,
                            cCheckoutTiming.UpdateDatetime,
                            cCheckoutTiming.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCheckoutTiming, nil
}

func (lp persistence) DeleteCCheckoutTiming(ctx *context.Context, checkoutTimingCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_checkout_timing WHERE checkout_timing_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, checkoutTimingCd, languageCd); err != nil {
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

func (lp persistence) GetCCheckoutTimingWithKey(ctx *context.Context, checkoutTimingCd int,languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCheckoutTimings []*model.CCheckoutTiming
    script := "SELECT checkout_timing_cd, language_cd, checkout_timing_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_checkout_timing WHERE checkout_timing_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, checkoutTimingCd,languageCd)
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
        cCheckoutTiming := &model.CCheckoutTiming{}
        var checkoutTimingName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCheckoutTiming.CheckoutTimingCd, &cCheckoutTiming.LanguageCd, &checkoutTimingName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if checkoutTimingName.Valid {
            cCheckoutTiming.CheckoutTimingName = checkoutTimingName.String
        }
        if createDatetime.Valid {
            cCheckoutTiming.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCheckoutTiming.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCheckoutTiming.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCheckoutTiming.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCheckoutTimings = append(cCheckoutTimings, cCheckoutTiming)
    }

    return cCheckoutTimings, nil
}

func (lp persistence) GetCCheckoutTimingWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCheckoutTimings []*model.CCheckoutTiming
    script := "SELECT checkout_timing_cd, language_cd, checkout_timing_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_checkout_timing WHERE language_cd = $1 ORDER BY create_datetime"
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
        cCheckoutTiming := &model.CCheckoutTiming{}
        var checkoutTimingName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCheckoutTiming.CheckoutTimingCd, &cCheckoutTiming.LanguageCd, &checkoutTimingName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if checkoutTimingName.Valid {
            cCheckoutTiming.CheckoutTimingName = checkoutTimingName.String
        }
        if createDatetime.Valid {
            cCheckoutTiming.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCheckoutTiming.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCheckoutTiming.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCheckoutTiming.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCheckoutTimings = append(cCheckoutTimings, cCheckoutTiming)
    }

    return cCheckoutTimings, nil
}
