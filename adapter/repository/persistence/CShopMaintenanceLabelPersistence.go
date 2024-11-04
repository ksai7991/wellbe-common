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

func (lp persistence) CreateCShopMaintenanceLabel(ctx *context.Context, cShopMaintenanceLabel *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_shop_maintenance_label(shop_maintenance_label_cd, language_cd, shop_maintenance_label_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopMaintenanceLabel.ShopMaintenanceLabelCd,
                            cShopMaintenanceLabel.LanguageCd,
                            cShopMaintenanceLabel.ShopMaintenanceLabelName,
                            cShopMaintenanceLabel.CreateDatetime,
                            cShopMaintenanceLabel.CreateFunction,
                            cShopMaintenanceLabel.UpdateDatetime,
                            cShopMaintenanceLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopMaintenanceLabel, nil
}

func (lp persistence) UpdateCShopMaintenanceLabel(ctx *context.Context, cShopMaintenanceLabel *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_shop_maintenance_label "
    script = script + "SET shop_maintenance_label_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE shop_maintenance_label_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopMaintenanceLabel.ShopMaintenanceLabelCd,
                            cShopMaintenanceLabel.LanguageCd,
                            cShopMaintenanceLabel.ShopMaintenanceLabelName,
                            cShopMaintenanceLabel.CreateDatetime,
                            cShopMaintenanceLabel.CreateFunction,
                            cShopMaintenanceLabel.UpdateDatetime,
                            cShopMaintenanceLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopMaintenanceLabel, nil
}

func (lp persistence) DeleteCShopMaintenanceLabel(ctx *context.Context, shopMaintenanceLabelCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_shop_maintenance_label WHERE shop_maintenance_label_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, shopMaintenanceLabelCd, languageCd); err != nil {
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

func (lp persistence) GetCShopMaintenanceLabelWithKey(ctx *context.Context, shopMaintenanceLabelCd int,languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopMaintenanceLabels []*model.CShopMaintenanceLabel
    script := "SELECT shop_maintenance_label_cd, language_cd, shop_maintenance_label_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_maintenance_label WHERE shop_maintenance_label_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, shopMaintenanceLabelCd,languageCd)
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
        cShopMaintenanceLabel := &model.CShopMaintenanceLabel{}
        var shopMaintenanceLabelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopMaintenanceLabel.ShopMaintenanceLabelCd, &cShopMaintenanceLabel.LanguageCd, &shopMaintenanceLabelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopMaintenanceLabelName.Valid {
            cShopMaintenanceLabel.ShopMaintenanceLabelName = shopMaintenanceLabelName.String
        }
        if createDatetime.Valid {
            cShopMaintenanceLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopMaintenanceLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopMaintenanceLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopMaintenanceLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopMaintenanceLabels = append(cShopMaintenanceLabels, cShopMaintenanceLabel)
    }

    return cShopMaintenanceLabels, nil
}

func (lp persistence) GetCShopMaintenanceLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopMaintenanceLabels []*model.CShopMaintenanceLabel
    script := "SELECT shop_maintenance_label_cd, language_cd, shop_maintenance_label_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_maintenance_label WHERE language_cd = $1 ORDER BY create_datetime"
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
        cShopMaintenanceLabel := &model.CShopMaintenanceLabel{}
        var shopMaintenanceLabelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopMaintenanceLabel.ShopMaintenanceLabelCd, &cShopMaintenanceLabel.LanguageCd, &shopMaintenanceLabelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopMaintenanceLabelName.Valid {
            cShopMaintenanceLabel.ShopMaintenanceLabelName = shopMaintenanceLabelName.String
        }
        if createDatetime.Valid {
            cShopMaintenanceLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopMaintenanceLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopMaintenanceLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopMaintenanceLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopMaintenanceLabels = append(cShopMaintenanceLabels, cShopMaintenanceLabel)
    }

    return cShopMaintenanceLabels, nil
}
