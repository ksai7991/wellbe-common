package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CTypeOfContactRepository interface {
    CreateCTypeOfContact(*context.Context, *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError)
    UpdateCTypeOfContact(*context.Context, *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError)
    DeleteCTypeOfContact(*context.Context, int, int) *errordef.LogicError
    GetCTypeOfContactWithKey(*context.Context, int,int) ([]*model.CTypeOfContact, *errordef.LogicError)
    GetCTypeOfContactWithLanguageCd(*context.Context, int) ([]*model.CTypeOfContact, *errordef.LogicError)
}
