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

func (lp persistence) CreateCReviewCategory(ctx *context.Context, cReviewCategory *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_review_category(review_category_cd, language_cd, review_category_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cReviewCategory.ReviewCategoryCd,
                            cReviewCategory.LanguageCd,
                            cReviewCategory.ReviewCategoryName,
                            cReviewCategory.CreateDatetime,
                            cReviewCategory.CreateFunction,
                            cReviewCategory.UpdateDatetime,
                            cReviewCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cReviewCategory, nil
}

func (lp persistence) UpdateCReviewCategory(ctx *context.Context, cReviewCategory *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_review_category "
    script = script + "SET review_category_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE review_category_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cReviewCategory.ReviewCategoryCd,
                            cReviewCategory.LanguageCd,
                            cReviewCategory.ReviewCategoryName,
                            cReviewCategory.CreateDatetime,
                            cReviewCategory.CreateFunction,
                            cReviewCategory.UpdateDatetime,
                            cReviewCategory.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cReviewCategory, nil
}

func (lp persistence) DeleteCReviewCategory(ctx *context.Context, reviewCategoryCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_review_category WHERE review_category_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, reviewCategoryCd, languageCd); err != nil {
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

func (lp persistence) GetCReviewCategoryWithKey(ctx *context.Context, reviewCategoryCd int,languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cReviewCategorys []*model.CReviewCategory
    script := "SELECT review_category_cd, language_cd, review_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_review_category WHERE review_category_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, reviewCategoryCd,languageCd)
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
        cReviewCategory := &model.CReviewCategory{}
        var reviewCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cReviewCategory.ReviewCategoryCd, &cReviewCategory.LanguageCd, &reviewCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if reviewCategoryName.Valid {
            cReviewCategory.ReviewCategoryName = reviewCategoryName.String
        }
        if createDatetime.Valid {
            cReviewCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cReviewCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cReviewCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cReviewCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cReviewCategorys = append(cReviewCategorys, cReviewCategory)
    }

    return cReviewCategorys, nil
}

func (lp persistence) GetCReviewCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cReviewCategorys []*model.CReviewCategory
    script := "SELECT review_category_cd, language_cd, review_category_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_review_category WHERE language_cd = $1 ORDER BY create_datetime"
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
        cReviewCategory := &model.CReviewCategory{}
        var reviewCategoryName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cReviewCategory.ReviewCategoryCd, &cReviewCategory.LanguageCd, &reviewCategoryName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if reviewCategoryName.Valid {
            cReviewCategory.ReviewCategoryName = reviewCategoryName.String
        }
        if createDatetime.Valid {
            cReviewCategory.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cReviewCategory.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cReviewCategory.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cReviewCategory.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cReviewCategorys = append(cReviewCategorys, cReviewCategory)
    }

    return cReviewCategorys, nil
}
