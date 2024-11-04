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

func (lp persistence) CreateCShopImageFilterCategory(ctx *context.Context, cShopImageFilterCategory *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_shop_image_filter_category(shop_image_filter_category_cd, language_cd, shop_image_filter_category_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopImageFilterCategory.ShopImageFilterCategoryCd,
                            cShopImageFilterCategory.LanguageCd,
                            cShopImageFilterCategory.ShopImageFilterCategoryName,
                            cShopImageFilterCategory.CreateDatetime,
                            cShopImageFilterCategory.CreateFunction,
                            cShopImageFilterCategory.UpdateDatetime,
                            cShopImageFilterCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopImageFilterCategory, nil
}

func (lp persistence) UpdateCShopImageFilterCategory(ctx *context.Context, cShopImageFilterCategory *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_shop_image_filter_category "
    script = script + "SET shop_image_filter_category_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE shop_image_filter_category_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cShopImageFilterCategory.ShopImageFilterCategoryCd,
                            cShopImageFilterCategory.LanguageCd,
                            cShopImageFilterCategory.ShopImageFilterCategoryName,
                            cShopImageFilterCategory.CreateDatetime,
                            cShopImageFilterCategory.CreateFunction,
                            cShopImageFilterCategory.UpdateDatetime,
                            cShopImageFilterCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cShopImageFilterCategory, nil
}

func (lp persistence) DeleteCShopImageFilterCategory(ctx *context.Context, shopImageFilterCategoryCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_shop_image_filter_category WHERE shop_image_filter_category_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, shopImageFilterCategoryCd, languageCd); err != nil {
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

func (lp persistence) GetCShopImageFilterCategoryWithKey(ctx *context.Context, shopImageFilterCategoryCd int,languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopImageFilterCategorys []*model.CShopImageFilterCategory
    script := "SELECT shop_image_filter_category_cd, language_cd, shop_image_filter_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_image_filter_category WHERE shop_image_filter_category_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, shopImageFilterCategoryCd,languageCd)
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
        cShopImageFilterCategory := &model.CShopImageFilterCategory{}
        var shopImageFilterCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopImageFilterCategory.ShopImageFilterCategoryCd, &cShopImageFilterCategory.LanguageCd, &shopImageFilterCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopImageFilterCategoryName.Valid {
            cShopImageFilterCategory.ShopImageFilterCategoryName = shopImageFilterCategoryName.String
        }
        if createDatetime.Valid {
            cShopImageFilterCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopImageFilterCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopImageFilterCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopImageFilterCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopImageFilterCategorys = append(cShopImageFilterCategorys, cShopImageFilterCategory)
    }

    return cShopImageFilterCategorys, nil
}

func (lp persistence) GetCShopImageFilterCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cShopImageFilterCategorys []*model.CShopImageFilterCategory
    script := "SELECT shop_image_filter_category_cd, language_cd, shop_image_filter_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_shop_image_filter_category WHERE language_cd = $1 ORDER BY create_datetime"
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
        cShopImageFilterCategory := &model.CShopImageFilterCategory{}
        var shopImageFilterCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cShopImageFilterCategory.ShopImageFilterCategoryCd, &cShopImageFilterCategory.LanguageCd, &shopImageFilterCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if shopImageFilterCategoryName.Valid {
            cShopImageFilterCategory.ShopImageFilterCategoryName = shopImageFilterCategoryName.String
        }
        if createDatetime.Valid {
            cShopImageFilterCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cShopImageFilterCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cShopImageFilterCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cShopImageFilterCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cShopImageFilterCategorys = append(cShopImageFilterCategorys, cShopImageFilterCategory)
    }

    return cShopImageFilterCategorys, nil
}
