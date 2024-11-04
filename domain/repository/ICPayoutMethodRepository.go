package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CPayoutMethodRepository interface {
    CreateCPayoutMethod(*context.Context, *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError)
    UpdateCPayoutMethod(*context.Context, *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError)
    DeleteCPayoutMethod(*context.Context, int) *errordef.LogicError
    GetCPayoutMethodWithKey(*context.Context, int) ([]*model.CPayoutMethod, *errordef.LogicError)
    GetCPayoutMethodWithLanguageCd(*context.Context, int) ([]*model.CPayoutMethod, *errordef.LogicError)
}
