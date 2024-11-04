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

func (lp persistence) CreateCReviewStatus(ctx *context.Context, cReviewStatus *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_review_status(review_status_cd, language_cd, review_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cReviewStatus.ReviewStatusCd,
                            cReviewStatus.LanguageCd,
                            cReviewStatus.ReviewStatusName,
                            cReviewStatus.CreateDatetime,
                            cReviewStatus.CreateFunction,
                            cReviewStatus.UpdateDatetime,
                            cReviewStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cReviewStatus, nil
}

func (lp persistence) UpdateCReviewStatus(ctx *context.Context, cReviewStatus *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_review_status "
    script = script + "SET review_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE review_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cReviewStatus.ReviewStatusCd,
                            cReviewStatus.LanguageCd,
                            cReviewStatus.ReviewStatusName,
                            cReviewStatus.CreateDatetime,
                            cReviewStatus.CreateFunction,
                            cReviewStatus.UpdateDatetime,
                            cReviewStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cReviewStatus, nil
}

func (lp persistence) DeleteCReviewStatus(ctx *context.Context, reviewStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_review_status WHERE review_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, reviewStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCReviewStatusWithKey(ctx *context.Context, reviewStatusCd int,languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cReviewStatuss []*model.CReviewStatus
    script := "SELECT review_status_cd, language_cd, review_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_review_status WHERE review_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, reviewStatusCd,languageCd)
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
        cReviewStatus := &model.CReviewStatus{}
        var reviewStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cReviewStatus.ReviewStatusCd, &cReviewStatus.LanguageCd, &reviewStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if reviewStatusName.Valid {
            cReviewStatus.ReviewStatusName = reviewStatusName.String
        }
        if createDatetime.Valid {
            cReviewStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cReviewStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cReviewStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cReviewStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cReviewStatuss = append(cReviewStatuss, cReviewStatus)
    }

    return cReviewStatuss, nil
}

func (lp persistence) GetCReviewStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cReviewStatuss []*model.CReviewStatus
    script := "SELECT review_status_cd, language_cd, review_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_review_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cReviewStatus := &model.CReviewStatus{}
        var reviewStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cReviewStatus.ReviewStatusCd, &cReviewStatus.LanguageCd, &reviewStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if reviewStatusName.Valid {
            cReviewStatus.ReviewStatusName = reviewStatusName.String
        }
        if createDatetime.Valid {
            cReviewStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cReviewStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cReviewStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cReviewStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cReviewStatuss = append(cReviewStatuss, cReviewStatus)
    }

    return cReviewStatuss, nil
}
