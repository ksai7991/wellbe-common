package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CServiceApplication interface {
    GetCServiceWithKey(*context.Context, int,int) ([]*model.CService, *errordef.LogicError)
    GetCServiceWithLanguageCd(*context.Context, int) ([]*model.CService, *errordef.LogicError)
}

type cServiceApplication struct {
    cServiceService service.CServiceService
    transaction repository.Transaction
}

func NewCServiceApplication(ls service.CServiceService, tr repository.Transaction) CServiceApplication {
    return &cServiceApplication{
        cServiceService :ls,
        transaction :tr,
    }
}

func (sa cServiceApplication) GetCServiceWithKey(ctx *context.Context, serviceCd int,languageCd int) ([]*model.CService, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cServiceService.GetCServiceWithKey(ctx, serviceCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cServiceApplication) GetCServiceWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CService, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cServiceService.GetCServiceWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
