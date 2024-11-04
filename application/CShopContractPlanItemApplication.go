package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CShopContractPlanItemApplication interface {
    GetCShopContractPlanItemWithKey(*context.Context, int,int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
    GetCShopContractPlanItemWithLanguageCd(*context.Context, int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
}

type cShopContractPlanItemApplication struct {
    cShopContractPlanItemService service.CShopContractPlanItemService
    transaction repository.Transaction
}

func NewCShopContractPlanItemApplication(ls service.CShopContractPlanItemService, tr repository.Transaction) CShopContractPlanItemApplication {
    return &cShopContractPlanItemApplication{
        cShopContractPlanItemService :ls,
        transaction :tr,
    }
}

func (sa cShopContractPlanItemApplication) GetCShopContractPlanItemWithKey(ctx *context.Context, shopContractPlanItemCd int,languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopContractPlanItemService.GetCShopContractPlanItemWithKey(ctx, shopContractPlanItemCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cShopContractPlanItemApplication) GetCShopContractPlanItemWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopContractPlanItemService.GetCShopContractPlanItemWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
