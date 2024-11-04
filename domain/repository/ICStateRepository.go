package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CStateRepository interface {
    CreateCState(*context.Context, *model.CState) (*model.CState, *errordef.LogicError)
    UpdateCState(*context.Context, *model.CState) (*model.CState, *errordef.LogicError)
    DeleteCState(*context.Context, int, int) *errordef.LogicError
    GetCStateWithKey(*context.Context, int,int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithStateCdIso(*context.Context, string) ([]*model.CState, *errordef.LogicError)
    GetCStateWithLanguageCd(*context.Context, int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithCountryCd(*context.Context, int,int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithStateCd(*context.Context, int) ([]*model.CState, *errordef.LogicError)
}
