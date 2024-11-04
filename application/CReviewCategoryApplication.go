package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CReviewCategoryApplication interface {
    GetCReviewCategoryWithKey(*context.Context, int,int) ([]*model.CReviewCategory, *errordef.LogicError)
    GetCReviewCategoryWithLanguageCd(*context.Context, int) ([]*model.CReviewCategory, *errordef.LogicError)
}

type cReviewCategoryApplication struct {
    cReviewCategoryService service.CReviewCategoryService
    transaction repository.Transaction
}

func NewCReviewCategoryApplication(ls service.CReviewCategoryService, tr repository.Transaction) CReviewCategoryApplication {
    return &cReviewCategoryApplication{
        cReviewCategoryService :ls,
        transaction :tr,
    }
}

func (sa cReviewCategoryApplication) GetCReviewCategoryWithKey(ctx *context.Context, reviewCategoryCd int,languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cReviewCategoryService.GetCReviewCategoryWithKey(ctx, reviewCategoryCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cReviewCategoryApplication) GetCReviewCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cReviewCategoryService.GetCReviewCategoryWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
