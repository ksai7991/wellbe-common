package service

import (
    model "wellbe-common/domain/model"
    repository "wellbe-common/domain/repository"
    errordef "wellbe-common/share/errordef"
    number "wellbe-common/share/number"
    commonconstants "wellbe-common/share/commonsettings/constants"
    datetime "wellbe-common/share/datetime"
    messages "wellbe-common/share/messages"

    "context"
)

type CCouponTargetAttrService interface {
    CreateCCouponTargetAttr(*context.Context, *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError)
    UpdateCCouponTargetAttr(*context.Context, *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError)
    GetCCouponTargetAttrWithKey(*context.Context, int,int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
    GetCCouponTargetAttrWithLanguageCd(*context.Context, int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
    DeleteCCouponTargetAttr(*context.Context, int, int) *errordef.LogicError
}

type cCouponTargetAttrService struct {
    cCouponTargetAttrRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCCouponTargetAttrService(pr repository.Repository, nu number.NumberUtil) CCouponTargetAttrService {
    return &cCouponTargetAttrService{
        cCouponTargetAttrRepository :pr,
        numberUtil :nu,
    }
}

func (ss cCouponTargetAttrService) CreateCCouponTargetAttr(ctx *context.Context, cCouponTargetAttr *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError) {
    cCouponTargetAttr.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCouponTargetAttr.CreateFunction = "CreateCCouponTargetAttr"
    cCouponTargetAttr.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCouponTargetAttr.UpdateFunction = "CreateCCouponTargetAttr"
    cCouponTargetAttrRepository := ss.cCouponTargetAttrRepository
    created, err := cCouponTargetAttrRepository.CreateCCouponTargetAttr(ctx, cCouponTargetAttr)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cCouponTargetAttrService) UpdateCCouponTargetAttr(ctx *context.Context, cCouponTargetAttr *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError) {
    cCouponTargetAttr.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCouponTargetAttr.UpdateFunction = "UpdateCCouponTargetAttr"
    cCouponTargetAttrRepository := ss.cCouponTargetAttrRepository
    results, err := cCouponTargetAttrRepository.GetCCouponTargetAttrWithKey(ctx, cCouponTargetAttr.CouponTargetAttrCd, cCouponTargetAttr.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cCouponTargetAttrRepository.UpdateCCouponTargetAttr(ctx, cCouponTargetAttr)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cCouponTargetAttrService) DeleteCCouponTargetAttr(ctx *context.Context, couponTargetAttrCd int, languageCd int) *errordef.LogicError {
    cCouponTargetAttrRepository := ss.cCouponTargetAttrRepository
    results, err := cCouponTargetAttrRepository.GetCCouponTargetAttrWithKey(ctx, couponTargetAttrCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cCouponTargetAttrRepository.DeleteCCouponTargetAttr(ctx, couponTargetAttrCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cCouponTargetAttrService) GetCCouponTargetAttrWithKey(ctx *context.Context, couponTargetAttrCd int,languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    cCouponTargetAttrRepository := ss.cCouponTargetAttrRepository
    cCouponTargetAttr, err := cCouponTargetAttrRepository.GetCCouponTargetAttrWithKey(ctx, couponTargetAttrCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCouponTargetAttr, nil
}

func (ss cCouponTargetAttrService) GetCCouponTargetAttrWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    cCouponTargetAttrRepository := ss.cCouponTargetAttrRepository
    cCouponTargetAttr, err := cCouponTargetAttrRepository.GetCCouponTargetAttrWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCouponTargetAttr, nil
}

