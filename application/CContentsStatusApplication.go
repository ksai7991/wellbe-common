package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CContentsStatusApplication interface {
    GetCContentsStatusWithKey(*context.Context, int,int) ([]*model.CContentsStatus, *errordef.LogicError)
    GetCContentsStatusWithLanguageCd(*context.Context, int) ([]*model.CContentsStatus, *errordef.LogicError)
}

type cContentsStatusApplication struct {
    cContentsStatusService service.CContentsStatusService
    transaction repository.Transaction
}

func NewCContentsStatusApplication(ls service.CContentsStatusService, tr repository.Transaction) CContentsStatusApplication {
    return &cContentsStatusApplication{
        cContentsStatusService :ls,
        transaction :tr,
    }
}

func (sa cContentsStatusApplication) GetCContentsStatusWithKey(ctx *context.Context, contentsStatusCd int,languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsStatusService.GetCContentsStatusWithKey(ctx, contentsStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cContentsStatusApplication) GetCContentsStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsStatusService.GetCContentsStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
