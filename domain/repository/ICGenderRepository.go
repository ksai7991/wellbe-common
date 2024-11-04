package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CGenderRepository interface {
    CreateCGender(*context.Context, *model.CGender) (*model.CGender, *errordef.LogicError)
    UpdateCGender(*context.Context, *model.CGender) (*model.CGender, *errordef.LogicError)
    DeleteCGender(*context.Context, int, int) *errordef.LogicError
    GetCGenderWithKey(*context.Context, int,int) ([]*model.CGender, *errordef.LogicError)
    GetCGenderWithLanguageCd(*context.Context, int) ([]*model.CGender, *errordef.LogicError)
}
