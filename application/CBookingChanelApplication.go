package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CBookingChanelApplication interface {
    GetCBookingChanelWithKey(*context.Context, int,int) ([]*model.CBookingChanel, *errordef.LogicError)
    GetCBookingChanelWithLanguageCd(*context.Context, int) ([]*model.CBookingChanel, *errordef.LogicError)
}

type cBookingChanelApplication struct {
    cBookingChanelService service.CBookingChanelService
    transaction repository.Transaction
}

func NewCBookingChanelApplication(ls service.CBookingChanelService, tr repository.Transaction) CBookingChanelApplication {
    return &cBookingChanelApplication{
        cBookingChanelService :ls,
        transaction :tr,
    }
}

func (sa cBookingChanelApplication) GetCBookingChanelWithKey(ctx *context.Context, bookingChanelCd int,languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBookingChanelService.GetCBookingChanelWithKey(ctx, bookingChanelCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cBookingChanelApplication) GetCBookingChanelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBookingChanelService.GetCBookingChanelWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
