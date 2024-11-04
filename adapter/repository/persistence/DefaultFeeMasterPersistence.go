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

func (lp persistence) CreateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.default_fee_master(id, fee_rate, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            defaultFeeMaster.Id,
                            defaultFeeMaster.FeeRate,
                            defaultFeeMaster.CreateDatetime,
                            defaultFeeMaster.CreateFunction,
                            defaultFeeMaster.UpdateDatetime,
                            defaultFeeMaster.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return defaultFeeMaster, nil
}

func (lp persistence) UpdateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.default_fee_master "
    script = script + "SET fee_rate = $2, create_datetime = $3, create_function = $4, update_datetime = $5, update_function = $6 "
    script = script + "WHERE id = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            defaultFeeMaster.Id,
                            defaultFeeMaster.FeeRate,
                            defaultFeeMaster.CreateDatetime,
                            defaultFeeMaster.CreateFunction,
                            defaultFeeMaster.UpdateDatetime,
                            defaultFeeMaster.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return defaultFeeMaster, nil
}

func (lp persistence) DeleteDefaultFeeMaster(ctx *context.Context, id string) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.default_fee_master WHERE id = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, id); err != nil {
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

func (lp persistence) GetDefaultFeeMasterWithKey(ctx *context.Context, id string) ([]*model.DefaultFeeMaster, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var defaultFeeMasters []*model.DefaultFeeMaster
    script := "SELECT id, fee_rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.default_fee_master WHERE id = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, id)
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
        defaultFeeMaster := &model.DefaultFeeMaster{}
        var id sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&id, &defaultFeeMaster.FeeRate, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if id.Valid {
            defaultFeeMaster.Id = id.String
        }
        if createDatetime.Valid {
            defaultFeeMaster.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            defaultFeeMaster.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            defaultFeeMaster.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            defaultFeeMaster.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        defaultFeeMasters = append(defaultFeeMasters, defaultFeeMaster)
    }

    return defaultFeeMasters, nil
}
