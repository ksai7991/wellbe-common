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

type CurrencyForPaymentService interface {
    CreateCurrencyForPayment(*context.Context, *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError)
    UpdateCurrencyForPayment(*context.Context, *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError)
    GetCurrencyForPaymentWithKey(*context.Context, int) ([]*model.CurrencyForPayment, *errordef.LogicError)
    DeleteCurrencyForPayment(*context.Context, int) *errordef.LogicError
}

type currencyForPaymentService struct {
    currencyForPaymentRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCurrencyForPaymentService(pr repository.Repository, nu number.NumberUtil) CurrencyForPaymentService {
    return &currencyForPaymentService{
        currencyForPaymentRepository :pr,
        numberUtil :nu,
    }
}

func (ss currencyForPaymentService) CreateCurrencyForPayment(ctx *context.Context, currencyForPayment *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError) {
    currencyForPayment.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    currencyForPayment.CreateFunction = "CreateCurrencyForPayment"
    currencyForPayment.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    currencyForPayment.UpdateFunction = "CreateCurrencyForPayment"
    currencyForPaymentRepository := ss.currencyForPaymentRepository
    created, err := currencyForPaymentRepository.CreateCurrencyForPayment(ctx, currencyForPayment)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss currencyForPaymentService) UpdateCurrencyForPayment(ctx *context.Context, currencyForPayment *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError) {
    currencyForPayment.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    currencyForPayment.UpdateFunction = "UpdateCurrencyForPayment"
    currencyForPaymentRepository := ss.currencyForPaymentRepository
    results, err := currencyForPaymentRepository.GetCurrencyForPaymentWithKey(ctx, currencyForPayment.CurrencyCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := currencyForPaymentRepository.UpdateCurrencyForPayment(ctx, currencyForPayment)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss currencyForPaymentService) DeleteCurrencyForPayment(ctx *context.Context, currencyCd int) *errordef.LogicError {
    currencyForPaymentRepository := ss.currencyForPaymentRepository
    results, err := currencyForPaymentRepository.GetCurrencyForPaymentWithKey(ctx, currencyCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = currencyForPaymentRepository.DeleteCurrencyForPayment(ctx, currencyCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss currencyForPaymentService) GetCurrencyForPaymentWithKey(ctx *context.Context, currencyCd int) ([]*model.CurrencyForPayment, *errordef.LogicError) {
    currencyForPaymentRepository := ss.currencyForPaymentRepository
    currencyForPayment, err := currencyForPaymentRepository.GetCurrencyForPaymentWithKey(ctx, currencyCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return currencyForPayment, nil
}

