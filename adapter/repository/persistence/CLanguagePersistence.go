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

func (lp persistence) CreateCLanguage(ctx *context.Context, cLanguage *model.CLanguage) (*model.CLanguage, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_language(language_cd, language_char_cd, language_name, sort_number, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cLanguage.LanguageCd,
                            cLanguage.LanguageCharCd,
                            cLanguage.LanguageName,
                            cLanguage.SortNumber,
                            cLanguage.CreateDatetime,
                            cLanguage.CreateFunction,
                            cLanguage.UpdateDatetime,
                            cLanguage.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cLanguage, nil
}

func (lp persistence) UpdateCLanguage(ctx *context.Context, cLanguage *model.CLanguage) (*model.CLanguage, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_language "
    script = script + "SET language_char_cd = $2, language_name = $3, sort_number = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE language_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cLanguage.LanguageCd,
                            cLanguage.LanguageCharCd,
                            cLanguage.LanguageName,
                            cLanguage.SortNumber,
                            cLanguage.CreateDatetime,
                            cLanguage.CreateFunction,
                            cLanguage.UpdateDatetime,
                            cLanguage.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cLanguage, nil
}

func (lp persistence) DeleteCLanguage(ctx *context.Context, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_language WHERE language_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, languageCd); err != nil {
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

func (lp persistence) GetCLanguageWithKey(ctx *context.Context, languageCd int) ([]*model.CLanguage, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cLanguages []*model.CLanguage
    script := "SELECT language_cd, language_char_cd, language_name, sort_number, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_language WHERE language_cd = $1 ORDER BY create_datetime"
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
        cLanguage := &model.CLanguage{}
        var languageCharCd sql.NullString
        var languageName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cLanguage.LanguageCd, &languageCharCd, &languageName, &cLanguage.SortNumber, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if languageCharCd.Valid {
            cLanguage.LanguageCharCd = languageCharCd.String
        }
        if languageName.Valid {
            cLanguage.LanguageName = languageName.String
        }
        if createDatetime.Valid {
            cLanguage.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cLanguage.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cLanguage.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cLanguage.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cLanguages = append(cLanguages, cLanguage)
    }

    return cLanguages, nil
}

func (lp persistence) GetCLanguageWithLanguageCharCd(ctx *context.Context, languageCharCd string) ([]*model.CLanguage, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cLanguages []*model.CLanguage
    script := "SELECT language_cd, language_char_cd, language_name, sort_number, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_language WHERE language_char_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCharCd)
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
        cLanguage := &model.CLanguage{}
        var languageCharCd sql.NullString
        var languageName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cLanguage.LanguageCd, &languageCharCd, &languageName, &cLanguage.SortNumber, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if languageCharCd.Valid {
            cLanguage.LanguageCharCd = languageCharCd.String
        }
        if languageName.Valid {
            cLanguage.LanguageName = languageName.String
        }
        if createDatetime.Valid {
            cLanguage.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cLanguage.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cLanguage.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cLanguage.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cLanguages = append(cLanguages, cLanguage)
    }

    return cLanguages, nil
}

func (lp persistence) GetCLanguageWithFilterCol(ctx *context.Context, languageCharCd string,languageName string) ([]*model.CLanguage, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cLanguages []*model.CLanguage
    script := "SELECT language_cd, language_char_cd, language_name, sort_number, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_language WHERE ($1 = '' or language_char_cd = $1) and language_name ILIKE concat('%', CAST($2 AS TEXT), '%') ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCharCd,languageName)
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
        cLanguage := &model.CLanguage{}
        var languageCharCd sql.NullString
        var languageName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cLanguage.LanguageCd, &languageCharCd, &languageName, &cLanguage.SortNumber, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if languageCharCd.Valid {
            cLanguage.LanguageCharCd = languageCharCd.String
        }
        if languageName.Valid {
            cLanguage.LanguageName = languageName.String
        }
        if createDatetime.Valid {
            cLanguage.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cLanguage.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cLanguage.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cLanguage.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cLanguages = append(cLanguages, cLanguage)
    }

    return cLanguages, nil
}
