package application

import (
	repository "wellbe-common/domain/repository"
	service "wellbe-common/domain/service"

	errordef "wellbe-common/share/errordef"

	"context"
)

type CurrencyExchangeRateBatchApplication interface {
    StoreCurrencyExchangeRate(*context.Context) (*errordef.LogicError)
}

type currencyExchangeRateBatchApplication struct {
    currencyExchangeRateService service.CurrencyExchangeRateBatchService
    transaction repository.Transaction
}

func NewCurrencyExchangeRateBatchApplication(ls service.CurrencyExchangeRateBatchService, tr repository.Transaction) CurrencyExchangeRateBatchApplication {
    return &currencyExchangeRateBatchApplication{
        currencyExchangeRateService :ls,
        transaction :tr,
    }
}

func (sa currencyExchangeRateBatchApplication) StoreCurrencyExchangeRate(ctx *context.Context) (*errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    err := sa.currencyExchangeRateService.StoreCurrencyExchangeRate(ctx)
    sa.transaction.Commit(ctx)
    return err
}