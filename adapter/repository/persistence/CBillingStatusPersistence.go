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

func (lp persistence) CreateCBillingStatus(ctx *context.Context, cBillingStatus *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_billing_status(billing_status_cd, language_cd, billing_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBillingStatus.BillingStatusCd,
                            cBillingStatus.LanguageCd,
                            cBillingStatus.BillingStatusName,
                            cBillingStatus.CreateDatetime,
                            cBillingStatus.CreateFunction,
                            cBillingStatus.UpdateDatetime,
                            cBillingStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBillingStatus, nil
}

func (lp persistence) UpdateCBillingStatus(ctx *context.Context, cBillingStatus *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_billing_status "
    script = script + "SET billing_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE billing_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBillingStatus.BillingStatusCd,
                            cBillingStatus.LanguageCd,
                            cBillingStatus.BillingStatusName,
                            cBillingStatus.CreateDatetime,
                            cBillingStatus.CreateFunction,
                            cBillingStatus.UpdateDatetime,
                            cBillingStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBillingStatus, nil
}

func (lp persistence) DeleteCBillingStatus(ctx *context.Context, billingStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_billing_status WHERE billing_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, billingStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCBillingStatusWithKey(ctx *context.Context, billingStatusCd int,languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBillingStatuss []*model.CBillingStatus
    script := "SELECT billing_status_cd, language_cd, billing_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_billing_status WHERE billing_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, billingStatusCd,languageCd)
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
        cBillingStatus := &model.CBillingStatus{}
        var billingStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBillingStatus.BillingStatusCd, &cBillingStatus.LanguageCd, &billingStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if billingStatusName.Valid {
            cBillingStatus.BillingStatusName = billingStatusName.String
        }
        if createDatetime.Valid {
            cBillingStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBillingStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBillingStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBillingStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBillingStatuss = append(cBillingStatuss, cBillingStatus)
    }

    return cBillingStatuss, nil
}

func (lp persistence) GetCBillingStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBillingStatuss []*model.CBillingStatus
    script := "SELECT billing_status_cd, language_cd, billing_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_billing_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cBillingStatus := &model.CBillingStatus{}
        var billingStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBillingStatus.BillingStatusCd, &cBillingStatus.LanguageCd, &billingStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if billingStatusName.Valid {
            cBillingStatus.BillingStatusName = billingStatusName.String
        }
        if createDatetime.Valid {
            cBillingStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBillingStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBillingStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBillingStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBillingStatuss = append(cBillingStatuss, cBillingStatus)
    }

    return cBillingStatuss, nil
}
