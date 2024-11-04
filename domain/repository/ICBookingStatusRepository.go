package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CBookingStatusRepository interface {
    CreateCBookingStatus(*context.Context, *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError)
    UpdateCBookingStatus(*context.Context, *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError)
    DeleteCBookingStatus(*context.Context, int, int) *errordef.LogicError
    GetCBookingStatusWithKey(*context.Context, int,int) ([]*model.CBookingStatus, *errordef.LogicError)
    GetCBookingStatusWithLanguageCd(*context.Context, int) ([]*model.CBookingStatus, *errordef.LogicError)
}
