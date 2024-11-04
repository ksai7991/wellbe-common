package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CCheckoutTimingRepository interface {
    CreateCCheckoutTiming(*context.Context, *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError)
    UpdateCCheckoutTiming(*context.Context, *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError)
    DeleteCCheckoutTiming(*context.Context, int, int) *errordef.LogicError
    GetCCheckoutTimingWithKey(*context.Context, int,int) ([]*model.CCheckoutTiming, *errordef.LogicError)
    GetCCheckoutTimingWithLanguageCd(*context.Context, int) ([]*model.CCheckoutTiming, *errordef.LogicError)
}
