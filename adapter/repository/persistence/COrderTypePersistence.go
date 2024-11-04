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

func (lp persistence) CreateCOrderType(ctx *context.Context, cOrderType *model.COrderType) (*model.COrderType, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_order_type(order_type_cd, language_cd, order_type_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cOrderType.OrderTypeCd,
                            cOrderType.LanguageCd,
                            cOrderType.OrderTypeName,
                            cOrderType.CreateDatetime,
                            cOrderType.CreateFunction,
                            cOrderType.UpdateDatetime,
                            cOrderType.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cOrderType, nil
}

func (lp persistence) UpdateCOrderType(ctx *context.Context, cOrderType *model.COrderType) (*model.COrderType, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_order_type "
    script = script + "SET language_cd = $2, order_type_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE order_type_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cOrderType.OrderTypeCd,
                            cOrderType.LanguageCd,
                            cOrderType.OrderTypeName,
                            cOrderType.CreateDatetime,
                            cOrderType.CreateFunction,
                            cOrderType.UpdateDatetime,
                            cOrderType.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cOrderType, nil
}

func (lp persistence) DeleteCOrderType(ctx *context.Context, orderTypeCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_order_type WHERE order_type_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, orderTypeCd); err != nil {
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

func (lp persistence) GetCOrderTypeWithKey(ctx *context.Context, orderTypeCd int) ([]*model.COrderType, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cOrderTypes []*model.COrderType
    script := "SELECT order_type_cd, language_cd, order_type_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_order_type WHERE order_type_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, orderTypeCd)
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
        cOrderType := &model.COrderType{}
        var orderTypeName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cOrderType.OrderTypeCd, &cOrderType.LanguageCd, &orderTypeName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if orderTypeName.Valid {
            cOrderType.OrderTypeName = orderTypeName.String
        }
        if createDatetime.Valid {
            cOrderType.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cOrderType.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cOrderType.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cOrderType.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cOrderTypes = append(cOrderTypes, cOrderType)
    }

    return cOrderTypes, nil
}

func (lp persistence) GetCOrderTypeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.COrderType, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cOrderTypes []*model.COrderType
    script := "SELECT order_type_cd, language_cd, order_type_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_order_type WHERE language_cd = $1 ORDER BY create_datetime"
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
        cOrderType := &model.COrderType{}
        var orderTypeName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cOrderType.OrderTypeCd, &cOrderType.LanguageCd, &orderTypeName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if orderTypeName.Valid {
            cOrderType.OrderTypeName = orderTypeName.String
        }
        if createDatetime.Valid {
            cOrderType.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cOrderType.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cOrderType.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cOrderType.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cOrderTypes = append(cOrderTypes, cOrderType)
    }

    return cOrderTypes, nil
}
