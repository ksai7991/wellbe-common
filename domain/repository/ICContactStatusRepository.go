package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CContactStatusRepository interface {
    CreateCContactStatus(*context.Context, *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError)
    UpdateCContactStatus(*context.Context, *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError)
    DeleteCContactStatus(*context.Context, int, int) *errordef.LogicError
    GetCContactStatusWithKey(*context.Context, int,int) ([]*model.CContactStatus, *errordef.LogicError)
    GetCContactStatusWithLanguageCd(*context.Context, int) ([]*model.CContactStatus, *errordef.LogicError)
}
