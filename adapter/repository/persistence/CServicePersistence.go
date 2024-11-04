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

func (lp persistence) CreateCService(ctx *context.Context, cService *model.CService) (*model.CService, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_service(service_cd, language_cd, service_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cService.ServiceCd,
                            cService.LanguageCd,
                            cService.ServiceName,
                            cService.CreateDatetime,
                            cService.CreateFunction,
                            cService.UpdateDatetime,
                            cService.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cService, nil
}

func (lp persistence) UpdateCService(ctx *context.Context, cService *model.CService) (*model.CService, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_service "
    script = script + "SET service_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE service_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cService.ServiceCd,
                            cService.LanguageCd,
                            cService.ServiceName,
                            cService.CreateDatetime,
                            cService.CreateFunction,
                            cService.UpdateDatetime,
                            cService.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cService, nil
}

func (lp persistence) DeleteCService(ctx *context.Context, serviceCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_service WHERE service_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, serviceCd, languageCd); err != nil {
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

func (lp persistence) GetCServiceWithKey(ctx *context.Context, serviceCd int,languageCd int) ([]*model.CService, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cServices []*model.CService
    script := "SELECT service_cd, language_cd, service_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_service WHERE service_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, serviceCd,languageCd)
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
        cService := &model.CService{}
        var serviceName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cService.ServiceCd, &cService.LanguageCd, &serviceName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if serviceName.Valid {
            cService.ServiceName = serviceName.String
        }
        if createDatetime.Valid {
            cService.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cService.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cService.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cService.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cServices = append(cServices, cService)
    }

    return cServices, nil
}

func (lp persistence) GetCServiceWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CService, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cServices []*model.CService
    script := "SELECT service_cd, language_cd, service_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_service WHERE language_cd = $1 ORDER BY create_datetime"
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
        cService := &model.CService{}
        var serviceName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cService.ServiceCd, &cService.LanguageCd, &serviceName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if serviceName.Valid {
            cService.ServiceName = serviceName.String
        }
        if createDatetime.Valid {
            cService.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cService.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cService.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cService.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cServices = append(cServices, cService)
    }

    return cServices, nil
}
