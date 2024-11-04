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

func (lp persistence) CreateCPayoutStatus(ctx *context.Context, cPayoutStatus *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_payout_status(payout_status_cd, language_cd, payout_status_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cPayoutStatus.PayoutStatusCd,
                            cPayoutStatus.LanguageCd,
                            cPayoutStatus.PayoutStatusName,
                            cPayoutStatus.CreateDatetime,
                            cPayoutStatus.CreateFunction,
                            cPayoutStatus.UpdateDatetime,
                            cPayoutStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cPayoutStatus, nil
}

func (lp persistence) UpdateCPayoutStatus(ctx *context.Context, cPayoutStatus *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_payout_status "
    script = script + "SET language_cd = $2, payout_status_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE payout_status_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cPayoutStatus.PayoutStatusCd,
                            cPayoutStatus.LanguageCd,
                            cPayoutStatus.PayoutStatusName,
                            cPayoutStatus.CreateDatetime,
                            cPayoutStatus.CreateFunction,
                            cPayoutStatus.UpdateDatetime,
                            cPayoutStatus.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cPayoutStatus, nil
}

func (lp persistence) DeleteCPayoutStatus(ctx *context.Context, payoutStatusCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_payout_status WHERE payout_status_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, payoutStatusCd); err != nil {
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

func (lp persistence) GetCPayoutStatusWithKey(ctx *context.Context, payoutStatusCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cPayoutStatuss []*model.CPayoutStatus
    script := "SELECT payout_status_cd, language_cd, payout_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_payout_status WHERE payout_status_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, payoutStatusCd)
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
        cPayoutStatus := &model.CPayoutStatus{}
        var payoutStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cPayoutStatus.PayoutStatusCd, &cPayoutStatus.LanguageCd, &payoutStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if payoutStatusName.Valid {
            cPayoutStatus.PayoutStatusName = payoutStatusName.String
        }
        if createDatetime.Valid {
            cPayoutStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cPayoutStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cPayoutStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cPayoutStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cPayoutStatuss = append(cPayoutStatuss, cPayoutStatus)
    }

    return cPayoutStatuss, nil
}

func (lp persistence) GetCPayoutStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cPayoutStatuss []*model.CPayoutStatus
    script := "SELECT payout_status_cd, language_cd, payout_status_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_payout_status WHERE language_cd = $1 ORDER BY create_datetime"
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
        cPayoutStatus := &model.CPayoutStatus{}
        var payoutStatusName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cPayoutStatus.PayoutStatusCd, &cPayoutStatus.LanguageCd, &payoutStatusName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if payoutStatusName.Valid {
            cPayoutStatus.PayoutStatusName = payoutStatusName.String
        }
        if createDatetime.Valid {
            cPayoutStatus.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cPayoutStatus.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cPayoutStatus.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cPayoutStatus.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cPayoutStatuss = append(cPayoutStatuss, cPayoutStatus)
    }

    return cPayoutStatuss, nil
}
