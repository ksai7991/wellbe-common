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

func (lp persistence) CreateCTypeOfContact(ctx *context.Context, cTypeOfContact *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_type_of_contact(type_of_contact_cd, language_cd, type_of_contact_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cTypeOfContact.TypeOfContactCd,
                            cTypeOfContact.LanguageCd,
                            cTypeOfContact.TypeOfContactName,
                            cTypeOfContact.CreateDatetime,
                            cTypeOfContact.CreateFunction,
                            cTypeOfContact.UpdateDatetime,
                            cTypeOfContact.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cTypeOfContact, nil
}

func (lp persistence) UpdateCTypeOfContact(ctx *context.Context, cTypeOfContact *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_type_of_contact "
    script = script + "SET type_of_contact_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE type_of_contact_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cTypeOfContact.TypeOfContactCd,
                            cTypeOfContact.LanguageCd,
                            cTypeOfContact.TypeOfContactName,
                            cTypeOfContact.CreateDatetime,
                            cTypeOfContact.CreateFunction,
                            cTypeOfContact.UpdateDatetime,
                            cTypeOfContact.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cTypeOfContact, nil
}

func (lp persistence) DeleteCTypeOfContact(ctx *context.Context, typeOfContactCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_type_of_contact WHERE type_of_contact_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, typeOfContactCd, languageCd); err != nil {
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

func (lp persistence) GetCTypeOfContactWithKey(ctx *context.Context, typeOfContactCd int,languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTypeOfContacts []*model.CTypeOfContact
    script := "SELECT type_of_contact_cd, language_cd, type_of_contact_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_type_of_contact WHERE type_of_contact_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, typeOfContactCd,languageCd)
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
        cTypeOfContact := &model.CTypeOfContact{}
        var typeOfContactName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTypeOfContact.TypeOfContactCd, &cTypeOfContact.LanguageCd, &typeOfContactName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if typeOfContactName.Valid {
            cTypeOfContact.TypeOfContactName = typeOfContactName.String
        }
        if createDatetime.Valid {
            cTypeOfContact.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTypeOfContact.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTypeOfContact.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTypeOfContact.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTypeOfContacts = append(cTypeOfContacts, cTypeOfContact)
    }

    return cTypeOfContacts, nil
}

func (lp persistence) GetCTypeOfContactWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTypeOfContacts []*model.CTypeOfContact
    script := "SELECT type_of_contact_cd, language_cd, type_of_contact_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_type_of_contact WHERE language_cd = $1 ORDER BY create_datetime"
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
        cTypeOfContact := &model.CTypeOfContact{}
        var typeOfContactName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTypeOfContact.TypeOfContactCd, &cTypeOfContact.LanguageCd, &typeOfContactName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if typeOfContactName.Valid {
            cTypeOfContact.TypeOfContactName = typeOfContactName.String
        }
        if createDatetime.Valid {
            cTypeOfContact.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTypeOfContact.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTypeOfContact.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTypeOfContact.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTypeOfContacts = append(cTypeOfContacts, cTypeOfContact)
    }

    return cTypeOfContacts, nil
}
