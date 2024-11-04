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

func (lp persistence) CreateCContentsCategory(ctx *context.Context, cContentsCategory *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_contents_category(contents_category_cd, language_cd, contents_category_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContentsCategory.ContentsCategoryCd,
                            cContentsCategory.LanguageCd,
                            cContentsCategory.ContentsCategoryName,
                            cContentsCategory.CreateDatetime,
                            cContentsCategory.CreateFunction,
                            cContentsCategory.UpdateDatetime,
                            cContentsCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContentsCategory, nil
}

func (lp persistence) UpdateCContentsCategory(ctx *context.Context, cContentsCategory *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_contents_category "
    script = script + "SET contents_category_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE contents_category_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContentsCategory.ContentsCategoryCd,
                            cContentsCategory.LanguageCd,
                            cContentsCategory.ContentsCategoryName,
                            cContentsCategory.CreateDatetime,
                            cContentsCategory.CreateFunction,
                            cContentsCategory.UpdateDatetime,
                            cContentsCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContentsCategory, nil
}

func (lp persistence) DeleteCContentsCategory(ctx *context.Context, contentsCategoryCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_contents_category WHERE contents_category_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, contentsCategoryCd, languageCd); err != nil {
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

func (lp persistence) GetCContentsCategoryWithKey(ctx *context.Context, contentsCategoryCd int,languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsCategorys []*model.CContentsCategory
    script := "SELECT contents_category_cd, language_cd, contents_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_category WHERE contents_category_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, contentsCategoryCd,languageCd)
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
        cContentsCategory := &model.CContentsCategory{}
        var contentsCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsCategory.ContentsCategoryCd, &cContentsCategory.LanguageCd, &contentsCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsCategoryName.Valid {
            cContentsCategory.ContentsCategoryName = contentsCategoryName.String
        }
        if createDatetime.Valid {
            cContentsCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsCategorys = append(cContentsCategorys, cContentsCategory)
    }

    return cContentsCategorys, nil
}

func (lp persistence) GetCContentsCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsCategorys []*model.CContentsCategory
    script := "SELECT contents_category_cd, language_cd, contents_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_category WHERE language_cd = $1 ORDER BY create_datetime"
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
        cContentsCategory := &model.CContentsCategory{}
        var contentsCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsCategory.ContentsCategoryCd, &cContentsCategory.LanguageCd, &contentsCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsCategoryName.Valid {
            cContentsCategory.ContentsCategoryName = contentsCategoryName.String
        }
        if createDatetime.Valid {
            cContentsCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsCategorys = append(cContentsCategorys, cContentsCategory)
    }

    return cContentsCategorys, nil
}
