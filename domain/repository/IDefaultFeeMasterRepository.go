package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type DefaultFeeMasterRepository interface {
    CreateDefaultFeeMaster(*context.Context, *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    UpdateDefaultFeeMaster(*context.Context, *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    DeleteDefaultFeeMaster(*context.Context, string) *errordef.LogicError
    GetDefaultFeeMasterWithKey(*context.Context, string) ([]*model.DefaultFeeMaster, *errordef.LogicError)
}
