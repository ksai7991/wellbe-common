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

func (lp persistence) CreateCShopStatus(ctx *context.Context, cShopStatus *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_shop_status(shop_status_cd, language_cd, shop_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopStatus.ShopStatusCd,
                            cShopStatus.LanguageCd,
                            cShopStatus.ShopStatusName,
                            cShopStatus.CreateDatetime,
                            cShopStatus.CreateFunction,
                            cShopStatus.UpdateDatetime,
                            cShopStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopStatus, nil
}

func (lp persistence) UpdateCShopStatus(ctx *context.Context, cShopStatus *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_shop_status "
    script = script + "SET shop_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE shop_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopStatus.ShopStatusCd,
                            cShopStatus.LanguageCd,
                            cShopStatus.ShopStatusName,
                            cShopStatus.CreateDatetime,
                            cShopStatus.CreateFunction,
                            cShopStatus.UpdateDatetime,
                            cShopStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopStatus, nil
}

func (lp persistence) DeleteCShopStatus(ctx *context.Context, shopStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_shop_status WHERE shop_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, shopStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCShopStatusWithKey(ctx *context.Context, shopStatusCd int,languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopStatuss []*model.CShopStatus
    script := "SELECT shop_status_cd, language_cd, shop_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_status WHERE shop_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, shopStatusCd,languageCd)
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
        cShopStatus := &model.CShopStatus{}
        var shopStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopStatus.ShopStatusCd, &cShopStatus.LanguageCd, &shopStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopStatusName.Valid {
            cShopStatus.ShopStatusName = shopStatusName.String
        }
        if createDatetime.Valid {
            cShopStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopStatuss = append(cShopStatuss, cShopStatus)
    }

    return cShopStatuss, nil
}

func (lp persistence) GetCShopStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopStatuss []*model.CShopStatus
    script := "SELECT shop_status_cd, language_cd, shop_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cShopStatus := &model.CShopStatus{}
        var shopStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopStatus.ShopStatusCd, &cShopStatus.LanguageCd, &shopStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopStatusName.Valid {
            cShopStatus.ShopStatusName = shopStatusName.String
        }
        if createDatetime.Valid {
            cShopStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopStatuss = append(cShopStatuss, cShopStatus)
    }

    return cShopStatuss, nil
}
