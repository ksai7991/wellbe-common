package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CShopImageFilterCategoryApplication interface {
    GetCShopImageFilterCategoryWithKey(*context.Context, int,int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
    GetCShopImageFilterCategoryWithLanguageCd(*context.Context, int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
}

type cShopImageFilterCategoryApplication struct {
    cShopImageFilterCategoryService service.CShopImageFilterCategoryService
    transaction repository.Transaction
}

func NewCShopImageFilterCategoryApplication(ls service.CShopImageFilterCategoryService, tr repository.Transaction) CShopImageFilterCategoryApplication {
    return &cShopImageFilterCategoryApplication{
        cShopImageFilterCategoryService :ls,
        transaction :tr,
    }
}

func (sa cShopImageFilterCategoryApplication) GetCShopImageFilterCategoryWithKey(ctx *context.Context, shopImageFilterCategoryCd int,languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopImageFilterCategoryService.GetCShopImageFilterCategoryWithKey(ctx, shopImageFilterCategoryCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cShopImageFilterCategoryApplication) GetCShopImageFilterCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopImageFilterCategoryService.GetCShopImageFilterCategoryWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
