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

func (lp persistence) CreateCShopContractPlanItem(ctx *context.Context, cShopContractPlanItem *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_shop_contract_plan_item(shop_contract_plan_item_cd, language_cd, shop_contract_plan_name, unit, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopContractPlanItem.ShopContractPlanItemCd,
                            cShopContractPlanItem.LanguageCd,
                            cShopContractPlanItem.ShopContractPlanName,
                            cShopContractPlanItem.Unit,
                            cShopContractPlanItem.CreateDatetime,
                            cShopContractPlanItem.CreateFunction,
                            cShopContractPlanItem.UpdateDatetime,
                            cShopContractPlanItem.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopContractPlanItem, nil
}

func (lp persistence) UpdateCShopContractPlanItem(ctx *context.Context, cShopContractPlanItem *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_shop_contract_plan_item "
    script = script + "SET shop_contract_plan_name = $3, unit = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE shop_contract_plan_item_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopContractPlanItem.ShopContractPlanItemCd,
                            cShopContractPlanItem.LanguageCd,
                            cShopContractPlanItem.ShopContractPlanName,
                            cShopContractPlanItem.Unit,
                            cShopContractPlanItem.CreateDatetime,
                            cShopContractPlanItem.CreateFunction,
                            cShopContractPlanItem.UpdateDatetime,
                            cShopContractPlanItem.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopContractPlanItem, nil
}

func (lp persistence) DeleteCShopContractPlanItem(ctx *context.Context, shopContractPlanItemCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_shop_contract_plan_item WHERE shop_contract_plan_item_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, shopContractPlanItemCd, languageCd); err != nil {
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

func (lp persistence) GetCShopContractPlanItemWithKey(ctx *context.Context, shopContractPlanItemCd int,languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopContractPlanItems []*model.CShopContractPlanItem
    script := "SELECT shop_contract_plan_item_cd, language_cd, shop_contract_plan_name, unit, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_contract_plan_item WHERE shop_contract_plan_item_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, shopContractPlanItemCd,languageCd)
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
        cShopContractPlanItem := &model.CShopContractPlanItem{}
        var shopContractPlanName sql.NullString
        var unit sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopContractPlanItem.ShopContractPlanItemCd, &cShopContractPlanItem.LanguageCd, &shopContractPlanName, &unit, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopContractPlanName.Valid {
            cShopContractPlanItem.ShopContractPlanName = shopContractPlanName.String
        }
        if unit.Valid {
            cShopContractPlanItem.Unit = unit.String
        }
        if createDatetime.Valid {
            cShopContractPlanItem.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopContractPlanItem.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopContractPlanItem.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopContractPlanItem.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopContractPlanItems = append(cShopContractPlanItems, cShopContractPlanItem)
    }

    return cShopContractPlanItems, nil
}

func (lp persistence) GetCShopContractPlanItemWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopContractPlanItems []*model.CShopContractPlanItem
    script := "SELECT shop_contract_plan_item_cd, language_cd, shop_contract_plan_name, unit, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_contract_plan_item WHERE language_cd = $1 ORDER BY create_datetime"
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
        cShopContractPlanItem := &model.CShopContractPlanItem{}
        var shopContractPlanName sql.NullString
        var unit sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopContractPlanItem.ShopContractPlanItemCd, &cShopContractPlanItem.LanguageCd, &shopContractPlanName, &unit, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopContractPlanName.Valid {
            cShopContractPlanItem.ShopContractPlanName = shopContractPlanName.String
        }
        if unit.Valid {
            cShopContractPlanItem.Unit = unit.String
        }
        if createDatetime.Valid {
            cShopContractPlanItem.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopContractPlanItem.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopContractPlanItem.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopContractPlanItem.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopContractPlanItems = append(cShopContractPlanItems, cShopContractPlanItem)
    }

    return cShopContractPlanItems, nil
}
