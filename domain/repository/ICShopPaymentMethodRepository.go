package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CShopPaymentMethodRepository interface {
    CreateCShopPaymentMethod(*context.Context, *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError)
    UpdateCShopPaymentMethod(*context.Context, *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError)
    DeleteCShopPaymentMethod(*context.Context, int) *errordef.LogicError
    GetCShopPaymentMethodWithKey(*context.Context, int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
    GetCShopPaymentMethodWithLanguageCd(*context.Context, int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
}
