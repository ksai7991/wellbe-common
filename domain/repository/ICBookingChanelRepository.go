package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CBookingChanelRepository interface {
    CreateCBookingChanel(*context.Context, *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError)
    UpdateCBookingChanel(*context.Context, *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError)
    DeleteCBookingChanel(*context.Context, int, int) *errordef.LogicError
    GetCBookingChanelWithKey(*context.Context, int,int) ([]*model.CBookingChanel, *errordef.LogicError)
    GetCBookingChanelWithLanguageCd(*context.Context, int) ([]*model.CBookingChanel, *errordef.LogicError)
}
