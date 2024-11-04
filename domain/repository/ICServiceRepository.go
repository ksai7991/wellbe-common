package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CServiceRepository interface {
    CreateCService(*context.Context, *model.CService) (*model.CService, *errordef.LogicError)
    UpdateCService(*context.Context, *model.CService) (*model.CService, *errordef.LogicError)
    DeleteCService(*context.Context, int, int) *errordef.LogicError
    GetCServiceWithKey(*context.Context, int,int) ([]*model.CService, *errordef.LogicError)
    GetCServiceWithLanguageCd(*context.Context, int) ([]*model.CService, *errordef.LogicError)
}
