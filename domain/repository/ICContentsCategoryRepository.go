package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CContentsCategoryRepository interface {
    CreateCContentsCategory(*context.Context, *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError)
    UpdateCContentsCategory(*context.Context, *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError)
    DeleteCContentsCategory(*context.Context, int, int) *errordef.LogicError
    GetCContentsCategoryWithKey(*context.Context, int,int) ([]*model.CContentsCategory, *errordef.LogicError)
    GetCContentsCategoryWithLanguageCd(*context.Context, int) ([]*model.CContentsCategory, *errordef.LogicError)
}
