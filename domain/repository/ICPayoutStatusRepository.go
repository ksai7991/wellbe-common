package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CPayoutStatusRepository interface {
    CreateCPayoutStatus(*context.Context, *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError)
    UpdateCPayoutStatus(*context.Context, *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError)
    DeleteCPayoutStatus(*context.Context, int) *errordef.LogicError
    GetCPayoutStatusWithKey(*context.Context, int) ([]*model.CPayoutStatus, *errordef.LogicError)
    GetCPayoutStatusWithLanguageCd(*context.Context, int) ([]*model.CPayoutStatus, *errordef.LogicError)
}
