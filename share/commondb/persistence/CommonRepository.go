package persistence

import (
	commondb "wellbe-common/share/commondb"
	commonmodel "wellbe-common/share/commonmodel"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type commonPersistence struct{
	db *sql.DB
} 

func NewCommonPersistence(db *sql.DB) *commonPersistence {
	return &commonPersistence{db: db}
}


func (lp commonPersistence) CreateNumbering(ctx *context.Context, numbering *commonmodel.Numbering) (*commonmodel.Numbering, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "INSERT INTO wellbe_common.numbering_definition("
	script = script + " numbering_key, initial_value, current_value, max_value, fix_length, create_datetime, create_func, update_datetime, update_func)"
	script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ;"
	_, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		if _, err := tx.Exec(script, numbering.NumberingKey, numbering.InitialValue, numbering.CurrentValue, numbering.MaxValue, numbering.FixLength, numbering.CreateDatetime, numbering.CreateFunc, numbering.UpdateDatetime, numbering.UpdateFunc); err != nil {
			return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
		}
		return nil, nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return numbering, nil
}

func (lp commonPersistence) CreateNumberingMaster(ctx *context.Context, numbering *commonmodel.Numbering) (*commonmodel.Numbering, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "INSERT INTO wellbe_common.numbering_master_definition("
	script = script + " numbering_key, initial_value, current_value, max_value, fix_length, create_datetime, create_func, update_datetime, update_func)"
	script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ;"
	_, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		if _, err := tx.Exec(script, numbering.NumberingKey, numbering.InitialValue, numbering.CurrentValue, numbering.MaxValue, numbering.FixLength, numbering.CreateDatetime, numbering.CreateFunc, numbering.UpdateDatetime, numbering.UpdateFunc); err != nil {
			return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
		}
		return nil, nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return numbering, nil
}

func (lp commonPersistence) DeleteNumbering(ctx *context.Context, numberingKey string) *errordef.LogicError {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "DELETE FROM wellbe_common.numbering_definition WHERE numbering_key = $1"
	_, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		if _, err := tx.Exec(script, numberingKey); err != nil {
			return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
		}
		return nil, nil
	})
	if err != nil {
		return nil
	}
	return nil
}
func (lp commonPersistence) DeleteNumberingMaster(ctx *context.Context, numberingKey string) *errordef.LogicError {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "DELETE FROM wellbe_common.numbering_master_definition WHERE numbering_key = $1"
	_, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		if _, err := tx.Exec(script, numberingKey); err != nil {
			return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
		}
		return nil, nil
	})
	if err != nil {
		return nil
	}
	return nil
}

func (lp commonPersistence) UpdateNumbering(ctx *context.Context, numbering *commonmodel.Numbering) (*commonmodel.Numbering, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "UPDATE wellbe_common.numbering_definition SET initial_value = $1, current_value=$2, max_value=$3, fix_length=$4, create_datetime=$5, create_func=$6, update_datetime=$7, update_func=$8 WHERE numbering_key = $9"
	_, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		if _, err := tx.Exec(script, numbering.InitialValue, numbering.CurrentValue, numbering.MaxValue, numbering.FixLength, numbering.CreateDatetime, numbering.CreateFunc, numbering.UpdateDatetime, numbering.UpdateFunc, numbering.NumberingKey); err != nil {
			return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
		}
		return nil, nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return numbering, nil
}

func (lp commonPersistence) GetOneNumbering(ctx *context.Context, numberingKey string) (*commonmodel.Numbering, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "SELECT numbering_key, initial_value, current_value, max_value, fix_length, create_datetime, create_func, update_datetime, update_func FROM wellbe_common.numbering_definition WHERE numbering_key = $1 ORDER BY numbering_key FOR UPDATE"
	row, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		row := tx.QueryRow(script, numberingKey)
		return row, nil
	})
	rowv, _ := row.(*sql.Row)
	if err != nil {
		logger.Error(err.Error())
		return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}
	numbering := &commonmodel.Numbering{}
	var updateDatetime sql.NullString
	var updateFunc sql.NullString
	errScan := rowv.Scan(&numbering.NumberingKey, &numbering.InitialValue, &numbering.CurrentValue, &numbering.MaxValue, &numbering.FixLength, &numbering.CreateDatetime, &numbering.CreateFunc, &updateDatetime, &updateFunc)
	if errScan != nil && errScan == sql.ErrNoRows {
		return nil, nil
	}
	if updateDatetime.Valid {
		numbering.UpdateDatetime = updateDatetime.String
	}
	if updateFunc.Valid {
		numbering.UpdateFunc = updateFunc.String
	}
	if errScan != nil {
		logger.Error(errScan.Error())
		return nil, &errordef.LogicError{Msg: errScan.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}

	return numbering, nil
}

