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

func (lp persistence) CreateCTreatmentTimeRange(ctx *context.Context, cTreatmentTimeRange *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_treatment_time_range(treatment_time_cd, language_cd, treatment_time_name, min_time, max_time, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cTreatmentTimeRange.TreatmentTimeCd,
                            cTreatmentTimeRange.LanguageCd,
                            cTreatmentTimeRange.TreatmentTimeName,
                            cTreatmentTimeRange.MinTime,
                            cTreatmentTimeRange.MaxTime,
                            cTreatmentTimeRange.CreateDatetime,
                            cTreatmentTimeRange.CreateFunction,
                            cTreatmentTimeRange.UpdateDatetime,
                            cTreatmentTimeRange.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cTreatmentTimeRange, nil
}

func (lp persistence) UpdateCTreatmentTimeRange(ctx *context.Context, cTreatmentTimeRange *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_treatment_time_range "
    script = script + "SET treatment_time_name = $3, min_time = $4, max_time = $5, create_datetime = $6, create_function = $7, update_datetime = $8, update_function = $9 "
    script = script + "WHERE treatment_time_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cTreatmentTimeRange.TreatmentTimeCd,
                            cTreatmentTimeRange.LanguageCd,
                            cTreatmentTimeRange.TreatmentTimeName,
                            cTreatmentTimeRange.MinTime,
                            cTreatmentTimeRange.MaxTime,
                            cTreatmentTimeRange.CreateDatetime,
                            cTreatmentTimeRange.CreateFunction,
                            cTreatmentTimeRange.UpdateDatetime,
                            cTreatmentTimeRange.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cTreatmentTimeRange, nil
}

func (lp persistence) DeleteCTreatmentTimeRange(ctx *context.Context, treatmentTimeCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_treatment_time_range WHERE treatment_time_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, treatmentTimeCd, languageCd); err != nil {
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

func (lp persistence) GetCTreatmentTimeRangeWithKey(ctx *context.Context, treatmentTimeCd int,languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTreatmentTimeRanges []*model.CTreatmentTimeRange
    script := "SELECT treatment_time_cd, language_cd, treatment_time_name, min_time, max_time, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_treatment_time_range WHERE treatment_time_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, treatmentTimeCd,languageCd)
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
        cTreatmentTimeRange := &model.CTreatmentTimeRange{}
        var treatmentTimeName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTreatmentTimeRange.TreatmentTimeCd, &cTreatmentTimeRange.LanguageCd, &treatmentTimeName, &cTreatmentTimeRange.MinTime, &cTreatmentTimeRange.MaxTime, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if treatmentTimeName.Valid {
            cTreatmentTimeRange.TreatmentTimeName = treatmentTimeName.String
        }
        if createDatetime.Valid {
            cTreatmentTimeRange.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTreatmentTimeRange.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTreatmentTimeRange.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTreatmentTimeRange.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTreatmentTimeRanges = append(cTreatmentTimeRanges, cTreatmentTimeRange)
    }

    return cTreatmentTimeRanges, nil
}

func (lp persistence) GetCTreatmentTimeRangeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTreatmentTimeRanges []*model.CTreatmentTimeRange
    script := "SELECT treatment_time_cd, language_cd, treatment_time_name, min_time, max_time, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_treatment_time_range WHERE language_cd = $1 ORDER BY create_datetime"
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
        cTreatmentTimeRange := &model.CTreatmentTimeRange{}
        var treatmentTimeName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTreatmentTimeRange.TreatmentTimeCd, &cTreatmentTimeRange.LanguageCd, &treatmentTimeName, &cTreatmentTimeRange.MinTime, &cTreatmentTimeRange.MaxTime, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if treatmentTimeName.Valid {
            cTreatmentTimeRange.TreatmentTimeName = treatmentTimeName.String
        }
        if createDatetime.Valid {
            cTreatmentTimeRange.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTreatmentTimeRange.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTreatmentTimeRange.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTreatmentTimeRange.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTreatmentTimeRanges = append(cTreatmentTimeRanges, cTreatmentTimeRange)
    }

    return cTreatmentTimeRanges, nil
}
