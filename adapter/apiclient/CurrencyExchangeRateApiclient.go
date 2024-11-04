package apiclient

import (
	"encoding/json"
	"io"
	"net/http"
	model "wellbe-common/domain/model"
	settings "wellbe-common/settings"
	constants "wellbe-common/settings/constants"
	commonconstants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
	"wellbe-common/share/messages"

	"context"

	_ "github.com/lib/pq"
)


type CurrencyExchangeRateEntity struct {
    Success bool `json:"success"`
    Timestamp int `json:"timestamp"`
    Date string `json:"date"`
    model.CurrencyExchangeRateApi
}

type CurrencyExchangeRateErrorEntity struct {
    Success bool `json:"success"`
    Error CurrencyExchangeRateErrorDetailEntity `json:"error"`
}

type CurrencyExchangeRateErrorDetailEntity struct {
    Code string `json:"code"`
    Info string `json:"message"`
}

func (a api) GetRate(ctx *context.Context, baseCurrencyCdIso string) (*model.CurrencyExchangeRateApi, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()

    url := settings.GetExchangerateDomain()+constants.ENV_EXCHANGERATE_LATEST + "?access_key=" + settings.GetExchangerateAccessKey() + "&base=" + baseCurrencyCdIso

    req, err := http.NewRequest(
        "GET",
        url,
        nil,
    )
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code: commonconstants.LOGIC_ERROR_CODE_API_INITIATE_ERROR}
    }

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        logger.Error(err.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code: commonconstants.LOGIC_ERROR_CODE_API_INITIATE_ERROR}
    }
    defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
    if err != nil {
        logger.Error(err.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code: commonconstants.LOGIC_ERROR_CODE_API_INITIATE_ERROR}
    }

    if resp.StatusCode != 200 {
        var currencyExchangeRateErrorEntity CurrencyExchangeRateErrorEntity
        if err := json.Unmarshal(body, &currencyExchangeRateErrorEntity); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: commonconstants.LOGIC_ERROR_CODE_API_INITIATE_ERROR}
        }
        logger.Error(currencyExchangeRateErrorEntity.Error.Info)
        return nil, &errordef.LogicError{Msg: currencyExchangeRateErrorEntity.Error.Info, Code: commonconstants.LOGIC_ERROR_CODE_API_EXCHANGE_RATE}
    } else {
        var currencyExchangeRateEntity CurrencyExchangeRateEntity
        if err := json.Unmarshal(body, &currencyExchangeRateEntity); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: commonconstants.LOGIC_ERROR_CODE_API_INITIATE_ERROR}
        }
        if !currencyExchangeRateEntity.Success {
            return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_SERVER_ERROR, Code: commonconstants.LOGIC_ERROR_CODE_API_EXCHANGE_RATE}
        }
        return &currencyExchangeRateEntity.CurrencyExchangeRateApi, nil
    }
}