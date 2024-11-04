package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CAreaRepository interface {
    CreateCArea(*context.Context, *model.CArea) (*model.CArea, *errordef.LogicError)
    UpdateCArea(*context.Context, *model.CArea) (*model.CArea, *errordef.LogicError)
    DeleteCArea(*context.Context, int, int) *errordef.LogicError
    GetCAreaWithKey(*context.Context, int,int) ([]*model.CArea, *errordef.LogicError)
    GetCAreaWithLanguageCd(*context.Context, int) ([]*model.CArea, *errordef.LogicError)
    GetCAreaWithStateCd(*context.Context, int) ([]*model.CArea, *errordef.LogicError)
}
