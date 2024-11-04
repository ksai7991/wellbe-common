package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CCouponTargetAttrApplication interface {
    GetCCouponTargetAttrWithKey(*context.Context, int,int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
    GetCCouponTargetAttrWithLanguageCd(*context.Context, int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
}

type cCouponTargetAttrApplication struct {
    cCouponTargetAttrService service.CCouponTargetAttrService
    transaction repository.Transaction
}

func NewCCouponTargetAttrApplication(ls service.CCouponTargetAttrService, tr repository.Transaction) CCouponTargetAttrApplication {
    return &cCouponTargetAttrApplication{
        cCouponTargetAttrService :ls,
        transaction :tr,
    }
}

func (sa cCouponTargetAttrApplication) GetCCouponTargetAttrWithKey(ctx *context.Context, couponTargetAttrCd int,languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCouponTargetAttrService.GetCCouponTargetAttrWithKey(ctx, couponTargetAttrCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCouponTargetAttrApplication) GetCCouponTargetAttrWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCouponTargetAttrService.GetCCouponTargetAttrWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
