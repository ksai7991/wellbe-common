package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CReviewContentApplication interface {
    GetCReviewContentWithKey(*context.Context, int,int) ([]*model.CReviewContent, *errordef.LogicError)
    GetCReviewContentWithLanguageCd(*context.Context, int) ([]*model.CReviewContent, *errordef.LogicError)
}

type cReviewContentApplication struct {
    cReviewContentService service.CReviewContentService
    transaction repository.Transaction
}

func NewCReviewContentApplication(ls service.CReviewContentService, tr repository.Transaction) CReviewContentApplication {
    return &cReviewContentApplication{
        cReviewContentService :ls,
        transaction :tr,
    }
}

func (sa cReviewContentApplication) GetCReviewContentWithKey(ctx *context.Context, reviewContentCd int,languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cReviewContentService.GetCReviewContentWithKey(ctx, reviewContentCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cReviewContentApplication) GetCReviewContentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cReviewContentService.GetCReviewContentWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
