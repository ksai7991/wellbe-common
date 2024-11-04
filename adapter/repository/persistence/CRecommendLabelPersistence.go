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

func (lp persistence) CreateCRecommendLabel(ctx *context.Context, cRecommendLabel *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_recommend_label(recommend_label_cd, language_cd, recommend_label_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cRecommendLabel.RecommendLabelCd,
                            cRecommendLabel.LanguageCd,
                            cRecommendLabel.RecommendLabelName,
                            cRecommendLabel.CreateDatetime,
                            cRecommendLabel.CreateFunction,
                            cRecommendLabel.UpdateDatetime,
                            cRecommendLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cRecommendLabel, nil
}

func (lp persistence) UpdateCRecommendLabel(ctx *context.Context, cRecommendLabel *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_recommend_label "
    script = script + "SET recommend_label_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE recommend_label_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cRecommendLabel.RecommendLabelCd,
                            cRecommendLabel.LanguageCd,
                            cRecommendLabel.RecommendLabelName,
                            cRecommendLabel.CreateDatetime,
                            cRecommendLabel.CreateFunction,
                            cRecommendLabel.UpdateDatetime,
                            cRecommendLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cRecommendLabel, nil
}

func (lp persistence) DeleteCRecommendLabel(ctx *context.Context, recommendLabelCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_recommend_label WHERE recommend_label_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, recommendLabelCd, languageCd); err != nil {
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

func (lp persistence) GetCRecommendLabelWithKey(ctx *context.Context, recommendLabelCd int,languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cRecommendLabels []*model.CRecommendLabel
    script := "SELECT recommend_label_cd, language_cd, recommend_label_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_recommend_label WHERE recommend_label_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, recommendLabelCd,languageCd)
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
        cRecommendLabel := &model.CRecommendLabel{}
        var recommendLabelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cRecommendLabel.RecommendLabelCd, &cRecommendLabel.LanguageCd, &recommendLabelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if recommendLabelName.Valid {
            cRecommendLabel.RecommendLabelName = recommendLabelName.String
        }
        if createDatetime.Valid {
            cRecommendLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cRecommendLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cRecommendLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cRecommendLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cRecommendLabels = append(cRecommendLabels, cRecommendLabel)
    }

    return cRecommendLabels, nil
}

func (lp persistence) GetCRecommendLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cRecommendLabels []*model.CRecommendLabel
    script := "SELECT recommend_label_cd, language_cd, recommend_label_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_recommend_label WHERE language_cd = $1 ORDER BY create_datetime"
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
        cRecommendLabel := &model.CRecommendLabel{}
        var recommendLabelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cRecommendLabel.RecommendLabelCd, &cRecommendLabel.LanguageCd, &recommendLabelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if recommendLabelName.Valid {
            cRecommendLabel.RecommendLabelName = recommendLabelName.String
        }
        if createDatetime.Valid {
            cRecommendLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cRecommendLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cRecommendLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cRecommendLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cRecommendLabels = append(cRecommendLabels, cRecommendLabel)
    }

    return cRecommendLabels, nil
}
