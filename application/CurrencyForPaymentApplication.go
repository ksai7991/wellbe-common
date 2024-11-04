package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CurrencyForPaymentApplication interface {
    GetCurrencyForPaymentWithKey(*context.Context, int) ([]*model.CurrencyForPayment, *errordef.LogicError)
}

type currencyForPaymentApplication struct {
    currencyForPaymentService service.CurrencyForPaymentService
    transaction repository.Transaction
}

func NewCurrencyForPaymentApplication(ls service.CurrencyForPaymentService, tr repository.Transaction) CurrencyForPaymentApplication {
    return &currencyForPaymentApplication{
        currencyForPaymentService :ls,
        transaction :tr,
    }
}

func (sa currencyForPaymentApplication) GetCurrencyForPaymentWithKey(ctx *context.Context, currencyCd int) ([]*model.CurrencyForPayment, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.currencyForPaymentService.GetCurrencyForPaymentWithKey(ctx, currencyCd)
    sa.transaction.Commit(ctx)
    return result, err
}
