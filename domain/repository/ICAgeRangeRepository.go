package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CAgeRangeRepository interface {
    CreateCAgeRange(*context.Context, *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError)
    UpdateCAgeRange(*context.Context, *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError)
    DeleteCAgeRange(*context.Context, int, int) *errordef.LogicError
    GetCAgeRangeWithKey(*context.Context, int,int) ([]*model.CAgeRange, *errordef.LogicError)
    GetCAgeRangeWithLanguageCd(*context.Context, int) ([]*model.CAgeRange, *errordef.LogicError)
}
