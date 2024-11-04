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

func (lp persistence) CreateCInvoiceStatus(ctx *context.Context, cInvoiceStatus *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_invoice_status(invoice_status_cd, language_cd, invoice_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cInvoiceStatus.InvoiceStatusCd,
                            cInvoiceStatus.LanguageCd,
                            cInvoiceStatus.InvoiceStatusName,
                            cInvoiceStatus.CreateDatetime,
                            cInvoiceStatus.CreateFunction,
                            cInvoiceStatus.UpdateDatetime,
                            cInvoiceStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cInvoiceStatus, nil
}

func (lp persistence) UpdateCInvoiceStatus(ctx *context.Context, cInvoiceStatus *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_invoice_status "
    script = script + "SET invoice_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE invoice_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cInvoiceStatus.InvoiceStatusCd,
                            cInvoiceStatus.LanguageCd,
                            cInvoiceStatus.InvoiceStatusName,
                            cInvoiceStatus.CreateDatetime,
                            cInvoiceStatus.CreateFunction,
                            cInvoiceStatus.UpdateDatetime,
                            cInvoiceStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cInvoiceStatus, nil
}

func (lp persistence) DeleteCInvoiceStatus(ctx *context.Context, invoiceStatusCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_invoice_status WHERE invoice_status_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, invoiceStatusCd, languageCd); err != nil {
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

func (lp persistence) GetCInvoiceStatusWithKey(ctx *context.Context, invoiceStatusCd int,languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cInvoiceStatuss []*model.CInvoiceStatus
    script := "SELECT invoice_status_cd, language_cd, invoice_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_invoice_status WHERE invoice_status_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, invoiceStatusCd,languageCd)
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
        cInvoiceStatus := &model.CInvoiceStatus{}
        var invoiceStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cInvoiceStatus.InvoiceStatusCd, &cInvoiceStatus.LanguageCd, &invoiceStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if invoiceStatusName.Valid {
            cInvoiceStatus.InvoiceStatusName = invoiceStatusName.String
        }
        if createDatetime.Valid {
            cInvoiceStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cInvoiceStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cInvoiceStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cInvoiceStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cInvoiceStatuss = append(cInvoiceStatuss, cInvoiceStatus)
    }

    return cInvoiceStatuss, nil
}

func (lp persistence) GetCInvoiceStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cInvoiceStatuss []*model.CInvoiceStatus
    script := "SELECT invoice_status_cd, language_cd, invoice_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_invoice_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cInvoiceStatus := &model.CInvoiceStatus{}
        var invoiceStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cInvoiceStatus.InvoiceStatusCd, &cInvoiceStatus.LanguageCd, &invoiceStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if invoiceStatusName.Valid {
            cInvoiceStatus.InvoiceStatusName = invoiceStatusName.String
        }
        if createDatetime.Valid {
            cInvoiceStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cInvoiceStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cInvoiceStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cInvoiceStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cInvoiceStatuss = append(cInvoiceStatuss, cInvoiceStatus)
    }

    return cInvoiceStatuss, nil
}
