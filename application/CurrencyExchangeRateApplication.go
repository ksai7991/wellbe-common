package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CurrencyExchangeRateApplication interface {
    CreateCurrencyExchangeRate(*context.Context, *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    UpdateCurrencyExchangeRate(*context.Context, *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    DeleteCurrencyExchangeRate(*context.Context, int, int) *errordef.LogicError
    GetCurrencyExchangeRateWithKey(*context.Context, int,int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithPaireName(*context.Context, string) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithBase(*context.Context, int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
}

type currencyExchangeRateApplication struct {
    currencyExchangeRateService service.CurrencyExchangeRateService
    transaction repository.Transaction
}

func NewCurrencyExchangeRateApplication(ls service.CurrencyExchangeRateService, tr repository.Transaction) CurrencyExchangeRateApplication {
    return &currencyExchangeRateApplication{
        currencyExchangeRateService :ls,
        transaction :tr,
    }
}

func (sa currencyExchangeRateApplication) CreateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.currencyExchangeRateService.CreateCurrencyExchangeRate(ctx, currencyExchangeRate)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa currencyExchangeRateApplication) UpdateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.currencyExchangeRateService.UpdateCurrencyExchangeRate(ctx, currencyExchangeRate)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa currencyExchangeRateApplication) DeleteCurrencyExchangeRate(ctx *context.Context, baseCurrencyCd int, targetCurrencyCd int) *errordef.LogicError {
    defer sa.transaction.Rollback(ctx)
    err := sa.currencyExchangeRateService.DeleteCurrencyExchangeRate(ctx, baseCurrencyCd, targetCurrencyCd)
    sa.transaction.Commit(ctx)
    return err
}

func (sa currencyExchangeRateApplication) GetCurrencyExchangeRateWithKey(ctx *context.Context, baseCurrencyCd int,targetCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.currencyExchangeRateService.GetCurrencyExchangeRateWithKey(ctx, baseCurrencyCd,targetCurrencyCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa currencyExchangeRateApplication) GetCurrencyExchangeRateWithPaireName(ctx *context.Context, paireName string) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.currencyExchangeRateService.GetCurrencyExchangeRateWithPaireName(ctx, paireName)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa currencyExchangeRateApplication) GetCurrencyExchangeRateWithBase(ctx *context.Context, baseCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.currencyExchangeRateService.GetCurrencyExchangeRateWithBase(ctx, baseCurrencyCd)
    sa.transaction.Commit(ctx)
    return result, err
}
