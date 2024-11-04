package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CCurrencyApplication interface {
    GetCCurrencyWithKey(*context.Context, int,int) ([]*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithCurrencyCdIso(*context.Context, string,int) ([]*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithLanguageCd(*context.Context, int) ([]*model.CCurrency, *errordef.LogicError)
}

type cCurrencyApplication struct {
    cCurrencyService service.CCurrencyService
    transaction repository.Transaction
}

func NewCCurrencyApplication(ls service.CCurrencyService, tr repository.Transaction) CCurrencyApplication {
    return &cCurrencyApplication{
        cCurrencyService :ls,
        transaction :tr,
    }
}

func (sa cCurrencyApplication) GetCCurrencyWithKey(ctx *context.Context, currencyCd int,languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCurrencyService.GetCCurrencyWithKey(ctx, currencyCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCurrencyApplication) GetCCurrencyWithCurrencyCdIso(ctx *context.Context, currencyCdIso string,significantDigit int) ([]*model.CCurrency, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCurrencyService.GetCCurrencyWithCurrencyCdIso(ctx, currencyCdIso,significantDigit)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCurrencyApplication) GetCCurrencyWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCurrencyService.GetCCurrencyWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
