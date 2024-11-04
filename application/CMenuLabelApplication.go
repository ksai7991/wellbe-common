package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CMenuLabelApplication interface {
    GetCMenuLabelWithKey(*context.Context, int) ([]*model.CMenuLabel, *errordef.LogicError)
    GetCMenuLabelWithLanguageCd(*context.Context, int) ([]*model.CMenuLabel, *errordef.LogicError)
}

type cMenuLabelApplication struct {
    cMenuLabelService service.CMenuLabelService
    transaction repository.Transaction
}

func NewCMenuLabelApplication(ls service.CMenuLabelService, tr repository.Transaction) CMenuLabelApplication {
    return &cMenuLabelApplication{
        cMenuLabelService :ls,
        transaction :tr,
    }
}

func (sa cMenuLabelApplication) GetCMenuLabelWithKey(ctx *context.Context, menuLabelCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cMenuLabelService.GetCMenuLabelWithKey(ctx, menuLabelCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cMenuLabelApplication) GetCMenuLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cMenuLabelService.GetCMenuLabelWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
