package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CCheckoutMethodApplication interface {
    GetCCheckoutMethodWithKey(*context.Context, int,int) ([]*model.CCheckoutMethod, *errordef.LogicError)
    GetCCheckoutMethodWithLanguageCd(*context.Context, int) ([]*model.CCheckoutMethod, *errordef.LogicError)
}

type cCheckoutMethodApplication struct {
    cCheckoutMethodService service.CCheckoutMethodService
    transaction repository.Transaction
}

func NewCCheckoutMethodApplication(ls service.CCheckoutMethodService, tr repository.Transaction) CCheckoutMethodApplication {
    return &cCheckoutMethodApplication{
        cCheckoutMethodService :ls,
        transaction :tr,
    }
}

func (sa cCheckoutMethodApplication) GetCCheckoutMethodWithKey(ctx *context.Context, checkoutMethodCd int,languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCheckoutMethodService.GetCCheckoutMethodWithKey(ctx, checkoutMethodCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCheckoutMethodApplication) GetCCheckoutMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCheckoutMethodService.GetCCheckoutMethodWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
