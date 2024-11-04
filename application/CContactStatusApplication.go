package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CContactStatusApplication interface {
    GetCContactStatusWithKey(*context.Context, int,int) ([]*model.CContactStatus, *errordef.LogicError)
    GetCContactStatusWithLanguageCd(*context.Context, int) ([]*model.CContactStatus, *errordef.LogicError)
}

type cContactStatusApplication struct {
    cContactStatusService service.CContactStatusService
    transaction repository.Transaction
}

func NewCContactStatusApplication(ls service.CContactStatusService, tr repository.Transaction) CContactStatusApplication {
    return &cContactStatusApplication{
        cContactStatusService :ls,
        transaction :tr,
    }
}

func (sa cContactStatusApplication) GetCContactStatusWithKey(ctx *context.Context, contactStatusCd int,languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContactStatusService.GetCContactStatusWithKey(ctx, contactStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cContactStatusApplication) GetCContactStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContactStatusService.GetCContactStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
