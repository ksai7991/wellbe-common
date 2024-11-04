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

func (lp persistence) CreateCPayoutMethod(ctx *context.Context, cPayoutMethod *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_payout_method(payout_method_cd, language_cd, payout_method_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cPayoutMethod.PayoutMethodCd,
                            cPayoutMethod.LanguageCd,
                            cPayoutMethod.PayoutMethodName,
                            cPayoutMethod.CreateDatetime,
                            cPayoutMethod.CreateFunction,
                            cPayoutMethod.UpdateDatetime,
                            cPayoutMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cPayoutMethod, nil
}

func (lp persistence) UpdateCPayoutMethod(ctx *context.Context, cPayoutMethod *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_payout_method "
    script = script + "SET language_cd = $2, payout_method_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE payout_method_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cPayoutMethod.PayoutMethodCd,
                            cPayoutMethod.LanguageCd,
                            cPayoutMethod.PayoutMethodName,
                            cPayoutMethod.CreateDatetime,
                            cPayoutMethod.CreateFunction,
                            cPayoutMethod.UpdateDatetime,
                            cPayoutMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cPayoutMethod, nil
}

func (lp persistence) DeleteCPayoutMethod(ctx *context.Context, payoutMethodCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_payout_method WHERE payout_method_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, payoutMethodCd); err != nil {
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

func (lp persistence) GetCPayoutMethodWithKey(ctx *context.Context, payoutMethodCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cPayoutMethods []*model.CPayoutMethod
    script := "SELECT payout_method_cd, language_cd, payout_method_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_payout_method WHERE payout_method_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, payoutMethodCd)
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
        cPayoutMethod := &model.CPayoutMethod{}
        var payoutMethodName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cPayoutMethod.PayoutMethodCd, &cPayoutMethod.LanguageCd, &payoutMethodName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if payoutMethodName.Valid {
            cPayoutMethod.PayoutMethodName = payoutMethodName.String
        }
        if createDatetime.Valid {
            cPayoutMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cPayoutMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cPayoutMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cPayoutMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cPayoutMethods = append(cPayoutMethods, cPayoutMethod)
    }

    return cPayoutMethods, nil
}

func (lp persistence) GetCPayoutMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cPayoutMethods []*model.CPayoutMethod
    script := "SELECT payout_method_cd, language_cd, payout_method_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_payout_method WHERE language_cd = $1 ORDER BY create_datetime"
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
        cPayoutMethod := &model.CPayoutMethod{}
        var payoutMethodName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cPayoutMethod.PayoutMethodCd, &cPayoutMethod.LanguageCd, &payoutMethodName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if payoutMethodName.Valid {
            cPayoutMethod.PayoutMethodName = payoutMethodName.String
        }
        if createDatetime.Valid {
            cPayoutMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cPayoutMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cPayoutMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cPayoutMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cPayoutMethods = append(cPayoutMethods, cPayoutMethod)
    }

    return cPayoutMethods, nil
}