func (lp commonPersistence) GetOneNumberingMaster(ctx *context.Context, numberingKey string) (*commonmodel.Numbering, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "SELECT numbering_key, initial_value, current_value, max_value, fix_length, create_datetime, create_func, update_datetime, update_func FROM wellbe_common.numbering_master_definition WHERE numbering_key = $1 ORDER BY numbering_key FOR UPDATE"
	row, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		row := tx.QueryRow(script, numberingKey)
		return row, nil
	})
	rowv, _ := row.(*sql.Row)
	if err != nil {
		logger.Error(err.Error())
		return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}
	numbering := &commonmodel.Numbering{}
	var updateDatetime sql.NullString
	var updateFunc sql.NullString
	errScan := rowv.Scan(&numbering.NumberingKey, &numbering.InitialValue, &numbering.CurrentValue, &numbering.MaxValue, &numbering.FixLength, &numbering.CreateDatetime, &numbering.CreateFunc, &updateDatetime, &updateFunc)
	if errScan != nil && errScan == sql.ErrNoRows {
		return nil, nil
	}
	if updateDatetime.Valid {
		numbering.UpdateDatetime = updateDatetime.String
	}
	if updateFunc.Valid {
		numbering.UpdateFunc = updateFunc.String
	}
	if errScan != nil {
		logger.Error(errScan.Error())
		return nil, &errordef.LogicError{Msg: errScan.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}

	return numbering, nil
}

func (lp commonPersistence) GetMailTemplate(ctx *context.Context, mailTemplateCd string, languageCd string) (*commonmodel.CMailTemplate, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	script := "SELECT mail_template_cd, language_cd, subject, body FROM wellbe_common.c_mail_template WHERE mail_template_cd = CAST($1 AS int) AND language_cd = CAST($2 AS int)"
	row, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
		row := tx.QueryRow(script, mailTemplateCd, languageCd)
		return row, nil
	})
	rowv, _ := row.(*sql.Row)
	if err != nil {
		logger.Error(err.Error())
		return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}
	cMailTemplate := &commonmodel.CMailTemplate{}
	errScan := rowv.Scan(&cMailTemplate.MailTemplateCd, &cMailTemplate.LanguageCd, &cMailTemplate.Subject, &cMailTemplate.Body)
	if errScan != nil {
		logger.Error(errScan.Error())
		return nil, &errordef.LogicError{Msg: errScan.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}

	return cMailTemplate, nil
}

func (lp commonPersistence) GetCCurrencyWithKey(ctx *context.Context, currencyCd int,languageCd int) ([]*commonmodel.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCurrencys []*commonmodel.CCurrency
    script := "SELECT currency_cd, language_cd, currency_name, currency_cd_iso, significant_digit, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_currency WHERE currency_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, currencyCd,languageCd)
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
        cCurrency := &commonmodel.CCurrency{}
        var currencyName sql.NullString
        var currencyCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCurrency.CurrencyCd, &cCurrency.LanguageCd, &currencyName, &currencyCdIso, &cCurrency.SignificantDigit, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if currencyName.Valid {
            cCurrency.CurrencyName = currencyName.String
        }
        if currencyCdIso.Valid {
            cCurrency.CurrencyCdIso = currencyCdIso.String
        }
        if createDatetime.Valid {
            cCurrency.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCurrency.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCurrency.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCurrency.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCurrencys = append(cCurrencys, cCurrency)
    }

    return cCurrencys, nil
}

