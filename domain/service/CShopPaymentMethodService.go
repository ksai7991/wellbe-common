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

type CShopPaymentMethodService interface {
    CreateCShopPaymentMethod(*context.Context, *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError)
    UpdateCShopPaymentMethod(*context.Context, *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError)
    GetCShopPaymentMethodWithKey(*context.Context, int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
    GetCShopPaymentMethodWithLanguageCd(*context.Context, int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
    DeleteCShopPaymentMethod(*context.Context, int) *errordef.LogicError
}

type cShopPaymentMethodService struct {
    cShopPaymentMethodRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCShopPaymentMethodService(pr repository.Repository, nu number.NumberUtil) CShopPaymentMethodService {
    return &cShopPaymentMethodService{
        cShopPaymentMethodRepository :pr,
        numberUtil :nu,
    }
}

func (ss cShopPaymentMethodService) CreateCShopPaymentMethod(ctx *context.Context, cShopPaymentMethod *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError) {
    cShopPaymentMethod.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopPaymentMethod.CreateFunction = "CreateCShopPaymentMethod"
    cShopPaymentMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopPaymentMethod.UpdateFunction = "CreateCShopPaymentMethod"
    cShopPaymentMethodRepository := ss.cShopPaymentMethodRepository
    created, err := cShopPaymentMethodRepository.CreateCShopPaymentMethod(ctx, cShopPaymentMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cShopPaymentMethodService) UpdateCShopPaymentMethod(ctx *context.Context, cShopPaymentMethod *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError) {
    cShopPaymentMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopPaymentMethod.UpdateFunction = "UpdateCShopPaymentMethod"
    cShopPaymentMethodRepository := ss.cShopPaymentMethodRepository
    results, err := cShopPaymentMethodRepository.GetCShopPaymentMethodWithKey(ctx, cShopPaymentMethod.ShopPaymentMethodCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cShopPaymentMethodRepository.UpdateCShopPaymentMethod(ctx, cShopPaymentMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cShopPaymentMethodService) DeleteCShopPaymentMethod(ctx *context.Context, shopPaymentMethodCd int) *errordef.LogicError {
    cShopPaymentMethodRepository := ss.cShopPaymentMethodRepository
    results, err := cShopPaymentMethodRepository.GetCShopPaymentMethodWithKey(ctx, shopPaymentMethodCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cShopPaymentMethodRepository.DeleteCShopPaymentMethod(ctx, shopPaymentMethodCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cShopPaymentMethodService) GetCShopPaymentMethodWithKey(ctx *context.Context, shopPaymentMethodCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    cShopPaymentMethodRepository := ss.cShopPaymentMethodRepository
    cShopPaymentMethod, err := cShopPaymentMethodRepository.GetCShopPaymentMethodWithKey(ctx, shopPaymentMethodCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopPaymentMethod, nil
}

func (ss cShopPaymentMethodService) GetCShopPaymentMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    cShopPaymentMethodRepository := ss.cShopPaymentMethodRepository
    cShopPaymentMethod, err := cShopPaymentMethodRepository.GetCShopPaymentMethodWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopPaymentMethod, nil
}

