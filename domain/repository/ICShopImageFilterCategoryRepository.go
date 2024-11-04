package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CShopImageFilterCategoryRepository interface {
    CreateCShopImageFilterCategory(*context.Context, *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError)
    UpdateCShopImageFilterCategory(*context.Context, *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError)
    DeleteCShopImageFilterCategory(*context.Context, int, int) *errordef.LogicError
    GetCShopImageFilterCategoryWithKey(*context.Context, int,int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
    GetCShopImageFilterCategoryWithLanguageCd(*context.Context, int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
}
