package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type COrderTypeRepository interface {
    CreateCOrderType(*context.Context, *model.COrderType) (*model.COrderType, *errordef.LogicError)
    UpdateCOrderType(*context.Context, *model.COrderType) (*model.COrderType, *errordef.LogicError)
    DeleteCOrderType(*context.Context, int) *errordef.LogicError
    GetCOrderTypeWithKey(*context.Context, int) ([]*model.COrderType, *errordef.LogicError)
    GetCOrderTypeWithLanguageCd(*context.Context, int) ([]*model.COrderType, *errordef.LogicError)
}
