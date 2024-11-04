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

func (lp persistence) CreateCCheckoutMethod(ctx *context.Context, cCheckoutMethod *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_checkout_method(checkout_method_cd, language_cd, checkout_method_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCheckoutMethod.CheckoutMethodCd,
                            cCheckoutMethod.LanguageCd,
                            cCheckoutMethod.CheckoutMethodName,
                            cCheckoutMethod.CreateDatetime,
                            cCheckoutMethod.CreateFunction,
                            cCheckoutMethod.UpdateDatetime,
                            cCheckoutMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCheckoutMethod, nil
}

func (lp persistence) UpdateCCheckoutMethod(ctx *context.Context, cCheckoutMethod *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_checkout_method "
    script = script + "SET checkout_method_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE checkout_method_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCheckoutMethod.CheckoutMethodCd,
                            cCheckoutMethod.LanguageCd,
                            cCheckoutMethod.CheckoutMethodName,
                            cCheckoutMethod.CreateDatetime,
                            cCheckoutMethod.CreateFunction,
                            cCheckoutMethod.UpdateDatetime,
                            cCheckoutMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCheckoutMethod, nil
}

func (lp persistence) DeleteCCheckoutMethod(ctx *context.Context, checkoutMethodCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_checkout_method WHERE checkout_method_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, checkoutMethodCd, languageCd); err != nil {
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

func (lp persistence) GetCCheckoutMethodWithKey(ctx *context.Context, checkoutMethodCd int,languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCheckoutMethods []*model.CCheckoutMethod
    script := "SELECT checkout_method_cd, language_cd, checkout_method_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_checkout_method WHERE checkout_method_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, checkoutMethodCd,languageCd)
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
        cCheckoutMethod := &model.CCheckoutMethod{}
        var checkoutMethodName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCheckoutMethod.CheckoutMethodCd, &cCheckoutMethod.LanguageCd, &checkoutMethodName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if checkoutMethodName.Valid {
            cCheckoutMethod.CheckoutMethodName = checkoutMethodName.String
        }
        if createDatetime.Valid {
            cCheckoutMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCheckoutMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCheckoutMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCheckoutMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCheckoutMethods = append(cCheckoutMethods, cCheckoutMethod)
    }

    return cCheckoutMethods, nil
}

func (lp persistence) GetCCheckoutMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCheckoutMethods []*model.CCheckoutMethod
    script := "SELECT checkout_method_cd, language_cd, checkout_method_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_checkout_method WHERE language_cd = $1 ORDER BY create_datetime"
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
        cCheckoutMethod := &model.CCheckoutMethod{}
        var checkoutMethodName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCheckoutMethod.CheckoutMethodCd, &cCheckoutMethod.LanguageCd, &checkoutMethodName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if checkoutMethodName.Valid {
            cCheckoutMethod.CheckoutMethodName = checkoutMethodName.String
        }
        if createDatetime.Valid {
            cCheckoutMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCheckoutMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCheckoutMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCheckoutMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCheckoutMethods = append(cCheckoutMethods, cCheckoutMethod)
    }

    return cCheckoutMethods, nil
}
