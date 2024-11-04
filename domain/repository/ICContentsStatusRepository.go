package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CContentsStatusRepository interface {
    CreateCContentsStatus(*context.Context, *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError)
    UpdateCContentsStatus(*context.Context, *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError)
    DeleteCContentsStatus(*context.Context, int, int) *errordef.LogicError
    GetCContentsStatusWithKey(*context.Context, int,int) ([]*model.CContentsStatus, *errordef.LogicError)
    GetCContentsStatusWithLanguageCd(*context.Context, int) ([]*model.CContentsStatus, *errordef.LogicError)
}
