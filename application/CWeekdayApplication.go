package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CWeekdayApplication interface {
    GetCWeekdayWithKey(*context.Context, int) ([]*model.CWeekday, *errordef.LogicError)
    GetCWeekdayWithLanguageCd(*context.Context, int) ([]*model.CWeekday, *errordef.LogicError)
}

type cWeekdayApplication struct {
    cWeekdayService service.CWeekdayService
    transaction repository.Transaction
}

func NewCWeekdayApplication(ls service.CWeekdayService, tr repository.Transaction) CWeekdayApplication {
    return &cWeekdayApplication{
        cWeekdayService :ls,
        transaction :tr,
    }
}

func (sa cWeekdayApplication) GetCWeekdayWithKey(ctx *context.Context, weekdayCd int) ([]*model.CWeekday, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cWeekdayService.GetCWeekdayWithKey(ctx, weekdayCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cWeekdayApplication) GetCWeekdayWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CWeekday, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cWeekdayService.GetCWeekdayWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
