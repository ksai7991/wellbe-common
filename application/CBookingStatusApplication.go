package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CBookingStatusApplication interface {
    GetCBookingStatusWithKey(*context.Context, int,int) ([]*model.CBookingStatus, *errordef.LogicError)
    GetCBookingStatusWithLanguageCd(*context.Context, int) ([]*model.CBookingStatus, *errordef.LogicError)
}

type cBookingStatusApplication struct {
    cBookingStatusService service.CBookingStatusService
    transaction repository.Transaction
}

func NewCBookingStatusApplication(ls service.CBookingStatusService, tr repository.Transaction) CBookingStatusApplication {
    return &cBookingStatusApplication{
        cBookingStatusService :ls,
        transaction :tr,
    }
}

func (sa cBookingStatusApplication) GetCBookingStatusWithKey(ctx *context.Context, bookingStatusCd int,languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBookingStatusService.GetCBookingStatusWithKey(ctx, bookingStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cBookingStatusApplication) GetCBookingStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBookingStatusService.GetCBookingStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
