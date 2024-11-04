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

func (lp persistence) CreateCWeekday(ctx *context.Context, cWeekday *model.CWeekday) (*model.CWeekday, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_weekday(weekday_cd, language_cd, weekday_name, weekday_abbreviation, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cWeekday.WeekdayCd,
                            cWeekday.LanguageCd,
                            cWeekday.WeekdayName,
                            cWeekday.WeekdayAbbreviation,
                            cWeekday.CreateDatetime,
                            cWeekday.CreateFunction,
                            cWeekday.UpdateDatetime,
                            cWeekday.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cWeekday, nil
}

func (lp persistence) UpdateCWeekday(ctx *context.Context, cWeekday *model.CWeekday) (*model.CWeekday, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_weekday "
    script = script + "SET language_cd = $2, weekday_name = $3, weekday_abbreviation = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE weekday_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cWeekday.WeekdayCd,
                            cWeekday.LanguageCd,
                            cWeekday.WeekdayName,
                            cWeekday.WeekdayAbbreviation,
                            cWeekday.CreateDatetime,
                            cWeekday.CreateFunction,
                            cWeekday.UpdateDatetime,
                            cWeekday.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cWeekday, nil
}

func (lp persistence) DeleteCWeekday(ctx *context.Context, weekdayCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_weekday WHERE weekday_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, weekdayCd); err != nil {
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

func (lp persistence) GetCWeekdayWithKey(ctx *context.Context, weekdayCd int) ([]*model.CWeekday, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cWeekdays []*model.CWeekday
    script := "SELECT weekday_cd, language_cd, weekday_name, weekday_abbreviation, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_weekday WHERE weekday_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, weekdayCd)
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
        cWeekday := &model.CWeekday{}
        var weekdayName sql.NullString
        var weekdayAbbreviation sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cWeekday.WeekdayCd, &cWeekday.LanguageCd, &weekdayName, &weekdayAbbreviation, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if weekdayName.Valid {
            cWeekday.WeekdayName = weekdayName.String
        }
        if weekdayAbbreviation.Valid {
            cWeekday.WeekdayAbbreviation = weekdayAbbreviation.String
        }
        if createDatetime.Valid {
            cWeekday.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cWeekday.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cWeekday.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cWeekday.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cWeekdays = append(cWeekdays, cWeekday)
    }

    return cWeekdays, nil
}

func (lp persistence) GetCWeekdayWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CWeekday, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cWeekdays []*model.CWeekday
    script := "SELECT weekday_cd, language_cd, weekday_name, weekday_abbreviation, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_weekday WHERE language_cd = $1 ORDER BY create_datetime"
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
        cWeekday := &model.CWeekday{}
        var weekdayName sql.NullString
        var weekdayAbbreviation sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cWeekday.WeekdayCd, &cWeekday.LanguageCd, &weekdayName, &weekdayAbbreviation, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if weekdayName.Valid {
            cWeekday.WeekdayName = weekdayName.String
        }
        if weekdayAbbreviation.Valid {
            cWeekday.WeekdayAbbreviation = weekdayAbbreviation.String
        }
        if createDatetime.Valid {
            cWeekday.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cWeekday.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cWeekday.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cWeekday.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cWeekdays = append(cWeekdays, cWeekday)
    }

    return cWeekdays, nil
}
