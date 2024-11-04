package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type DefaultFeeMasterApplication interface {
    CreateDefaultFeeMaster(*context.Context, *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    UpdateDefaultFeeMaster(*context.Context, *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    DeleteDefaultFeeMaster(*context.Context, string) *errordef.LogicError
}

type defaultFeeMasterApplication struct {
    defaultFeeMasterService service.DefaultFeeMasterService
    transaction repository.Transaction
}

func NewDefaultFeeMasterApplication(ls service.DefaultFeeMasterService, tr repository.Transaction) DefaultFeeMasterApplication {
    return &defaultFeeMasterApplication{
        defaultFeeMasterService :ls,
        transaction :tr,
    }
}

func (sa defaultFeeMasterApplication) CreateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.defaultFeeMasterService.CreateDefaultFeeMaster(ctx, defaultFeeMaster)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa defaultFeeMasterApplication) UpdateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.defaultFeeMasterService.UpdateDefaultFeeMaster(ctx, defaultFeeMaster)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa defaultFeeMasterApplication) DeleteDefaultFeeMaster(ctx *context.Context, id string) *errordef.LogicError {
    defer sa.transaction.Rollback(ctx)
    err := sa.defaultFeeMasterService.DeleteDefaultFeeMaster(ctx, id)
    sa.transaction.Commit(ctx)
    return err
}
