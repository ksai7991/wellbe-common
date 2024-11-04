package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CBookingMethodRepository interface {
    CreateCBookingMethod(*context.Context, *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError)
    UpdateCBookingMethod(*context.Context, *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError)
    DeleteCBookingMethod(*context.Context, int, int) *errordef.LogicError
    GetCBookingMethodWithKey(*context.Context, int,int) ([]*model.CBookingMethod, *errordef.LogicError)
    GetCBookingMethodWithLanguageCd(*context.Context, int) ([]*model.CBookingMethod, *errordef.LogicError)
}
