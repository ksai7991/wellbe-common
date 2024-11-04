package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CCheckoutMethodRepository interface {
    CreateCCheckoutMethod(*context.Context, *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError)
    UpdateCCheckoutMethod(*context.Context, *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError)
    DeleteCCheckoutMethod(*context.Context, int, int) *errordef.LogicError
    GetCCheckoutMethodWithKey(*context.Context, int,int) ([]*model.CCheckoutMethod, *errordef.LogicError)
    GetCCheckoutMethodWithLanguageCd(*context.Context, int) ([]*model.CCheckoutMethod, *errordef.LogicError)
}
