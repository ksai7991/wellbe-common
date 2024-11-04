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

func (lp persistence) CreateCContactStatus(ctx *context.Context, cContactStatus *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_contact_status(contact_status_cd, language_cd, contact_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContactStatus.ContactStatusCd,
                            cContactStatus.LanguageCd,
                            cContactStatus.ContactStatusName,
                            cContactStatus.CreateDatetime,
                            cContactStatus.CreateFunction,
                            cContactStatus.UpdateDatetime,
                            cContactStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContactStatus, nil
}

func (lp persistence) UpdateCContactStatus(ctx *context.Context, cContactStatus *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_contact_status "
    script = script + "SET contact_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE contact_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cContactStatus.ContactStatusCd,
                            cContactStatus.LanguageCd,
                            cContactStatus.ContactStatusName,
                            cContactStatus.CreateDatetime,
                            cContactStatus.CreateFunction,
                            cContactStatus.UpdateDatetime,
                            cContactStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cContactStatus, nil
}

func (lp persistence) DeleteCContactStatus(ctx *context.Context, contactStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_contact_status WHERE contact_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, contactStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCContactStatusWithKey(ctx *context.Context, contactStatusCd int,languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContactStatuss []*model.CContactStatus
    script := "SELECT contact_status_cd, language_cd, contact_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contact_status WHERE contact_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, contactStatusCd,languageCd)
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
        cContactStatus := &model.CContactStatus{}
        var contactStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContactStatus.ContactStatusCd, &cContactStatus.LanguageCd, &contactStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contactStatusName.Valid {
            cContactStatus.ContactStatusName = contactStatusName.String
        }
        if createDatetime.Valid {
            cContactStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContactStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContactStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContactStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContactStatuss = append(cContactStatuss, cContactStatus)
    }

    return cContactStatuss, nil
}

func (lp persistence) GetCContactStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cContactStatuss []*model.CContactStatus
    script := "SELECT contact_status_cd, language_cd, contact_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_contact_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cContactStatus := &model.CContactStatus{}
        var contactStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cContactStatus.ContactStatusCd, &cContactStatus.LanguageCd, &contactStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if contactStatusName.Valid {
            cContactStatus.ContactStatusName = contactStatusName.String
        }
        if createDatetime.Valid {
            cContactStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cContactStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cContactStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cContactStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cContactStatuss = append(cContactStatuss, cContactStatus)
    }

    return cContactStatuss, nil
}
