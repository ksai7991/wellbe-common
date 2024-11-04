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

func (lp persistence) CreateCPayoutItemCategory(ctx *context.Context, cPayoutItemCategory *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_payout_item_category(payout_item_category_cd, language_cd, payout_item_category_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cPayoutItemCategory.PayoutItemCategoryCd,
                            cPayoutItemCategory.LanguageCd,
                            cPayoutItemCategory.PayoutItemCategoryName,
                            cPayoutItemCategory.CreateDatetime,
                            cPayoutItemCategory.CreateFunction,
                            cPayoutItemCategory.UpdateDatetime,
                            cPayoutItemCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cPayoutItemCategory, nil
}

func (lp persistence) UpdateCPayoutItemCategory(ctx *context.Context, cPayoutItemCategory *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_payout_item_category "
    script = script + "SET language_cd = $2, payout_item_category_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE payout_item_category_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cPayoutItemCategory.PayoutItemCategoryCd,
                            cPayoutItemCategory.LanguageCd,
                            cPayoutItemCategory.PayoutItemCategoryName,
                            cPayoutItemCategory.CreateDatetime,
                            cPayoutItemCategory.CreateFunction,
                            cPayoutItemCategory.UpdateDatetime,
                            cPayoutItemCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cPayoutItemCategory, nil
}

func (lp persistence) DeleteCPayoutItemCategory(ctx *context.Context, payoutItemCategoryCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_payout_item_category WHERE payout_item_category_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, payoutItemCategoryCd); err != nil {
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

func (lp persistence) GetCPayoutItemCategoryWithKey(ctx *context.Context, payoutItemCategoryCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cPayoutItemCategorys []*model.CPayoutItemCategory
    script := "SELECT payout_item_category_cd, language_cd, payout_item_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_payout_item_category WHERE payout_item_category_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, payoutItemCategoryCd)
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
        cPayoutItemCategory := &model.CPayoutItemCategory{}
        var payoutItemCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cPayoutItemCategory.PayoutItemCategoryCd, &cPayoutItemCategory.LanguageCd, &payoutItemCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if payoutItemCategoryName.Valid {
            cPayoutItemCategory.PayoutItemCategoryName = payoutItemCategoryName.String
        }
        if createDatetime.Valid {
            cPayoutItemCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cPayoutItemCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cPayoutItemCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cPayoutItemCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cPayoutItemCategorys = append(cPayoutItemCategorys, cPayoutItemCategory)
    }

    return cPayoutItemCategorys, nil
}

func (lp persistence) GetCPayoutItemCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cPayoutItemCategorys []*model.CPayoutItemCategory
    script := "SELECT payout_item_category_cd, language_cd, payout_item_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_payout_item_category WHERE language_cd = $1 ORDER BY create_datetime"
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
        cPayoutItemCategory := &model.CPayoutItemCategory{}
        var payoutItemCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cPayoutItemCategory.PayoutItemCategoryCd, &cPayoutItemCategory.LanguageCd, &payoutItemCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if payoutItemCategoryName.Valid {
            cPayoutItemCategory.PayoutItemCategoryName = payoutItemCategoryName.String
        }
        if createDatetime.Valid {
            cPayoutItemCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cPayoutItemCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cPayoutItemCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cPayoutItemCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cPayoutItemCategorys = append(cPayoutItemCategorys, cPayoutItemCategory)
    }

    return cPayoutItemCategorys, nil
}
