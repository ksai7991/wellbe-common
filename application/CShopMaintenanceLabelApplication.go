package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CShopMaintenanceLabelApplication interface {
    GetCShopMaintenanceLabelWithKey(*context.Context, int,int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
    GetCShopMaintenanceLabelWithLanguageCd(*context.Context, int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
}

type cShopMaintenanceLabelApplication struct {
    cShopMaintenanceLabelService service.CShopMaintenanceLabelService
    transaction repository.Transaction
}

func NewCShopMaintenanceLabelApplication(ls service.CShopMaintenanceLabelService, tr repository.Transaction) CShopMaintenanceLabelApplication {
    return &cShopMaintenanceLabelApplication{
        cShopMaintenanceLabelService :ls,
        transaction :tr,
    }
}

func (sa cShopMaintenanceLabelApplication) GetCShopMaintenanceLabelWithKey(ctx *context.Context, shopMaintenanceLabelCd int,languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopMaintenanceLabelService.GetCShopMaintenanceLabelWithKey(ctx, shopMaintenanceLabelCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cShopMaintenanceLabelApplication) GetCShopMaintenanceLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopMaintenanceLabelService.GetCShopMaintenanceLabelWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
