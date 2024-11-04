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

func (lp persistence) CreateCShopEquipment(ctx *context.Context, cShopEquipment *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_shop_equipment(shop_equipment_cd, language_cd, shop_equipment_name, unit_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopEquipment.ShopEquipmentCd,
                            cShopEquipment.LanguageCd,
                            cShopEquipment.ShopEquipmentName,
                            cShopEquipment.UnitName,
                            cShopEquipment.CreateDatetime,
                            cShopEquipment.CreateFunction,
                            cShopEquipment.UpdateDatetime,
                            cShopEquipment.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopEquipment, nil
}

func (lp persistence) UpdateCShopEquipment(ctx *context.Context, cShopEquipment *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_shop_equipment "
    script = script + "SET shop_equipment_name = $3, unit_name = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE shop_equipment_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopEquipment.ShopEquipmentCd,
                            cShopEquipment.LanguageCd,
                            cShopEquipment.ShopEquipmentName,
                            cShopEquipment.UnitName,
                            cShopEquipment.CreateDatetime,
                            cShopEquipment.CreateFunction,
                            cShopEquipment.UpdateDatetime,
                            cShopEquipment.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopEquipment, nil
}

func (lp persistence) DeleteCShopEquipment(ctx *context.Context, shopEquipmentCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_shop_equipment WHERE shop_equipment_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, shopEquipmentCd, languageCd); err != nil {
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

func (lp persistence) GetCShopEquipmentWithKey(ctx *context.Context, shopEquipmentCd int,languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopEquipments []*model.CShopEquipment
    script := "SELECT shop_equipment_cd, language_cd, shop_equipment_name, unit_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_equipment WHERE shop_equipment_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, shopEquipmentCd,languageCd)
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
        cShopEquipment := &model.CShopEquipment{}
        var shopEquipmentName sql.NullString
        var unitName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopEquipment.ShopEquipmentCd, &cShopEquipment.LanguageCd, &shopEquipmentName, &unitName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopEquipmentName.Valid {
            cShopEquipment.ShopEquipmentName = shopEquipmentName.String
        }
        if unitName.Valid {
            cShopEquipment.UnitName = unitName.String
        }
        if createDatetime.Valid {
            cShopEquipment.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopEquipment.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopEquipment.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopEquipment.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopEquipments = append(cShopEquipments, cShopEquipment)
    }

    return cShopEquipments, nil
}

func (lp persistence) GetCShopEquipmentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopEquipments []*model.CShopEquipment
    script := "SELECT shop_equipment_cd, language_cd, shop_equipment_name, unit_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_equipment WHERE language_cd = $1 ORDER BY create_datetime"
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
        cShopEquipment := &model.CShopEquipment{}
        var shopEquipmentName sql.NullString
        var unitName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopEquipment.ShopEquipmentCd, &cShopEquipment.LanguageCd, &shopEquipmentName, &unitName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopEquipmentName.Valid {
            cShopEquipment.ShopEquipmentName = shopEquipmentName.String
        }
        if unitName.Valid {
            cShopEquipment.UnitName = unitName.String
        }
        if createDatetime.Valid {
            cShopEquipment.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopEquipment.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopEquipment.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopEquipment.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopEquipments = append(cShopEquipments, cShopEquipment)
    }

    return cShopEquipments, nil
}
