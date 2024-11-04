package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CShopStatusRepository interface {
    CreateCShopStatus(*context.Context, *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError)
    UpdateCShopStatus(*context.Context, *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError)
    DeleteCShopStatus(*context.Context, int, int) *errordef.LogicError
    GetCShopStatusWithKey(*context.Context, int,int) ([]*model.CShopStatus, *errordef.LogicError)
    GetCShopStatusWithLanguageCd(*context.Context, int) ([]*model.CShopStatus, *errordef.LogicError)
}
