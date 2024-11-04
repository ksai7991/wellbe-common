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

func (lp persistence) CreateCContentsStatus(ctx *context.Context, cContentsStatus *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_contents_status(contents_status_cd, language_cd, contents_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContentsStatus.ContentsStatusCd,
                            cContentsStatus.LanguageCd,
                            cContentsStatus.ContentsStatusName,
                            cContentsStatus.CreateDatetime,
                            cContentsStatus.CreateFunction,
                            cContentsStatus.UpdateDatetime,
                            cContentsStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContentsStatus, nil
}

func (lp persistence) UpdateCContentsStatus(ctx *context.Context, cContentsStatus *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_contents_status "
    script = script + "SET contents_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE contents_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContentsStatus.ContentsStatusCd,
                            cContentsStatus.LanguageCd,
                            cContentsStatus.ContentsStatusName,
                            cContentsStatus.CreateDatetime,
                            cContentsStatus.CreateFunction,
                            cContentsStatus.UpdateDatetime,
                            cContentsStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContentsStatus, nil
}

func (lp persistence) DeleteCContentsStatus(ctx *context.Context, contentsStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_contents_status WHERE contents_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, contentsStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCContentsStatusWithKey(ctx *context.Context, contentsStatusCd int,languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsStatuss []*model.CContentsStatus
    script := "SELECT contents_status_cd, language_cd, contents_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_status WHERE contents_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, contentsStatusCd,languageCd)
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
        cContentsStatus := &model.CContentsStatus{}
        var contentsStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsStatus.ContentsStatusCd, &cContentsStatus.LanguageCd, &contentsStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsStatusName.Valid {
            cContentsStatus.ContentsStatusName = contentsStatusName.String
        }
        if createDatetime.Valid {
            cContentsStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsStatuss = append(cContentsStatuss, cContentsStatus)
    }

    return cContentsStatuss, nil
}

func (lp persistence) GetCContentsStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContentsStatuss []*model.CContentsStatus
    script := "SELECT contents_status_cd, language_cd, contents_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contents_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cContentsStatus := &model.CContentsStatus{}
        var contentsStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContentsStatus.ContentsStatusCd, &cContentsStatus.LanguageCd, &contentsStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contentsStatusName.Valid {
            cContentsStatus.ContentsStatusName = contentsStatusName.String
        }
        if createDatetime.Valid {
            cContentsStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContentsStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContentsStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContentsStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContentsStatuss = append(cContentsStatuss, cContentsStatus)
    }

    return cContentsStatuss, nil
}
