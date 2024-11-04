package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CBookingMethodApplication interface {
    GetCBookingMethodWithKey(*context.Context, int,int) ([]*model.CBookingMethod, *errordef.LogicError)
    GetCBookingMethodWithLanguageCd(*context.Context, int) ([]*model.CBookingMethod, *errordef.LogicError)
}

type cBookingMethodApplication struct {
    cBookingMethodService service.CBookingMethodService
    transaction repository.Transaction
}

func NewCBookingMethodApplication(ls service.CBookingMethodService, tr repository.Transaction) CBookingMethodApplication {
    return &cBookingMethodApplication{
        cBookingMethodService :ls,
        transaction :tr,
    }
}

func (sa cBookingMethodApplication) GetCBookingMethodWithKey(ctx *context.Context, bookingMethodCd int,languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBookingMethodService.GetCBookingMethodWithKey(ctx, bookingMethodCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cBookingMethodApplication) GetCBookingMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cBookingMethodService.GetCBookingMethodWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
