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

func (lp persistence) CreateCShopPaymentMethod(ctx *context.Context, cShopPaymentMethod *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_shop_payment_method(shop_payment_method_cd, language_cd, shop_payment_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopPaymentMethod.ShopPaymentMethodCd,
                            cShopPaymentMethod.LanguageCd,
                            cShopPaymentMethod.ShopPaymentName,
                            cShopPaymentMethod.CreateDatetime,
                            cShopPaymentMethod.CreateFunction,
                            cShopPaymentMethod.UpdateDatetime,
                            cShopPaymentMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopPaymentMethod, nil
}

func (lp persistence) UpdateCShopPaymentMethod(ctx *context.Context, cShopPaymentMethod *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_shop_payment_method "
    script = script + "SET language_cd = $2, shop_payment_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE shop_payment_method_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopPaymentMethod.ShopPaymentMethodCd,
                            cShopPaymentMethod.LanguageCd,
                            cShopPaymentMethod.ShopPaymentName,
                            cShopPaymentMethod.CreateDatetime,
                            cShopPaymentMethod.CreateFunction,
                            cShopPaymentMethod.UpdateDatetime,
                            cShopPaymentMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopPaymentMethod, nil
}

func (lp persistence) DeleteCShopPaymentMethod(ctx *context.Context, shopPaymentMethodCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_shop_payment_method WHERE shop_payment_method_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, shopPaymentMethodCd); err != nil {
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

func (lp persistence) GetCShopPaymentMethodWithKey(ctx *context.Context, shopPaymentMethodCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopPaymentMethods []*model.CShopPaymentMethod
    script := "SELECT shop_payment_method_cd, language_cd, shop_payment_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_payment_method WHERE shop_payment_method_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, shopPaymentMethodCd)
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
        cShopPaymentMethod := &model.CShopPaymentMethod{}
        var shopPaymentName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopPaymentMethod.ShopPaymentMethodCd, &cShopPaymentMethod.LanguageCd, &shopPaymentName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopPaymentName.Valid {
            cShopPaymentMethod.ShopPaymentName = shopPaymentName.String
        }
        if createDatetime.Valid {
            cShopPaymentMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopPaymentMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopPaymentMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopPaymentMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopPaymentMethods = append(cShopPaymentMethods, cShopPaymentMethod)
    }

    return cShopPaymentMethods, nil
}

func (lp persistence) GetCShopPaymentMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopPaymentMethods []*model.CShopPaymentMethod
    script := "SELECT shop_payment_method_cd, language_cd, shop_payment_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_payment_method WHERE language_cd = $1 ORDER BY create_datetime"
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
        cShopPaymentMethod := &model.CShopPaymentMethod{}
        var shopPaymentName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopPaymentMethod.ShopPaymentMethodCd, &cShopPaymentMethod.LanguageCd, &shopPaymentName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopPaymentName.Valid {
            cShopPaymentMethod.ShopPaymentName = shopPaymentName.String
        }
        if createDatetime.Valid {
            cShopPaymentMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopPaymentMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopPaymentMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopPaymentMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopPaymentMethods = append(cShopPaymentMethods, cShopPaymentMethod)
    }

    return cShopPaymentMethods, nil
}
