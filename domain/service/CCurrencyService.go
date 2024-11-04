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

type CCurrencyService interface {
    CreateCCurrency(*context.Context, *model.CCurrency) (*model.CCurrency, *errordef.LogicError)
    UpdateCCurrency(*context.Context, *model.CCurrency) (*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithKey(*context.Context, int,int) ([]*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithCurrencyCdIso(*context.Context, string,int) ([]*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithLanguageCd(*context.Context, int) ([]*model.CCurrency, *errordef.LogicError)
    DeleteCCurrency(*context.Context, int, int) *errordef.LogicError
}

type cCurrencyService struct {
    cCurrencyRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCCurrencyService(pr repository.Repository, nu number.NumberUtil) CCurrencyService {
    return &cCurrencyService{
        cCurrencyRepository :pr,
        numberUtil :nu,
    }
}

func (ss cCurrencyService) CreateCCurrency(ctx *context.Context, cCurrency *model.CCurrency) (*model.CCurrency, *errordef.LogicError) {
    cCurrency.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCurrency.CreateFunction = "CreateCCurrency"
    cCurrency.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCurrency.UpdateFunction = "CreateCCurrency"
    cCurrencyRepository := ss.cCurrencyRepository
    created, err := cCurrencyRepository.CreateCCurrency(ctx, cCurrency)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cCurrencyService) UpdateCCurrency(ctx *context.Context, cCurrency *model.CCurrency) (*model.CCurrency, *errordef.LogicError) {
    cCurrency.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCurrency.UpdateFunction = "UpdateCCurrency"
    cCurrencyRepository := ss.cCurrencyRepository
    results, err := cCurrencyRepository.GetCCurrencyWithKey(ctx, cCurrency.CurrencyCd, cCurrency.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cCurrencyRepository.UpdateCCurrency(ctx, cCurrency)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cCurrencyService) DeleteCCurrency(ctx *context.Context, currencyCd int, languageCd int) *errordef.LogicError {
    cCurrencyRepository := ss.cCurrencyRepository
    results, err := cCurrencyRepository.GetCCurrencyWithKey(ctx, currencyCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cCurrencyRepository.DeleteCCurrency(ctx, currencyCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cCurrencyService) GetCCurrencyWithKey(ctx *context.Context, currencyCd int,languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
    cCurrencyRepository := ss.cCurrencyRepository
    cCurrency, err := cCurrencyRepository.GetCCurrencyWithKey(ctx, currencyCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCurrency, nil
}

func (ss cCurrencyService) GetCCurrencyWithCurrencyCdIso(ctx *context.Context, currencyCdIso string,significantDigit int) ([]*model.CCurrency, *errordef.LogicError) {
    cCurrencyRepository := ss.cCurrencyRepository
    cCurrency, err := cCurrencyRepository.GetCCurrencyWithCurrencyCdIso(ctx, currencyCdIso,significantDigit)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCurrency, nil
}

func (ss cCurrencyService) GetCCurrencyWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
    cCurrencyRepository := ss.cCurrencyRepository
    cCurrency, err := cCurrencyRepository.GetCCurrencyWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCurrency, nil
}

