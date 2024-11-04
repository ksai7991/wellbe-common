package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CAreaApplication interface {
    GetCAreaWithKey(*context.Context, int,int) ([]*model.CArea, *errordef.LogicError)
    GetCAreaWithLanguageCd(*context.Context, int) ([]*model.CArea, *errordef.LogicError)
    GetCAreaWithStateCd(*context.Context, int) ([]*model.CArea, *errordef.LogicError)
}

type cAreaApplication struct {
    cAreaService service.CAreaService
    transaction repository.Transaction
}

func NewCAreaApplication(ls service.CAreaService, tr repository.Transaction) CAreaApplication {
    return &cAreaApplication{
        cAreaService :ls,
        transaction :tr,
    }
}

func (sa cAreaApplication) GetCAreaWithKey(ctx *context.Context, languageCd int,areaCd int) ([]*model.CArea, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAreaService.GetCAreaWithKey(ctx, languageCd,areaCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cAreaApplication) GetCAreaWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CArea, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAreaService.GetCAreaWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cAreaApplication) GetCAreaWithStateCd(ctx *context.Context, stateCd int) ([]*model.CArea, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAreaService.GetCAreaWithStateCd(ctx, stateCd)
    sa.transaction.Commit(ctx)
    return result, err
}
