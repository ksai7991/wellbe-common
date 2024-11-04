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

func (lp persistence) CreateCBookingStatus(ctx *context.Context, cBookingStatus *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_booking_status(booking_status_cd, language_cd, booking_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBookingStatus.BookingStatusCd,
                            cBookingStatus.LanguageCd,
                            cBookingStatus.BookingStatusName,
                            cBookingStatus.CreateDatetime,
                            cBookingStatus.CreateFunction,
                            cBookingStatus.UpdateDatetime,
                            cBookingStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBookingStatus, nil
}

func (lp persistence) UpdateCBookingStatus(ctx *context.Context, cBookingStatus *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_booking_status "
    script = script + "SET booking_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE booking_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBookingStatus.BookingStatusCd,
                            cBookingStatus.LanguageCd,
                            cBookingStatus.BookingStatusName,
                            cBookingStatus.CreateDatetime,
                            cBookingStatus.CreateFunction,
                            cBookingStatus.UpdateDatetime,
                            cBookingStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBookingStatus, nil
}

func (lp persistence) DeleteCBookingStatus(ctx *context.Context, bookingStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_booking_status WHERE booking_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, bookingStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCBookingStatusWithKey(ctx *context.Context, bookingStatusCd int,languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBookingStatuss []*model.CBookingStatus
    script := "SELECT booking_status_cd, language_cd, booking_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_booking_status WHERE booking_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, bookingStatusCd,languageCd)
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
        cBookingStatus := &model.CBookingStatus{}
        var bookingStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBookingStatus.BookingStatusCd, &cBookingStatus.LanguageCd, &bookingStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if bookingStatusName.Valid {
            cBookingStatus.BookingStatusName = bookingStatusName.String
        }
        if createDatetime.Valid {
            cBookingStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBookingStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBookingStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBookingStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBookingStatuss = append(cBookingStatuss, cBookingStatus)
    }

    return cBookingStatuss, nil
}

func (lp persistence) GetCBookingStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBookingStatuss []*model.CBookingStatus
    script := "SELECT booking_status_cd, language_cd, booking_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_booking_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cBookingStatus := &model.CBookingStatus{}
        var bookingStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBookingStatus.BookingStatusCd, &cBookingStatus.LanguageCd, &bookingStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if bookingStatusName.Valid {
            cBookingStatus.BookingStatusName = bookingStatusName.String
        }
        if createDatetime.Valid {
            cBookingStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBookingStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBookingStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBookingStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBookingStatuss = append(cBookingStatuss, cBookingStatus)
    }

    return cBookingStatuss, nil
}
