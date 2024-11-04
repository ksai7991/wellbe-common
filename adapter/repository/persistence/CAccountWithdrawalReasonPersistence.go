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

func (lp persistence) CreateCAccountWithdrawalReason(ctx *context.Context, cAccountWithdrawalReason *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_account_withdrawal_reason(account_withdrawal_reason_cd, language_cd, account_withdrawal_reason_name, account_withdrawal_reason_abbreviation, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cAccountWithdrawalReason.AccountWithdrawalReasonCd,
                            cAccountWithdrawalReason.LanguageCd,
                            cAccountWithdrawalReason.AccountWithdrawalReasonName,
                            cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation,
                            cAccountWithdrawalReason.CreateDatetime,
                            cAccountWithdrawalReason.CreateFunction,
                            cAccountWithdrawalReason.UpdateDatetime,
                            cAccountWithdrawalReason.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cAccountWithdrawalReason, nil
}

func (lp persistence) UpdateCAccountWithdrawalReason(ctx *context.Context, cAccountWithdrawalReason *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_account_withdrawal_reason "
    script = script + "SET account_withdrawal_reason_name = $3, account_withdrawal_reason_abbreviation = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE account_withdrawal_reason_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cAccountWithdrawalReason.AccountWithdrawalReasonCd,
                            cAccountWithdrawalReason.LanguageCd,
                            cAccountWithdrawalReason.AccountWithdrawalReasonName,
                            cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation,
                            cAccountWithdrawalReason.CreateDatetime,
                            cAccountWithdrawalReason.CreateFunction,
                            cAccountWithdrawalReason.UpdateDatetime,
                            cAccountWithdrawalReason.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cAccountWithdrawalReason, nil
}

func (lp persistence) DeleteCAccountWithdrawalReason(ctx *context.Context, accountWithdrawalReasonCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_account_withdrawal_reason WHERE account_withdrawal_reason_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, accountWithdrawalReasonCd, languageCd); err != nil {
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

func (lp persistence) GetCAccountWithdrawalReasonWithKey(ctx *context.Context, accountWithdrawalReasonCd int,languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAccountWithdrawalReasons []*model.CAccountWithdrawalReason
    script := "SELECT account_withdrawal_reason_cd, language_cd, account_withdrawal_reason_name, account_withdrawal_reason_abbreviation, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_account_withdrawal_reason WHERE account_withdrawal_reason_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, accountWithdrawalReasonCd,languageCd)
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
        cAccountWithdrawalReason := &model.CAccountWithdrawalReason{}
        var accountWithdrawalReasonName sql.NullString
        var accountWithdrawalReasonAbbreviation sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cAccountWithdrawalReason.AccountWithdrawalReasonCd, &cAccountWithdrawalReason.LanguageCd, &accountWithdrawalReasonName, &accountWithdrawalReasonAbbreviation, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if accountWithdrawalReasonName.Valid {
            cAccountWithdrawalReason.AccountWithdrawalReasonName = accountWithdrawalReasonName.String
        }
        if accountWithdrawalReasonAbbreviation.Valid {
            cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation = accountWithdrawalReasonAbbreviation.String
        }
        if createDatetime.Valid {
            cAccountWithdrawalReason.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cAccountWithdrawalReason.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cAccountWithdrawalReason.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cAccountWithdrawalReason.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAccountWithdrawalReasons = append(cAccountWithdrawalReasons, cAccountWithdrawalReason)
    }

    return cAccountWithdrawalReasons, nil
}

func (lp persistence) GetCAccountWithdrawalReasonWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAccountWithdrawalReasons []*model.CAccountWithdrawalReason
    script := "SELECT account_withdrawal_reason_cd, language_cd, account_withdrawal_reason_name, account_withdrawal_reason_abbreviation, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_account_withdrawal_reason WHERE language_cd = $1 ORDER BY create_datetime"
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
        cAccountWithdrawalReason := &model.CAccountWithdrawalReason{}
        var accountWithdrawalReasonName sql.NullString
        var accountWithdrawalReasonAbbreviation sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cAccountWithdrawalReason.AccountWithdrawalReasonCd, &cAccountWithdrawalReason.LanguageCd, &accountWithdrawalReasonName, &accountWithdrawalReasonAbbreviation, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if accountWithdrawalReasonName.Valid {
            cAccountWithdrawalReason.AccountWithdrawalReasonName = accountWithdrawalReasonName.String
        }
        if accountWithdrawalReasonAbbreviation.Valid {
            cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation = accountWithdrawalReasonAbbreviation.String
        }
        if createDatetime.Valid {
            cAccountWithdrawalReason.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cAccountWithdrawalReason.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cAccountWithdrawalReason.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cAccountWithdrawalReason.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAccountWithdrawalReasons = append(cAccountWithdrawalReasons, cAccountWithdrawalReason)
    }

    return cAccountWithdrawalReasons, nil
}
