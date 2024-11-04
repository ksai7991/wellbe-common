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

func (lp persistence) CreateCBookingChanel(ctx *context.Context, cBookingChanel *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_booking_chanel(booking_chanel_cd, language_cd, booking_chanel_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBookingChanel.BookingChanelCd,
                            cBookingChanel.LanguageCd,
                            cBookingChanel.BookingChanelName,
                            cBookingChanel.CreateDatetime,
                            cBookingChanel.CreateFunction,
                            cBookingChanel.UpdateDatetime,
                            cBookingChanel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBookingChanel, nil
}

func (lp persistence) UpdateCBookingChanel(ctx *context.Context, cBookingChanel *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_booking_chanel "
    script = script + "SET booking_chanel_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE booking_chanel_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cBookingChanel.BookingChanelCd,
                            cBookingChanel.LanguageCd,
                            cBookingChanel.BookingChanelName,
                            cBookingChanel.CreateDatetime,
                            cBookingChanel.CreateFunction,
                            cBookingChanel.UpdateDatetime,
                            cBookingChanel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cBookingChanel, nil
}

func (lp persistence) DeleteCBookingChanel(ctx *context.Context, bookingChanelCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_booking_chanel WHERE booking_chanel_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, bookingChanelCd, languageCd); err != nil {
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

func (lp persistence) GetCBookingChanelWithKey(ctx *context.Context, bookingChanelCd int,languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBookingChanels []*model.CBookingChanel
    script := "SELECT booking_chanel_cd, language_cd, booking_chanel_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_booking_chanel WHERE booking_chanel_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, bookingChanelCd,languageCd)
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
        cBookingChanel := &model.CBookingChanel{}
        var bookingChanelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBookingChanel.BookingChanelCd, &cBookingChanel.LanguageCd, &bookingChanelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if bookingChanelName.Valid {
            cBookingChanel.BookingChanelName = bookingChanelName.String
        }
        if createDatetime.Valid {
            cBookingChanel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBookingChanel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBookingChanel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBookingChanel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBookingChanels = append(cBookingChanels, cBookingChanel)
    }

    return cBookingChanels, nil
}

func (lp persistence) GetCBookingChanelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cBookingChanels []*model.CBookingChanel
    script := "SELECT booking_chanel_cd, language_cd, booking_chanel_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_booking_chanel WHERE language_cd = $1 ORDER BY create_datetime"
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
        cBookingChanel := &model.CBookingChanel{}
        var bookingChanelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cBookingChanel.BookingChanelCd, &cBookingChanel.LanguageCd, &bookingChanelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if bookingChanelName.Valid {
            cBookingChanel.BookingChanelName = bookingChanelName.String
        }
        if createDatetime.Valid {
            cBookingChanel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cBookingChanel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cBookingChanel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cBookingChanel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cBookingChanels = append(cBookingChanels, cBookingChanel)
    }

    return cBookingChanels, nil
}
