package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CShopPaymentMethodApplication interface {
    GetCShopPaymentMethodWithKey(*context.Context, int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
    GetCShopPaymentMethodWithLanguageCd(*context.Context, int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
}

type cShopPaymentMethodApplication struct {
    cShopPaymentMethodService service.CShopPaymentMethodService
    transaction repository.Transaction
}

func NewCShopPaymentMethodApplication(ls service.CShopPaymentMethodService, tr repository.Transaction) CShopPaymentMethodApplication {
    return &cShopPaymentMethodApplication{
        cShopPaymentMethodService :ls,
        transaction :tr,
    }
}

func (sa cShopPaymentMethodApplication) GetCShopPaymentMethodWithKey(ctx *context.Context, shopPaymentMethodCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopPaymentMethodService.GetCShopPaymentMethodWithKey(ctx, shopPaymentMethodCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cShopPaymentMethodApplication) GetCShopPaymentMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cShopPaymentMethodService.GetCShopPaymentMethodWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
