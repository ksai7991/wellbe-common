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

func (lp persistence) CreateCContentsLabel(ctx *context.Context, cContentsLabel *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_contents_label(contents_label_cd, language_cd, contents_category_cd, contents_label_name, contents_label_url, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContentsLabel.ContentsLabelCd,
                            cContentsLabel.LanguageCd,
                            cContentsLabel.ContentsCategoryCd,
                            cContentsLabel.ContentsLabelName,
                            cContentsLabel.ContentsLabelUrl,
                            cContentsLabel.CreateDatetime,
                            cContentsLabel.CreateFunction,
                            cContentsLabel.UpdateDatetime,
                            cContentsLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContentsLabel, nil
}

func (lp persistence) UpdateCContentsLabel(ctx *context.Context, cContentsLabel *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_contents_label "
    script = script + "SET contents_category_cd = $3, contents_label_name = $4, contents_label_url = $5, create_datetime = $6, create_function = $7, update_datetime = $8, update_function = $9 "
    script = script + "WHERE contents_label_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContentsLabel.ContentsLabelCd,
                            cContentsLabel.LanguageCd,
                            cContentsLabel.ContentsCategoryCd,
                            cContentsLabel.ContentsLabelName,
                            cContentsLabel.ContentsLabelUrl,
                            cContentsLabel.CreateDatetime,
                            cContentsLabel.CreateFunction,
                            cContentsLabel.UpdateDatetime,
                            cContentsLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContentsLabel, nil
}

func (lp persistence) DeleteCContentsLabel(ctx *context.Context, contentsLabelCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_contents_label WHERE contents_label_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, contentsLabelCd, languageCd); err != nil {
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

func (lp persistence) GetCContentsLabelWithKey(ctx *context.Context, contentsLabelCd int,languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsLabels []*model.CContentsLabel
    script := "SELECT contents_label_cd, language_cd, contents_category_cd, contents_label_name, contents_label_url, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_label WHERE contents_label_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, contentsLabelCd,languageCd)
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
        cContentsLabel := &model.CContentsLabel{}
        var contentsLabelName sql.NullString
        var contentsLabelUrl sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsLabel.ContentsLabelCd, &cContentsLabel.LanguageCd, &cContentsLabel.ContentsCategoryCd, &contentsLabelName, &contentsLabelUrl, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsLabelName.Valid {
            cContentsLabel.ContentsLabelName = contentsLabelName.String
        }
        if contentsLabelUrl.Valid {
            cContentsLabel.ContentsLabelUrl = contentsLabelUrl.String
        }
        if createDatetime.Valid {
            cContentsLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsLabels = append(cContentsLabels, cContentsLabel)
    }

    return cContentsLabels, nil
}

func (lp persistence) GetCContentsLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsLabels []*model.CContentsLabel
    script := "SELECT contents_label_cd, language_cd, contents_category_cd, contents_label_name, contents_label_url, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_label WHERE language_cd = $1 ORDER BY create_datetime"
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
        cContentsLabel := &model.CContentsLabel{}
        var contentsLabelName sql.NullString
        var contentsLabelUrl sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsLabel.ContentsLabelCd, &cContentsLabel.LanguageCd, &cContentsLabel.ContentsCategoryCd, &contentsLabelName, &contentsLabelUrl, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsLabelName.Valid {
            cContentsLabel.ContentsLabelName = contentsLabelName.String
        }
        if contentsLabelUrl.Valid {
            cContentsLabel.ContentsLabelUrl = contentsLabelUrl.String
        }
        if createDatetime.Valid {
            cContentsLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsLabels = append(cContentsLabels, cContentsLabel)
    }

    return cContentsLabels, nil
}

func (lp persistence) GetCContentsLabelWithContentsCateogry(ctx *context.Context, languageCd int,contentsCategoryCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsLabels []*model.CContentsLabel
    script := "SELECT contents_label_cd, language_cd, contents_category_cd, contents_label_name, contents_label_url, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_label WHERE language_cd = $1 and contents_category_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCd,contentsCategoryCd)
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
        cContentsLabel := &model.CContentsLabel{}
        var contentsLabelName sql.NullString
        var contentsLabelUrl sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsLabel.ContentsLabelCd, &cContentsLabel.LanguageCd, &cContentsLabel.ContentsCategoryCd, &contentsLabelName, &contentsLabelUrl, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsLabelName.Valid {
            cContentsLabel.ContentsLabelName = contentsLabelName.String
        }
        if contentsLabelUrl.Valid {
            cContentsLabel.ContentsLabelUrl = contentsLabelUrl.String
        }
        if createDatetime.Valid {
            cContentsLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsLabels = append(cContentsLabels, cContentsLabel)
    }

    return cContentsLabels, nil
}
