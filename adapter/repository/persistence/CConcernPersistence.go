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

func (lp persistence) CreateCConcern(ctx *context.Context, cConcern *model.CConcern) (*model.CConcern, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_concern(concern_cd, language_cd, concern_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cConcern.ConcernCd,
                            cConcern.LanguageCd,
                            cConcern.ConcernName,
                            cConcern.CreateDatetime,
                            cConcern.CreateFunction,
                            cConcern.UpdateDatetime,
                            cConcern.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cConcern, nil
}

func (lp persistence) UpdateCConcern(ctx *context.Context, cConcern *model.CConcern) (*model.CConcern, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_concern "
    script = script + "SET concern_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE concern_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cConcern.ConcernCd,
                            cConcern.LanguageCd,
                            cConcern.ConcernName,
                            cConcern.CreateDatetime,
                            cConcern.CreateFunction,
                            cConcern.UpdateDatetime,
                            cConcern.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cConcern, nil
}

func (lp persistence) DeleteCConcern(ctx *context.Context, concernCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_concern WHERE concern_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, concernCd, languageCd); err != nil {
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

func (lp persistence) GetCConcernWithKey(ctx *context.Context, concernCd int,languageCd int) ([]*model.CConcern, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cConcerns []*model.CConcern
    script := "SELECT concern_cd, language_cd, concern_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_concern WHERE concern_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, concernCd,languageCd)
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
        cConcern := &model.CConcern{}
        var concernName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cConcern.ConcernCd, &cConcern.LanguageCd, &concernName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if concernName.Valid {
            cConcern.ConcernName = concernName.String
        }
        if createDatetime.Valid {
            cConcern.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cConcern.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cConcern.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cConcern.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cConcerns = append(cConcerns, cConcern)
    }

    return cConcerns, nil
}

func (lp persistence) GetCConcernWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CConcern, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cConcerns []*model.CConcern
    script := "SELECT concern_cd, language_cd, concern_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_concern WHERE language_cd = $1 ORDER BY create_datetime"
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
        cConcern := &model.CConcern{}
        var concernName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cConcern.ConcernCd, &cConcern.LanguageCd, &concernName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if concernName.Valid {
            cConcern.ConcernName = concernName.String
        }
        if createDatetime.Valid {
            cConcern.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cConcern.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cConcern.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cConcern.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cConcerns = append(cConcerns, cConcern)
    }

    return cConcerns, nil
}