func (lp commonPersistence) GetCurrencyExchangeRateWithKey(ctx *context.Context, baseCurrencyCd int,targetCurrencyCd int) ([]*commonmodel.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var currencyExchangeRates []*commonmodel.CurrencyExchangeRate
    script := "SELECT base_currency_cd, target_currency_cd, paire_name, rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.currency_exchange_rate WHERE base_currency_cd = $1 and target_currency_cd = $2 ORDER BY create_datetime"
    rows, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, baseCurrencyCd,targetCurrencyCd)
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
        currencyExchangeRate := &commonmodel.CurrencyExchangeRate{}
        var paireName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&currencyExchangeRate.BaseCurrencyCd, &currencyExchangeRate.TargetCurrencyCd, &paireName, &currencyExchangeRate.Rate, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if paireName.Valid {
            currencyExchangeRate.PaireName = paireName.String
        }
        if createDatetime.Valid {
            currencyExchangeRate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            currencyExchangeRate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            currencyExchangeRate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            currencyExchangeRate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        currencyExchangeRates = append(currencyExchangeRates, currencyExchangeRate)
    }

    return currencyExchangeRates, nil
}

func (lp commonPersistence) GetCurrencyExchangeRateAll(ctx *context.Context) ([]*commonmodel.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var currencyExchangeRates []*commonmodel.CurrencyExchangeRate
    script := "SELECT base_currency_cd, target_currency_cd, paire_name, rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.currency_exchange_rate ORDER BY create_datetime"
    rows, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script)
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
        currencyExchangeRate := &commonmodel.CurrencyExchangeRate{}
        var paireName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&currencyExchangeRate.BaseCurrencyCd, &currencyExchangeRate.TargetCurrencyCd, &paireName, &currencyExchangeRate.Rate, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if paireName.Valid {
            currencyExchangeRate.PaireName = paireName.String
        }
        if createDatetime.Valid {
            currencyExchangeRate.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            currencyExchangeRate.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            currencyExchangeRate.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            currencyExchangeRate.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        currencyExchangeRates = append(currencyExchangeRates, currencyExchangeRate)
    }

    return currencyExchangeRates, nil
}

func (lp commonPersistence) GetDefaultFeeMasterWithKey(ctx *context.Context, id string) ([]*commonmodel.DefaultFeeMaster, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var defaultFeeMasters []*commonmodel.DefaultFeeMaster
    script := "SELECT id, fee_rate, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.default_fee_master WHERE id = $1 ORDER BY create_datetime"
    rows, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
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
        defaultFeeMaster := &commonmodel.DefaultFeeMaster{}
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

func (lp commonPersistence) GetCTypeOfContactWithKey(ctx *context.Context, typeOfContactCd string,languageCd string) ([]*commonmodel.CTypeOfContact, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTypeOfContacts []*commonmodel.CTypeOfContact
    script := "SELECT type_of_contact_cd, language_cd, type_of_contact_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_type_of_contact WHERE type_of_contact_cd = CAST($1 AS INTEGER) and language_cd = CAST($2 AS INTEGER) ORDER BY create_datetime"
    rows, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, typeOfContactCd,languageCd)
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
        cTypeOfContact := &commonmodel.CTypeOfContact{}
        var typeOfContactName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTypeOfContact.TypeOfContactCd, &cTypeOfContact.LanguageCd, &typeOfContactName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if typeOfContactName.Valid {
            cTypeOfContact.TypeOfContactName = typeOfContactName.String
        }
        if createDatetime.Valid {
            cTypeOfContact.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTypeOfContact.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTypeOfContact.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTypeOfContact.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTypeOfContacts = append(cTypeOfContacts, cTypeOfContact)
    }

    return cTypeOfContacts, nil
}

func (lp commonPersistence) GetCLanguageAll(ctx *context.Context) ([]*commonmodel.CLanguage, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cLanguages []*commonmodel.CLanguage
    script := "SELECT language_cd, language_char_cd, language_name, sort_number, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_language ORDER BY language_cd"
    rows, err := commondb.DoInTx(ctx, lp.db, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script)
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
        cLanguage := &commonmodel.CLanguage{}
        var languageCharCd sql.NullString
        var languageName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cLanguage.LanguageCd, &languageCharCd, &languageName, &cLanguage.SortNumber, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if languageCharCd.Valid {
            cLanguage.LanguageCharCd = languageCharCd.String
        }
        if languageName.Valid {
            cLanguage.LanguageName = languageName.String
        }
        if createDatetime.Valid {
            cLanguage.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cLanguage.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cLanguage.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cLanguage.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cLanguages = append(cLanguages, cLanguage)
    }

    return cLanguages, nil
}