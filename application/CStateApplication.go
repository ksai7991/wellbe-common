package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CStateApplication interface {
    GetCStateWithKey(*context.Context, int,int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithStateCdIso(*context.Context, string) ([]*model.CState, *errordef.LogicError)
    GetCStateWithLanguageCd(*context.Context, int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithCountryCd(*context.Context, int,int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithStateCd(*context.Context, int) ([]*model.CState, *errordef.LogicError)
}

type cStateApplication struct {
    cStateService service.CStateService
    transaction repository.Transaction
}

func NewCStateApplication(ls service.CStateService, tr repository.Transaction) CStateApplication {
    return &cStateApplication{
        cStateService :ls,
        transaction :tr,
    }
}

func (sa cStateApplication) GetCStateWithKey(ctx *context.Context, stateCd int,languageCd int) ([]*model.CState, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cStateService.GetCStateWithKey(ctx, stateCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cStateApplication) GetCStateWithStateCdIso(ctx *context.Context, stateCdIso string) ([]*model.CState, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cStateService.GetCStateWithStateCdIso(ctx, stateCdIso)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cStateApplication) GetCStateWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CState, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cStateService.GetCStateWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cStateApplication) GetCStateWithCountryCd(ctx *context.Context, languageCd int,countryCd int) ([]*model.CState, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cStateService.GetCStateWithCountryCd(ctx, languageCd,countryCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cStateApplication) GetCStateWithStateCd(ctx *context.Context, stateCd int) ([]*model.CState, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cStateService.GetCStateWithStateCd(ctx, stateCd)
    sa.transaction.Commit(ctx)
    return result, err
}
