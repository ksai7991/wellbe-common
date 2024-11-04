package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CConcernRepository interface {
    CreateCConcern(*context.Context, *model.CConcern) (*model.CConcern, *errordef.LogicError)
    UpdateCConcern(*context.Context, *model.CConcern) (*model.CConcern, *errordef.LogicError)
    DeleteCConcern(*context.Context, int, int) *errordef.LogicError
    GetCConcernWithKey(*context.Context, int,int) ([]*model.CConcern, *errordef.LogicError)
    GetCConcernWithLanguageCd(*context.Context, int) ([]*model.CConcern, *errordef.LogicError)
}
