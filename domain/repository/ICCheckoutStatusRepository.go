package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CCheckoutStatusRepository interface {
    CreateCCheckoutStatus(*context.Context, *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError)
    UpdateCCheckoutStatus(*context.Context, *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError)
    DeleteCCheckoutStatus(*context.Context, int, int) *errordef.LogicError
    GetCCheckoutStatusWithKey(*context.Context, int,int) ([]*model.CCheckoutStatus, *errordef.LogicError)
    GetCCheckoutStatusWithLanguageCd(*context.Context, int) ([]*model.CCheckoutStatus, *errordef.LogicError)
}
