package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CBillingStatusRepository interface {
    CreateCBillingStatus(*context.Context, *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError)
    UpdateCBillingStatus(*context.Context, *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError)
    DeleteCBillingStatus(*context.Context, int, int) *errordef.LogicError
    GetCBillingStatusWithKey(*context.Context, int,int) ([]*model.CBillingStatus, *errordef.LogicError)
    GetCBillingStatusWithLanguageCd(*context.Context, int) ([]*model.CBillingStatus, *errordef.LogicError)
}
