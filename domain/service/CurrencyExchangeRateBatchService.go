package service

import (
	"wellbe-common/domain/apiclient"
	model "wellbe-common/domain/model"
	repository "wellbe-common/domain/repository"
	"wellbe-common/share/commonsettings/constants/code/cLanguage"
	errordef "wellbe-common/share/errordef"
	number "wellbe-common/share/number"

	"context"
)

type CurrencyExchangeRateBatchService interface {
    StoreCurrencyExchangeRate(*context.Context) (*errordef.LogicError)
}

type currencyExchangeRateBatchService struct {
    currencyExchangeRateRepository repository.Repository
    numberUtil number.NumberUtil
    api apiclient.Apiclient
}

func NewCurrencyExchangeRateBatchService(pr repository.Repository, nu number.NumberUtil, a apiclient.Apiclient) CurrencyExchangeRateBatchService {
    return &currencyExchangeRateBatchService{
        currencyExchangeRateRepository :pr,
        numberUtil :nu,
        api :a,
    }
}

func (ss currencyExchangeRateBatchService) StoreCurrencyExchangeRate(ctx *context.Context) (*errordef.LogicError) {

    // ** Service
    currencyService := NewCCurrencyService(ss.currencyExchangeRateRepository, ss.numberUtil)
    currencyExchangeRateService := NewCurrencyExchangeRateService(ss.currencyExchangeRateRepository, ss.numberUtil)

    currencies, err := currencyService.GetCCurrencyWithLanguageCd(ctx, cLanguage.ENGLISH)
    if err != nil {
        return err
    }

    for _, v := range currencies {
        rates, err := ss.api.GetRate(ctx, v.CurrencyCdIso)
        if err != nil {
            return err
        }

        oldExchangeCurrencies, err := currencyExchangeRateService.GetCurrencyExchangeRateWithBase(ctx, v.CurrencyCd)
        if err != nil {
            return err
        }
        for _, d := range oldExchangeCurrencies {
            err = currencyExchangeRateService.DeleteCurrencyExchangeRate(ctx, d.BaseCurrencyCd, d.TargetCurrencyCd)
            if err != nil {
                return err
            }
        }

        for _, d := range currencies {
            if d.CurrencyCdIso == v.CurrencyCdIso {
                continue
            }

            if val, ok := rates.Rates[d.CurrencyCdIso]; ok {
                exchangeRate := &model.CurrencyExchangeRate{}
                exchangeRate.BaseCurrencyCd = v.CurrencyCd
                exchangeRate.TargetCurrencyCd = d.CurrencyCd
                exchangeRate.PaireName = v.CurrencyCdIso + d.CurrencyCdIso
                exchangeRate.Rate = val
                _, err = currencyExchangeRateService.CreateCurrencyExchangeRate(ctx, exchangeRate)
                if err != nil {
                    return err
                }
            }
        }
    }



    return nil
}