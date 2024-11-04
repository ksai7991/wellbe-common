package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CReviewStatusApplication interface {
    GetCReviewStatusWithKey(*context.Context, int,int) ([]*model.CReviewStatus, *errordef.LogicError)
    GetCReviewStatusWithLanguageCd(*context.Context, int) ([]*model.CReviewStatus, *errordef.LogicError)
}

type cReviewStatusApplication struct {
    cReviewStatusService service.CReviewStatusService
    transaction repository.Transaction
}

func NewCReviewStatusApplication(ls service.CReviewStatusService, tr repository.Transaction) CReviewStatusApplication {
    return &cReviewStatusApplication{
        cReviewStatusService :ls,
        transaction :tr,
    }
}

func (sa cReviewStatusApplication) GetCReviewStatusWithKey(ctx *context.Context, reviewStatusCd int,languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cReviewStatusService.GetCReviewStatusWithKey(ctx, reviewStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cReviewStatusApplication) GetCReviewStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cReviewStatusService.GetCReviewStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
