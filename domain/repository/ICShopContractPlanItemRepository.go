package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CShopContractPlanItemRepository interface {
    CreateCShopContractPlanItem(*context.Context, *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError)
    UpdateCShopContractPlanItem(*context.Context, *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError)
    DeleteCShopContractPlanItem(*context.Context, int, int) *errordef.LogicError
    GetCShopContractPlanItemWithKey(*context.Context, int,int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
    GetCShopContractPlanItemWithLanguageCd(*context.Context, int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
}
