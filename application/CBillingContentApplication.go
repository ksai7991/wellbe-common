package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CBillingContentApplication interface {
    GetCBillingContentWithKey(*context.Context, int,int) ([]*model.CBillingContent, *errordef.LogicError)
    GetCBillingContentWithLanguageCd(*context.Context, int) ([]*model.CBillingContent, *errordef.LogicError)
}

type cBillingContentApplication struct {
    cBillingContentService service.CBillingContentService
    transaction repository.Transaction
}

func NewCBillingContentApplication(ls service.CBillingContentService, tr repository.Transaction) CBillingContentApplication {
    return &cBillingContentApplication{
        cBillingContentService :ls,
        transaction :tr,
    }
}

func (sa cBillingContentApplication) GetCBillingContentWithKey(ctx *context.Context, billingContentCd int,languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBillingContentService.GetCBillingContentWithKey(ctx, billingContentCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cBillingContentApplication) GetCBillingContentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBillingContentService.GetCBillingContentWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
