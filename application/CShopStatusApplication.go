package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CShopStatusApplication interface {
    GetCShopStatusWithKey(*context.Context, int,int) ([]*model.CShopStatus, *errordef.LogicError)
    GetCShopStatusWithLanguageCd(*context.Context, int) ([]*model.CShopStatus, *errordef.LogicError)
}

type cShopStatusApplication struct {
    cShopStatusService service.CShopStatusService
    transaction repository.Transaction
}

func NewCShopStatusApplication(ls service.CShopStatusService, tr repository.Transaction) CShopStatusApplication {
    return &cShopStatusApplication{
        cShopStatusService :ls,
        transaction :tr,
    }
}

func (sa cShopStatusApplication) GetCShopStatusWithKey(ctx *context.Context, shopStatusCd int,languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopStatusService.GetCShopStatusWithKey(ctx, shopStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cShopStatusApplication) GetCShopStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopStatusService.GetCShopStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
