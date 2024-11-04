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

func (lp persistence) CreateCState(ctx *context.Context, cState *model.CState) (*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_state(state_cd, language_cd, country_cd, state_name, state_cd_iso, timezone_iana, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cState.StateCd,
                            cState.LanguageCd,
                            cState.CountryCd,
                            cState.StateName,
                            cState.StateCdIso,
                            cState.TimezoneIana,
                            cState.CreateDatetime,
                            cState.CreateFunction,
                            cState.UpdateDatetime,
                            cState.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cState, nil
}

func (lp persistence) UpdateCState(ctx *context.Context, cState *model.CState) (*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_state "
    script = script + "SET country_cd = $3, state_name = $4, state_cd_iso = $5, timezone_iana = $6, create_datetime = $7, create_function = $8, update_datetime = $9, update_function = $10 "
    script = script + "WHERE state_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cState.StateCd,
                            cState.LanguageCd,
                            cState.CountryCd,
                            cState.StateName,
                            cState.StateCdIso,
                            cState.TimezoneIana,
                            cState.CreateDatetime,
                            cState.CreateFunction,
                            cState.UpdateDatetime,
                            cState.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cState, nil
}

func (lp persistence) DeleteCState(ctx *context.Context, stateCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_state WHERE state_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, stateCd, languageCd); err != nil {
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

func (lp persistence) GetCStateWithKey(ctx *context.Context, stateCd int,languageCd int) ([]*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cStates []*model.CState
    script := "SELECT state_cd, language_cd, country_cd, state_name, state_cd_iso, timezone_iana, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_state WHERE state_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, stateCd,languageCd)
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
        cState := &model.CState{}
        var stateName sql.NullString
        var stateCdIso sql.NullString
        var timezoneIana sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cState.StateCd, &cState.LanguageCd, &cState.CountryCd, &stateName, &stateCdIso, &timezoneIana, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if stateName.Valid {
            cState.StateName = stateName.String
        }
        if stateCdIso.Valid {
            cState.StateCdIso = stateCdIso.String
        }
        if timezoneIana.Valid {
            cState.TimezoneIana = timezoneIana.String
        }
        if createDatetime.Valid {
            cState.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cState.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cState.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cState.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cStates = append(cStates, cState)
    }

    return cStates, nil
}

func (lp persistence) GetCStateWithStateCdIso(ctx *context.Context, stateCdIso string) ([]*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cStates []*model.CState
    script := "SELECT state_cd, language_cd, country_cd, state_name, state_cd_iso, timezone_iana, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_state WHERE state_cd_iso = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, stateCdIso)
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
        cState := &model.CState{}
        var stateName sql.NullString
        var stateCdIso sql.NullString
        var timezoneIana sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cState.StateCd, &cState.LanguageCd, &cState.CountryCd, &stateName, &stateCdIso, &timezoneIana, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if stateName.Valid {
            cState.StateName = stateName.String
        }
        if stateCdIso.Valid {
            cState.StateCdIso = stateCdIso.String
        }
        if timezoneIana.Valid {
            cState.TimezoneIana = timezoneIana.String
        }
        if createDatetime.Valid {
            cState.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cState.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cState.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cState.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cStates = append(cStates, cState)
    }

    return cStates, nil
}

func (lp persistence) GetCStateWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cStates []*model.CState
    script := "SELECT state_cd, language_cd, country_cd, state_name, state_cd_iso, timezone_iana, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_state WHERE language_cd = $1 ORDER BY create_datetime"
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
        cState := &model.CState{}
        var stateName sql.NullString
        var stateCdIso sql.NullString
        var timezoneIana sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cState.StateCd, &cState.LanguageCd, &cState.CountryCd, &stateName, &stateCdIso, &timezoneIana, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if stateName.Valid {
            cState.StateName = stateName.String
        }
        if stateCdIso.Valid {
            cState.StateCdIso = stateCdIso.String
        }
        if timezoneIana.Valid {
            cState.TimezoneIana = timezoneIana.String
        }
        if createDatetime.Valid {
            cState.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cState.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cState.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cState.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cStates = append(cStates, cState)
    }

    return cStates, nil
}

func (lp persistence) GetCStateWithCountryCd(ctx *context.Context, languageCd int,countryCd int) ([]*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cStates []*model.CState
    script := "SELECT state_cd, language_cd, country_cd, state_name, state_cd_iso, timezone_iana, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_state WHERE language_cd = $1 and country_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCd,countryCd)
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
        cState := &model.CState{}
        var stateName sql.NullString
        var stateCdIso sql.NullString
        var timezoneIana sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cState.StateCd, &cState.LanguageCd, &cState.CountryCd, &stateName, &stateCdIso, &timezoneIana, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if stateName.Valid {
            cState.StateName = stateName.String
        }
        if stateCdIso.Valid {
            cState.StateCdIso = stateCdIso.String
        }
        if timezoneIana.Valid {
            cState.TimezoneIana = timezoneIana.String
        }
        if createDatetime.Valid {
            cState.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cState.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cState.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cState.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cStates = append(cStates, cState)
    }

    return cStates, nil
}

func (lp persistence) GetCStateWithStateCd(ctx *context.Context, stateCd int) ([]*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cStates []*model.CState
    script := "SELECT state_cd, language_cd, country_cd, state_name, state_cd_iso, timezone_iana, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_state WHERE state_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, stateCd)
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
        cState := &model.CState{}
        var stateName sql.NullString
        var stateCdIso sql.NullString
        var timezoneIana sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cState.StateCd, &cState.LanguageCd, &cState.CountryCd, &stateName, &stateCdIso, &timezoneIana, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if stateName.Valid {
            cState.StateName = stateName.String
        }
        if stateCdIso.Valid {
            cState.StateCdIso = stateCdIso.String
        }
        if timezoneIana.Valid {
            cState.TimezoneIana = timezoneIana.String
        }
        if createDatetime.Valid {
            cState.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cState.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cState.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cState.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cStates = append(cStates, cState)
    }

    return cStates, nil
}
