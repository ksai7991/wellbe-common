package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CCheckoutStatusApplication interface {
    GetCCheckoutStatusWithKey(*context.Context, int,int) ([]*model.CCheckoutStatus, *errordef.LogicError)
    GetCCheckoutStatusWithLanguageCd(*context.Context, int) ([]*model.CCheckoutStatus, *errordef.LogicError)
}

type cCheckoutStatusApplication struct {
    cCheckoutStatusService service.CCheckoutStatusService
    transaction repository.Transaction
}

func NewCCheckoutStatusApplication(ls service.CCheckoutStatusService, tr repository.Transaction) CCheckoutStatusApplication {
    return &cCheckoutStatusApplication{
        cCheckoutStatusService :ls,
        transaction :tr,
    }
}

func (sa cCheckoutStatusApplication) GetCCheckoutStatusWithKey(ctx *context.Context, checkoutStatusCd int,languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCheckoutStatusService.GetCCheckoutStatusWithKey(ctx, checkoutStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCheckoutStatusApplication) GetCCheckoutStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCheckoutStatusService.GetCCheckoutStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
