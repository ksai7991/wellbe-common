package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CWeekdayRepository interface {
    CreateCWeekday(*context.Context, *model.CWeekday) (*model.CWeekday, *errordef.LogicError)
    UpdateCWeekday(*context.Context, *model.CWeekday) (*model.CWeekday, *errordef.LogicError)
    DeleteCWeekday(*context.Context, int) *errordef.LogicError
    GetCWeekdayWithKey(*context.Context, int) ([]*model.CWeekday, *errordef.LogicError)
    GetCWeekdayWithLanguageCd(*context.Context, int) ([]*model.CWeekday, *errordef.LogicError)
}
