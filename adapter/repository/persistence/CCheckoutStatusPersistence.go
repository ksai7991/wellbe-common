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

func (lp persistence) CreateCCheckoutStatus(ctx *context.Context, cCheckoutStatus *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_checkout_status(checkout_status_cd, language_cd, checkout_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCheckoutStatus.CheckoutStatusCd,
                            cCheckoutStatus.LanguageCd,
                            cCheckoutStatus.CheckoutStatusName,
                            cCheckoutStatus.CreateDatetime,
                            cCheckoutStatus.CreateFunction,
                            cCheckoutStatus.UpdateDatetime,
                            cCheckoutStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCheckoutStatus, nil
}

func (lp persistence) UpdateCCheckoutStatus(ctx *context.Context, cCheckoutStatus *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_checkout_status "
    script = script + "SET checkout_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE checkout_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCheckoutStatus.CheckoutStatusCd,
                            cCheckoutStatus.LanguageCd,
                            cCheckoutStatus.CheckoutStatusName,
                            cCheckoutStatus.CreateDatetime,
                            cCheckoutStatus.CreateFunction,
                            cCheckoutStatus.UpdateDatetime,
                            cCheckoutStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCheckoutStatus, nil
}

func (lp persistence) DeleteCCheckoutStatus(ctx *context.Context, checkoutStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_checkout_status WHERE checkout_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, checkoutStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCCheckoutStatusWithKey(ctx *context.Context, checkoutStatusCd int,languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCheckoutStatuss []*model.CCheckoutStatus
    script := "SELECT checkout_status_cd, language_cd, checkout_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_checkout_status WHERE checkout_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, checkoutStatusCd,languageCd)
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
        cCheckoutStatus := &model.CCheckoutStatus{}
        var checkoutStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCheckoutStatus.CheckoutStatusCd, &cCheckoutStatus.LanguageCd, &checkoutStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if checkoutStatusName.Valid {
            cCheckoutStatus.CheckoutStatusName = checkoutStatusName.String
        }
        if createDatetime.Valid {
            cCheckoutStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCheckoutStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCheckoutStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCheckoutStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCheckoutStatuss = append(cCheckoutStatuss, cCheckoutStatus)
    }

    return cCheckoutStatuss, nil
}

func (lp persistence) GetCCheckoutStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCheckoutStatuss []*model.CCheckoutStatus
    script := "SELECT checkout_status_cd, language_cd, checkout_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_checkout_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cCheckoutStatus := &model.CCheckoutStatus{}
        var checkoutStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCheckoutStatus.CheckoutStatusCd, &cCheckoutStatus.LanguageCd, &checkoutStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if checkoutStatusName.Valid {
            cCheckoutStatus.CheckoutStatusName = checkoutStatusName.String
        }
        if createDatetime.Valid {
            cCheckoutStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCheckoutStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCheckoutStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCheckoutStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCheckoutStatuss = append(cCheckoutStatuss, cCheckoutStatus)
    }

    return cCheckoutStatuss, nil
}
