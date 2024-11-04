package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CRecommendLabelApplication interface {
    GetCRecommendLabelWithKey(*context.Context, int,int) ([]*model.CRecommendLabel, *errordef.LogicError)
    GetCRecommendLabelWithLanguageCd(*context.Context, int) ([]*model.CRecommendLabel, *errordef.LogicError)
}

type cRecommendLabelApplication struct {
    cRecommendLabelService service.CRecommendLabelService
    transaction repository.Transaction
}

func NewCRecommendLabelApplication(ls service.CRecommendLabelService, tr repository.Transaction) CRecommendLabelApplication {
    return &cRecommendLabelApplication{
        cRecommendLabelService :ls,
        transaction :tr,
    }
}

func (sa cRecommendLabelApplication) GetCRecommendLabelWithKey(ctx *context.Context, recommendLabelCd int,languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cRecommendLabelService.GetCRecommendLabelWithKey(ctx, recommendLabelCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cRecommendLabelApplication) GetCRecommendLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cRecommendLabelService.GetCRecommendLabelWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
