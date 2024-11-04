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

func (lp persistence) CreateCMailTemplate(ctx *context.Context, cMailTemplate *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_mail_template(mail_template_cd, language_cd, subject, body, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cMailTemplate.MailTemplateCd,
                            cMailTemplate.LanguageCd,
                            cMailTemplate.Subject,
                            cMailTemplate.Body,
                            cMailTemplate.CreateDatetime,
                            cMailTemplate.CreateFunction,
                            cMailTemplate.UpdateDatetime,
                            cMailTemplate.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cMailTemplate, nil
}

func (lp persistence) UpdateCMailTemplate(ctx *context.Context, cMailTemplate *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_mail_template "
    script = script + "SET subject = $3, body = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE mail_template_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cMailTemplate.MailTemplateCd,
                            cMailTemplate.LanguageCd,
                            cMailTemplate.Subject,
                            cMailTemplate.Body,
                            cMailTemplate.CreateDatetime,
                            cMailTemplate.CreateFunction,
                            cMailTemplate.UpdateDatetime,
                            cMailTemplate.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cMailTemplate, nil
}

func (lp persistence) DeleteCMailTemplate(ctx *context.Context, mailTemplateCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_mail_template WHERE mail_template_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, mailTemplateCd, languageCd); err != nil {
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

func (lp persistence) GetCMailTemplateWithKey(ctx *context.Context, mailTemplateCd int,languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cMailTemplates []*model.CMailTemplate
    script := "SELECT mail_template_cd, language_cd, subject, body, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_mail_template WHERE mail_template_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, mailTemplateCd,languageCd)
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
        cMailTemplate := &model.CMailTemplate{}
        var subject sql.NullString
        var body sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cMailTemplate.MailTemplateCd, &cMailTemplate.LanguageCd, &subject, &body, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if subject.Valid {
            cMailTemplate.Subject = subject.String
        }
        if body.Valid {
            cMailTemplate.Body = body.String
        }
        if createDatetime.Valid {
            cMailTemplate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cMailTemplate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cMailTemplate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cMailTemplate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cMailTemplates = append(cMailTemplates, cMailTemplate)
    }

    return cMailTemplates, nil
}

func (lp persistence) GetCMailTemplateWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cMailTemplates []*model.CMailTemplate
    script := "SELECT mail_template_cd, language_cd, subject, body, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_mail_template WHERE language_cd = $1 ORDER BY create_datetime"
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
        cMailTemplate := &model.CMailTemplate{}
        var subject sql.NullString
        var body sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cMailTemplate.MailTemplateCd, &cMailTemplate.LanguageCd, &subject, &body, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if subject.Valid {
            cMailTemplate.Subject = subject.String
        }
        if body.Valid {
            cMailTemplate.Body = body.String
        }
        if createDatetime.Valid {
            cMailTemplate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cMailTemplate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cMailTemplate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cMailTemplate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cMailTemplates = append(cMailTemplates, cMailTemplate)
    }

    return cMailTemplates, nil
}
