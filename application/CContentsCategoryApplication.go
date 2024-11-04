package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CContentsCategoryApplication interface {
    GetCContentsCategoryWithKey(*context.Context, int,int) ([]*model.CContentsCategory, *errordef.LogicError)
    GetCContentsCategoryWithLanguageCd(*context.Context, int) ([]*model.CContentsCategory, *errordef.LogicError)
}

type cContentsCategoryApplication struct {
    cContentsCategoryService service.CContentsCategoryService
    transaction repository.Transaction
}

func NewCContentsCategoryApplication(ls service.CContentsCategoryService, tr repository.Transaction) CContentsCategoryApplication {
    return &cContentsCategoryApplication{
        cContentsCategoryService :ls,
        transaction :tr,
    }
}

func (sa cContentsCategoryApplication) GetCContentsCategoryWithKey(ctx *context.Context, contentsCategoryCd int,languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsCategoryService.GetCContentsCategoryWithKey(ctx, contentsCategoryCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cContentsCategoryApplication) GetCContentsCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cContentsCategoryService.GetCContentsCategoryWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
