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

func (lp persistence) CreateCBillingContent(ctx *context.Context, cBillingContent *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_billing_content(billing_content_cd, language_cd, billing_content_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBillingContent.BillingContentCd,
                            cBillingContent.LanguageCd,
                            cBillingContent.BillingContentName,
                            cBillingContent.CreateDatetime,
                            cBillingContent.CreateFunction,
                            cBillingContent.UpdateDatetime,
                            cBillingContent.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBillingContent, nil
}

func (lp persistence) UpdateCBillingContent(ctx *context.Context, cBillingContent *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_billing_content "
    script = script + "SET billing_content_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE billing_content_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBillingContent.BillingContentCd,
                            cBillingContent.LanguageCd,
                            cBillingContent.BillingContentName,
                            cBillingContent.CreateDatetime,
                            cBillingContent.CreateFunction,
                            cBillingContent.UpdateDatetime,
                            cBillingContent.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBillingContent, nil
}

func (lp persistence) DeleteCBillingContent(ctx *context.Context, billingContentCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_billing_content WHERE billing_content_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, billingContentCd, languageCd); err != nil {
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

func (lp persistence) GetCBillingContentWithKey(ctx *context.Context, billingContentCd int,languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBillingContents []*model.CBillingContent
    script := "SELECT billing_content_cd, language_cd, billing_content_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_billing_content WHERE billing_content_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, billingContentCd,languageCd)
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
        cBillingContent := &model.CBillingContent{}
        var billingContentName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBillingContent.BillingContentCd, &cBillingContent.LanguageCd, &billingContentName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if billingContentName.Valid {
            cBillingContent.BillingContentName = billingContentName.String
        }
        if createDatetime.Valid {
            cBillingContent.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBillingContent.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBillingContent.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBillingContent.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBillingContents = append(cBillingContents, cBillingContent)
    }

    return cBillingContents, nil
}

func (lp persistence) GetCBillingContentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBillingContents []*model.CBillingContent
    script := "SELECT billing_content_cd, language_cd, billing_content_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_billing_content WHERE language_cd = $1 ORDER BY create_datetime"
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
        cBillingContent := &model.CBillingContent{}
        var billingContentName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBillingContent.BillingContentCd, &cBillingContent.LanguageCd, &billingContentName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if billingContentName.Valid {
            cBillingContent.BillingContentName = billingContentName.String
        }
        if createDatetime.Valid {
            cBillingContent.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBillingContent.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBillingContent.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBillingContent.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBillingContents = append(cBillingContents, cBillingContent)
    }

    return cBillingContents, nil
}
