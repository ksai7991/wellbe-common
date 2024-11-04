package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type COrderTypeApplication interface {
    GetCOrderTypeWithKey(*context.Context, int) ([]*model.COrderType, *errordef.LogicError)
    GetCOrderTypeWithLanguageCd(*context.Context, int) ([]*model.COrderType, *errordef.LogicError)
}

type cOrderTypeApplication struct {
    cOrderTypeService service.COrderTypeService
    transaction repository.Transaction
}

func NewCOrderTypeApplication(ls service.COrderTypeService, tr repository.Transaction) COrderTypeApplication {
    return &cOrderTypeApplication{
        cOrderTypeService :ls,
        transaction :tr,
    }
}

func (sa cOrderTypeApplication) GetCOrderTypeWithKey(ctx *context.Context, orderTypeCd int) ([]*model.COrderType, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cOrderTypeService.GetCOrderTypeWithKey(ctx, orderTypeCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cOrderTypeApplication) GetCOrderTypeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.COrderType, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cOrderTypeService.GetCOrderTypeWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
