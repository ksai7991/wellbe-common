package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CShopEquipmentApplication interface {
    GetCShopEquipmentWithKey(*context.Context, int,int) ([]*model.CShopEquipment, *errordef.LogicError)
    GetCShopEquipmentWithLanguageCd(*context.Context, int) ([]*model.CShopEquipment, *errordef.LogicError)
}

type cShopEquipmentApplication struct {
    cShopEquipmentService service.CShopEquipmentService
    transaction repository.Transaction
}

func NewCShopEquipmentApplication(ls service.CShopEquipmentService, tr repository.Transaction) CShopEquipmentApplication {
    return &cShopEquipmentApplication{
        cShopEquipmentService :ls,
        transaction :tr,
    }
}

func (sa cShopEquipmentApplication) GetCShopEquipmentWithKey(ctx *context.Context, shopEquipmentCd int,languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopEquipmentService.GetCShopEquipmentWithKey(ctx, shopEquipmentCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cShopEquipmentApplication) GetCShopEquipmentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopEquipmentService.GetCShopEquipmentWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
