package service

import (
    model "wellbe-common/domain/model"
    repository "wellbe-common/domain/repository"
    errordef "wellbe-common/share/errordef"
    number "wellbe-common/share/number"
    commonconstants "wellbe-common/share/commonsettings/constants"
    datetime "wellbe-common/share/datetime"
    messages "wellbe-common/share/messages"

    "context"
)

type CurrencyExchangeRateService interface {
    CreateCurrencyExchangeRate(*context.Context, *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    UpdateCurrencyExchangeRate(*context.Context, *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithKey(*context.Context, int,int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithPaireName(*context.Context, string) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithBase(*context.Context, int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    DeleteCurrencyExchangeRate(*context.Context, int, int) *errordef.LogicError
}

type currencyExchangeRateService struct {
    currencyExchangeRateRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCurrencyExchangeRateService(pr repository.Repository, nu number.NumberUtil) CurrencyExchangeRateService {
    return &currencyExchangeRateService{
        currencyExchangeRateRepository :pr,
        numberUtil :nu,
    }
}

func (ss currencyExchangeRateService) CreateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
    currencyExchangeRate.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    currencyExchangeRate.CreateFunction = "CreateCurrencyExchangeRate"
    currencyExchangeRate.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    currencyExchangeRate.UpdateFunction = "CreateCurrencyExchangeRate"
    currencyExchangeRateRepository := ss.currencyExchangeRateRepository
    created, err := currencyExchangeRateRepository.CreateCurrencyExchangeRate(ctx, currencyExchangeRate)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss currencyExchangeRateService) UpdateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
    currencyExchangeRate.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    currencyExchangeRate.UpdateFunction = "UpdateCurrencyExchangeRate"
    currencyExchangeRateRepository := ss.currencyExchangeRateRepository
    results, err := currencyExchangeRateRepository.GetCurrencyExchangeRateWithKey(ctx, currencyExchangeRate.BaseCurrencyCd, currencyExchangeRate.TargetCurrencyCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := currencyExchangeRateRepository.UpdateCurrencyExchangeRate(ctx, currencyExchangeRate)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss currencyExchangeRateService) DeleteCurrencyExchangeRate(ctx *context.Context, baseCurrencyCd int, targetCurrencyCd int) *errordef.LogicError {
    currencyExchangeRateRepository := ss.currencyExchangeRateRepository
    results, err := currencyExchangeRateRepository.GetCurrencyExchangeRateWithKey(ctx, baseCurrencyCd, targetCurrencyCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = currencyExchangeRateRepository.DeleteCurrencyExchangeRate(ctx, baseCurrencyCd, targetCurrencyCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss currencyExchangeRateService) GetCurrencyExchangeRateWithKey(ctx *context.Context, baseCurrencyCd int,targetCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    currencyExchangeRateRepository := ss.currencyExchangeRateRepository
    currencyExchangeRate, err := currencyExchangeRateRepository.GetCurrencyExchangeRateWithKey(ctx, baseCurrencyCd,targetCurrencyCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return currencyExchangeRate, nil
}

func (ss currencyExchangeRateService) GetCurrencyExchangeRateWithPaireName(ctx *context.Context, paireName string) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    currencyExchangeRateRepository := ss.currencyExchangeRateRepository
    currencyExchangeRate, err := currencyExchangeRateRepository.GetCurrencyExchangeRateWithPaireName(ctx, paireName)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return currencyExchangeRate, nil
}

func (ss currencyExchangeRateService) GetCurrencyExchangeRateWithBase(ctx *context.Context, baseCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    currencyExchangeRateRepository := ss.currencyExchangeRateRepository
    currencyExchangeRate, err := currencyExchangeRateRepository.GetCurrencyExchangeRateWithBase(ctx, baseCurrencyCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return currencyExchangeRate, nil
}

