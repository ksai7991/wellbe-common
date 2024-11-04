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

func (lp persistence) CreateCReviewContent(ctx *context.Context, cReviewContent *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_review_content(review_content_cd, language_cd, review_category_cd, review_content_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cReviewContent.ReviewContentCd,
                            cReviewContent.LanguageCd,
                            cReviewContent.ReviewCategoryCd,
                            cReviewContent.ReviewContentName,
                            cReviewContent.CreateDatetime,
                            cReviewContent.CreateFunction,
                            cReviewContent.UpdateDatetime,
                            cReviewContent.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cReviewContent, nil
}

func (lp persistence) UpdateCReviewContent(ctx *context.Context, cReviewContent *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_review_content "
    script = script + "SET review_category_cd = $3, review_content_name = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE review_content_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cReviewContent.ReviewContentCd,
                            cReviewContent.LanguageCd,
                            cReviewContent.ReviewCategoryCd,
                            cReviewContent.ReviewContentName,
                            cReviewContent.CreateDatetime,
                            cReviewContent.CreateFunction,
                            cReviewContent.UpdateDatetime,
                            cReviewContent.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cReviewContent, nil
}

func (lp persistence) DeleteCReviewContent(ctx *context.Context, reviewContentCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_review_content WHERE review_content_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, reviewContentCd, languageCd); err != nil {
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

func (lp persistence) GetCReviewContentWithKey(ctx *context.Context, reviewContentCd int,languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cReviewContents []*model.CReviewContent
    script := "SELECT review_content_cd, language_cd, review_category_cd, review_content_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_review_content WHERE review_content_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, reviewContentCd,languageCd)
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
        cReviewContent := &model.CReviewContent{}
        var reviewContentName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cReviewContent.ReviewContentCd, &cReviewContent.LanguageCd, &cReviewContent.ReviewCategoryCd, &reviewContentName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if reviewContentName.Valid {
            cReviewContent.ReviewContentName = reviewContentName.String
        }
        if createDatetime.Valid {
            cReviewContent.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cReviewContent.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cReviewContent.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cReviewContent.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cReviewContents = append(cReviewContents, cReviewContent)
    }

    return cReviewContents, nil
}

func (lp persistence) GetCReviewContentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cReviewContents []*model.CReviewContent
    script := "SELECT review_content_cd, language_cd, review_category_cd, review_content_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_review_content WHERE language_cd = $1 ORDER BY create_datetime"
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
        cReviewContent := &model.CReviewContent{}
        var reviewContentName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cReviewContent.ReviewContentCd, &cReviewContent.LanguageCd, &cReviewContent.ReviewCategoryCd, &reviewContentName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if reviewContentName.Valid {
            cReviewContent.ReviewContentName = reviewContentName.String
        }
        if createDatetime.Valid {
            cReviewContent.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cReviewContent.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cReviewContent.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cReviewContent.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cReviewContents = append(cReviewContents, cReviewContent)
    }

    return cReviewContents, nil
}
