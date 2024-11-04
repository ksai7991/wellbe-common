package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CContentsLabelRepository interface {
    CreateCContentsLabel(*context.Context, *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError)
    UpdateCContentsLabel(*context.Context, *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError)
    DeleteCContentsLabel(*context.Context, int, int) *errordef.LogicError
    GetCContentsLabelWithKey(*context.Context, int,int) ([]*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithLanguageCd(*context.Context, int) ([]*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithContentsCateogry(*context.Context, int,int) ([]*model.CContentsLabel, *errordef.LogicError)
}
