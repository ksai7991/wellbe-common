package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CCheckoutTimingApplication interface {
    GetCCheckoutTimingWithKey(*context.Context, int,int) ([]*model.CCheckoutTiming, *errordef.LogicError)
    GetCCheckoutTimingWithLanguageCd(*context.Context, int) ([]*model.CCheckoutTiming, *errordef.LogicError)
}

type cCheckoutTimingApplication struct {
    cCheckoutTimingService service.CCheckoutTimingService
    transaction repository.Transaction
}

func NewCCheckoutTimingApplication(ls service.CCheckoutTimingService, tr repository.Transaction) CCheckoutTimingApplication {
    return &cCheckoutTimingApplication{
        cCheckoutTimingService :ls,
        transaction :tr,
    }
}

func (sa cCheckoutTimingApplication) GetCCheckoutTimingWithKey(ctx *context.Context, checkoutTimingCd int,languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCheckoutTimingService.GetCCheckoutTimingWithKey(ctx, checkoutTimingCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCheckoutTimingApplication) GetCCheckoutTimingWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCheckoutTimingService.GetCCheckoutTimingWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
