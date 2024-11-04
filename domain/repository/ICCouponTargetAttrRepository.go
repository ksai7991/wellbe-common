package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CCouponTargetAttrRepository interface {
    CreateCCouponTargetAttr(*context.Context, *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError)
    UpdateCCouponTargetAttr(*context.Context, *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError)
    DeleteCCouponTargetAttr(*context.Context, int, int) *errordef.LogicError
    GetCCouponTargetAttrWithKey(*context.Context, int,int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
    GetCCouponTargetAttrWithLanguageCd(*context.Context, int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
}
