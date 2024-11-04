package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CContentsLabelApplication interface {
    GetCContentsLabelWithKey(*context.Context, int,int) ([]*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithLanguageCd(*context.Context, int) ([]*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithContentsCateogry(*context.Context, int,int) ([]*model.CContentsLabel, *errordef.LogicError)
}

type cContentsLabelApplication struct {
    cContentsLabelService service.CContentsLabelService
    transaction repository.Transaction
}

func NewCContentsLabelApplication(ls service.CContentsLabelService, tr repository.Transaction) CContentsLabelApplication {
    return &cContentsLabelApplication{
        cContentsLabelService :ls,
        transaction :tr,
    }
}

func (sa cContentsLabelApplication) GetCContentsLabelWithKey(ctx *context.Context, contentsLabelCd int,languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsLabelService.GetCContentsLabelWithKey(ctx, contentsLabelCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cContentsLabelApplication) GetCContentsLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsLabelService.GetCContentsLabelWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cContentsLabelApplication) GetCContentsLabelWithContentsCateogry(ctx *context.Context, languageCd int,contentsCategoryCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsLabelService.GetCContentsLabelWithContentsCateogry(ctx, languageCd,contentsCategoryCd)
    sa.transaction.Commit(ctx)
    return result, err
}
