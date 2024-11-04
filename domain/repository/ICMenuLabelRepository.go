package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CMenuLabelRepository interface {
    CreateCMenuLabel(*context.Context, *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError)
    UpdateCMenuLabel(*context.Context, *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError)
    DeleteCMenuLabel(*context.Context, int) *errordef.LogicError
    GetCMenuLabelWithKey(*context.Context, int) ([]*model.CMenuLabel, *errordef.LogicError)
    GetCMenuLabelWithLanguageCd(*context.Context, int) ([]*model.CMenuLabel, *errordef.LogicError)
}
