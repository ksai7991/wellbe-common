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

func (lp persistence) CreateCBookingMethod(ctx *context.Context, cBookingMethod *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_booking_method(booking_method_cd, language_cd, booking_method_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBookingMethod.BookingMethodCd,
                            cBookingMethod.LanguageCd,
                            cBookingMethod.BookingMethodName,
                            cBookingMethod.CreateDatetime,
                            cBookingMethod.CreateFunction,
                            cBookingMethod.UpdateDatetime,
                            cBookingMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBookingMethod, nil
}

func (lp persistence) UpdateCBookingMethod(ctx *context.Context, cBookingMethod *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_booking_method "
    script = script + "SET booking_method_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE booking_method_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBookingMethod.BookingMethodCd,
                            cBookingMethod.LanguageCd,
                            cBookingMethod.BookingMethodName,
                            cBookingMethod.CreateDatetime,
                            cBookingMethod.CreateFunction,
                            cBookingMethod.UpdateDatetime,
                            cBookingMethod.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBookingMethod, nil
}

func (lp persistence) DeleteCBookingMethod(ctx *context.Context, bookingMethodCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_booking_method WHERE booking_method_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, bookingMethodCd, languageCd); err != nil {
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

func (lp persistence) GetCBookingMethodWithKey(ctx *context.Context, bookingMethodCd int,languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBookingMethods []*model.CBookingMethod
    script := "SELECT booking_method_cd, language_cd, booking_method_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_booking_method WHERE booking_method_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, bookingMethodCd,languageCd)
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
        cBookingMethod := &model.CBookingMethod{}
        var bookingMethodName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBookingMethod.BookingMethodCd, &cBookingMethod.LanguageCd, &bookingMethodName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if bookingMethodName.Valid {
            cBookingMethod.BookingMethodName = bookingMethodName.String
        }
        if createDatetime.Valid {
            cBookingMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBookingMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBookingMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBookingMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBookingMethods = append(cBookingMethods, cBookingMethod)
    }

    return cBookingMethods, nil
}

func (lp persistence) GetCBookingMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBookingMethods []*model.CBookingMethod
    script := "SELECT booking_method_cd, language_cd, booking_method_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_booking_method WHERE language_cd = $1 ORDER BY create_datetime"
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
        cBookingMethod := &model.CBookingMethod{}
        var bookingMethodName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBookingMethod.BookingMethodCd, &cBookingMethod.LanguageCd, &bookingMethodName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if bookingMethodName.Valid {
            cBookingMethod.BookingMethodName = bookingMethodName.String
        }
        if createDatetime.Valid {
            cBookingMethod.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBookingMethod.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBookingMethod.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBookingMethod.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBookingMethods = append(cBookingMethods, cBookingMethod)
    }

    return cBookingMethods, nil
}
