package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CBillingContentRepository interface {
    CreateCBillingContent(*context.Context, *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError)
    UpdateCBillingContent(*context.Context, *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError)
    DeleteCBillingContent(*context.Context, int, int) *errordef.LogicError
    GetCBillingContentWithKey(*context.Context, int,int) ([]*model.CBillingContent, *errordef.LogicError)
    GetCBillingContentWithLanguageCd(*context.Context, int) ([]*model.CBillingContent, *errordef.LogicError)
}
