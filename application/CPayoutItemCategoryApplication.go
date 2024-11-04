package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CPayoutItemCategoryApplication interface {
    GetCPayoutItemCategoryWithKey(*context.Context, int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
    GetCPayoutItemCategoryWithLanguageCd(*context.Context, int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
}

type cPayoutItemCategoryApplication struct {
    cPayoutItemCategoryService service.CPayoutItemCategoryService
    transaction repository.Transaction
}

func NewCPayoutItemCategoryApplication(ls service.CPayoutItemCategoryService, tr repository.Transaction) CPayoutItemCategoryApplication {
    return &cPayoutItemCategoryApplication{
        cPayoutItemCategoryService :ls,
        transaction :tr,
    }
}

func (sa cPayoutItemCategoryApplication) GetCPayoutItemCategoryWithKey(ctx *context.Context, payoutItemCategoryCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cPayoutItemCategoryService.GetCPayoutItemCategoryWithKey(ctx, payoutItemCategoryCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cPayoutItemCategoryApplication) GetCPayoutItemCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cPayoutItemCategoryService.GetCPayoutItemCategoryWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
