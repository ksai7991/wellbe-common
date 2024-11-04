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

func (lp persistence) CreateCGender(ctx *context.Context, cGender *model.CGender) (*model.CGender, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_gender(gender_cd, language_cd, gender_name, gender_abbreviation, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cGender.GenderCd,
                            cGender.LanguageCd,
                            cGender.GenderName,
                            cGender.GenderAbbreviation,
                            cGender.CreateDatetime,
                            cGender.CreateFunction,
                            cGender.UpdateDatetime,
                            cGender.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cGender, nil
}

func (lp persistence) UpdateCGender(ctx *context.Context, cGender *model.CGender) (*model.CGender, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_gender "
    script = script + "SET gender_name = $3, gender_abbreviation = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE gender_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cGender.GenderCd,
                            cGender.LanguageCd,
                            cGender.GenderName,
                            cGender.GenderAbbreviation,
                            cGender.CreateDatetime,
                            cGender.CreateFunction,
                            cGender.UpdateDatetime,
                            cGender.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cGender, nil
}

func (lp persistence) DeleteCGender(ctx *context.Context, genderCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_gender WHERE gender_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, genderCd, languageCd); err != nil {
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

func (lp persistence) GetCGenderWithKey(ctx *context.Context, genderCd int,languageCd int) ([]*model.CGender, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cGenders []*model.CGender
    script := "SELECT gender_cd, language_cd, gender_name, gender_abbreviation, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_gender WHERE gender_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, genderCd,languageCd)
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
        cGender := &model.CGender{}
        var genderName sql.NullString
        var genderAbbreviation sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cGender.GenderCd, &cGender.LanguageCd, &genderName, &genderAbbreviation, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if genderName.Valid {
            cGender.GenderName = genderName.String
        }
        if genderAbbreviation.Valid {
            cGender.GenderAbbreviation = genderAbbreviation.String
        }
        if createDatetime.Valid {
            cGender.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cGender.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cGender.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cGender.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cGenders = append(cGenders, cGender)
    }

    return cGenders, nil
}

func (lp persistence) GetCGenderWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CGender, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cGenders []*model.CGender
    script := "SELECT gender_cd, language_cd, gender_name, gender_abbreviation, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_gender WHERE language_cd = $1 ORDER BY create_datetime"
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
        cGender := &model.CGender{}
        var genderName sql.NullString
        var genderAbbreviation sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cGender.GenderCd, &cGender.LanguageCd, &genderName, &genderAbbreviation, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if genderName.Valid {
            cGender.GenderName = genderName.String
        }
        if genderAbbreviation.Valid {
            cGender.GenderAbbreviation = genderAbbreviation.String
        }
        if createDatetime.Valid {
            cGender.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cGender.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cGender.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cGender.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cGenders = append(cGenders, cGender)
    }

    return cGenders, nil
}
