package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CBillingStatusApplication interface {
    GetCBillingStatusWithKey(*context.Context, int,int) ([]*model.CBillingStatus, *errordef.LogicError)
    GetCBillingStatusWithLanguageCd(*context.Context, int) ([]*model.CBillingStatus, *errordef.LogicError)
}

type cBillingStatusApplication struct {
    cBillingStatusService service.CBillingStatusService
    transaction repository.Transaction
}

func NewCBillingStatusApplication(ls service.CBillingStatusService, tr repository.Transaction) CBillingStatusApplication {
    return &cBillingStatusApplication{
        cBillingStatusService :ls,
        transaction :tr,
    }
}

func (sa cBillingStatusApplication) GetCBillingStatusWithKey(ctx *context.Context, billingStatusCd int,languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBillingStatusService.GetCBillingStatusWithKey(ctx, billingStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cBillingStatusApplication) GetCBillingStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBillingStatusService.GetCBillingStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
