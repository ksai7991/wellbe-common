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

func (lp persistence) CreateCAgeRange(ctx *context.Context, cAgeRange *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_age_range(age_range_cd, language_cd, age_range_gender, age_range_from, age_range_to, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cAgeRange.AgeRangeCd,
                            cAgeRange.LanguageCd,
                            cAgeRange.AgeRangeGender,
                            cAgeRange.AgeRangeFrom,
                            cAgeRange.AgeRangeTo,
                            cAgeRange.CreateDatetime,
                            cAgeRange.CreateFunction,
                            cAgeRange.UpdateDatetime,
                            cAgeRange.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cAgeRange, nil
}

func (lp persistence) UpdateCAgeRange(ctx *context.Context, cAgeRange *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_age_range "
    script = script + "SET age_range_gender = $3, age_range_from = $4, age_range_to = $5, create_datetime = $6, create_function = $7, update_datetime = $8, update_function = $9 "
    script = script + "WHERE age_range_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cAgeRange.AgeRangeCd,
                            cAgeRange.LanguageCd,
                            cAgeRange.AgeRangeGender,
                            cAgeRange.AgeRangeFrom,
                            cAgeRange.AgeRangeTo,
                            cAgeRange.CreateDatetime,
                            cAgeRange.CreateFunction,
                            cAgeRange.UpdateDatetime,
                            cAgeRange.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cAgeRange, nil
}

func (lp persistence) DeleteCAgeRange(ctx *context.Context, ageRangeCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_age_range WHERE age_range_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, ageRangeCd, languageCd); err != nil {
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

func (lp persistence) GetCAgeRangeWithKey(ctx *context.Context, ageRangeCd int,languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAgeRanges []*model.CAgeRange
    script := "SELECT age_range_cd, language_cd, age_range_gender, age_range_from, age_range_to, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_age_range WHERE age_range_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, ageRangeCd,languageCd)
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
        cAgeRange := &model.CAgeRange{}
        var ageRangeGender sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cAgeRange.AgeRangeCd, &cAgeRange.LanguageCd, &ageRangeGender, &cAgeRange.AgeRangeFrom, &cAgeRange.AgeRangeTo, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if ageRangeGender.Valid {
            cAgeRange.AgeRangeGender = ageRangeGender.String
        }
        if createDatetime.Valid {
            cAgeRange.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cAgeRange.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cAgeRange.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cAgeRange.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAgeRanges = append(cAgeRanges, cAgeRange)
    }

    return cAgeRanges, nil
}

func (lp persistence) GetCAgeRangeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAgeRanges []*model.CAgeRange
    script := "SELECT age_range_cd, language_cd, age_range_gender, age_range_from, age_range_to, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_age_range WHERE language_cd = $1 ORDER BY create_datetime"
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
        cAgeRange := &model.CAgeRange{}
        var ageRangeGender sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cAgeRange.AgeRangeCd, &cAgeRange.LanguageCd, &ageRangeGender, &cAgeRange.AgeRangeFrom, &cAgeRange.AgeRangeTo, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if ageRangeGender.Valid {
            cAgeRange.AgeRangeGender = ageRangeGender.String
        }
        if createDatetime.Valid {
            cAgeRange.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cAgeRange.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cAgeRange.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cAgeRange.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAgeRanges = append(cAgeRanges, cAgeRange)
    }

    return cAgeRanges, nil
}
