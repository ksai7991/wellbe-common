package amount

import (
	"context"
	"fmt"
	"strings"
	"wellbe-common/share/commondb"
	"wellbe-common/share/commondb/persistence"
	"wellbe-common/share/commonmodel"
	"wellbe-common/share/commonsettings/constants"
	"wellbe-common/share/commonsettings/constants/code/cLanguage"
	errordef "wellbe-common/share/errordef"
	"wellbe-common/share/log"
	"wellbe-common/share/messages"
	"wellbe-common/share/util"
)

func FormatAmount(number float64) string {
	// 数値を文字列に変換
	numberStr := fmt.Sprintf("%.2f", number)

	// 整数部と小数部に分割
	parts := strings.Split(numberStr, ".")

	// 整数部を3桁区切りにフォーマット
	integerPart := parts[0]
	formattedInteger := formatInteger(integerPart)

	// フォーマットされた整数部と小数部を結合
	formattedNumber := formattedInteger + "." + parts[1]

	return formattedNumber
}

func formatInteger(integer string) string {
	length := len(integer)
	separatorCount := length / 3

	// 整数部を3桁区切りにフォーマット
	formattedInteger := ""
	for i := 0; i < separatorCount; i++ {
		startIndex := length - (i+1)*3
		endIndex := length - i*3
		part := integer[startIndex:endIndex]
		formattedInteger = "," + part + formattedInteger
	}

	// 余りの部分を追加
	if length%3 != 0 {
		formattedInteger = integer[0:length%3] + formattedInteger
	}

	// 先頭がカンマの場合はカンマを削除
	if formattedInteger[0:1] == "," {
		formattedInteger = formattedInteger[1:]
	}

	return formattedInteger
}

func GetCurrencyCdIso(currencyCd int, languageCd int) (string, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	currencys, err := commonPersistence.GetCCurrencyWithKey(&ctx, currencyCd, languageCd)
	if err != nil {
		return "", err
	}
	if len(currencys) == 0 {
		errmessage := fmt.Sprintf(messages.MESSAGE_EN_MASTER_DATA_IS_UNSETUP, "c_currency", "Can not get currencyCd")
		logger.Error(errmessage)
		return "", &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_MASTER_DATA_IS_UNSETUP, Msg: errmessage}
	}

	return currencys[0].CurrencyCdIso, nil
}

func GetCurrency(currencyCd int, languageCd int) (*commonmodel.CCurrency, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	currencys, err := commonPersistence.GetCCurrencyWithKey(&ctx, currencyCd, languageCd)
	if err != nil {
		return nil, err
	}
	if len(currencys) == 0 {
		errmessage := fmt.Sprintf(messages.MESSAGE_EN_MASTER_DATA_IS_UNSETUP, "c_currency", "Can not get currencyCd")
		logger.Error(errmessage)
		return nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_MASTER_DATA_IS_UNSETUP, Msg: errmessage}
	}

	return currencys[0], nil
}

func ExchangeConversion(baseCurrencyCd, targetCurrencyCd int, amount float64) (float64, *commonmodel.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	CurrencyExchangeRates, err := commonPersistence.GetCurrencyExchangeRateWithKey(&ctx, baseCurrencyCd, targetCurrencyCd)
	if err != nil {
		return float64(0), nil, err
	}
	if len(CurrencyExchangeRates) > 0 {
		return (amount * CurrencyExchangeRates[0].Rate), CurrencyExchangeRates[0], nil
	}

	CurrencyExchangeRates, err = commonPersistence.GetCurrencyExchangeRateWithKey(&ctx, targetCurrencyCd, baseCurrencyCd)
	if err != nil {
		return float64(0), nil, err
	}
	if len(CurrencyExchangeRates) > 0 {
		return (amount / CurrencyExchangeRates[0].Rate), CurrencyExchangeRates[0], nil
	}

	errmessage := fmt.Sprintf(messages.MESSAGE_EN_EXCHANGE_PAIRE_UNSETUP, baseCurrencyCd, targetCurrencyCd)
	logger.Error(errmessage)
	return float64(0), nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_EXCHANGE_PAIRE_IS_UNSETUP, Msg: errmessage}
}

func ExchangeConversionExchangeRateInput(baseCurrencyCd, targetCurrencyCd int, amount float64, exchangeRates []*commonmodel.CurrencyExchangeRate) (float64, *commonmodel.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	
	for _, v := range exchangeRates {
		if v.BaseCurrencyCd == baseCurrencyCd && v.TargetCurrencyCd == targetCurrencyCd {
			return (amount * v.Rate), v, nil
		}
	}
	
	for _, v := range exchangeRates {
		if v.BaseCurrencyCd == targetCurrencyCd && v.TargetCurrencyCd == baseCurrencyCd {
			return (amount / v.Rate), v, nil
		}
	}

	errmessage := fmt.Sprintf(messages.MESSAGE_EN_EXCHANGE_PAIRE_UNSETUP, baseCurrencyCd, targetCurrencyCd)
	logger.Error(errmessage)
	return float64(0), nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_EXCHANGE_PAIRE_IS_UNSETUP, Msg: errmessage}
}

func GetExchangeRate() ([]*commonmodel.CurrencyExchangeRate, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	CurrencyExchangeRates, err := commonPersistence.GetCurrencyExchangeRateAll(&ctx)
	if err != nil {
		return nil, err
	}

	return CurrencyExchangeRates, nil
}

func RoundWithCurrency(amount float64, currencyCd int) (float64, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	currencys, err := commonPersistence.GetCCurrencyWithKey(&ctx, currencyCd, cLanguage.ENGLISH)
	if err != nil {
		return float64(0), err
	}
	if len(currencys) == 0 {
		errmessage := fmt.Sprintf(messages.MESSAGE_EN_MASTER_DATA_IS_UNSETUP, "c_currency", "Can not get currencyCd")
		logger.Error(errmessage)
		return float64(0), &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_MASTER_DATA_IS_UNSETUP, Msg: errmessage}
	}

	outAmount, _ := RoundWithCurrencyCurrencyInput(amount, currencys[0])

	return outAmount, nil
}

func RoundWithCurrencyCurrencyInput(amount float64, currency *commonmodel.CCurrency) (float64, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()

	outAmount := util.Round(amount, float64(currency.SignificantDigit))

	return outAmount, nil
}