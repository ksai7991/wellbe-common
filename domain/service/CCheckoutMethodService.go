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

type CCheckoutMethodService interface {
    CreateCCheckoutMethod(*context.Context, *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError)
    UpdateCCheckoutMethod(*context.Context, *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError)
    GetCCheckoutMethodWithKey(*context.Context, int,int) ([]*model.CCheckoutMethod, *errordef.LogicError)
    GetCCheckoutMethodWithLanguageCd(*context.Context, int) ([]*model.CCheckoutMethod, *errordef.LogicError)
    DeleteCCheckoutMethod(*context.Context, int, int) *errordef.LogicError
}

type cCheckoutMethodService struct {
    cCheckoutMethodRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCCheckoutMethodService(pr repository.Repository, nu number.NumberUtil) CCheckoutMethodService {
    return &cCheckoutMethodService{
        cCheckoutMethodRepository :pr,
        numberUtil :nu,
    }
}

func (ss cCheckoutMethodService) CreateCCheckoutMethod(ctx *context.Context, cCheckoutMethod *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError) {
    cCheckoutMethod.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutMethod.CreateFunction = "CreateCCheckoutMethod"
    cCheckoutMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutMethod.UpdateFunction = "CreateCCheckoutMethod"
    cCheckoutMethodRepository := ss.cCheckoutMethodRepository
    created, err := cCheckoutMethodRepository.CreateCCheckoutMethod(ctx, cCheckoutMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cCheckoutMethodService) UpdateCCheckoutMethod(ctx *context.Context, cCheckoutMethod *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError) {
    cCheckoutMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutMethod.UpdateFunction = "UpdateCCheckoutMethod"
    cCheckoutMethodRepository := ss.cCheckoutMethodRepository
    results, err := cCheckoutMethodRepository.GetCCheckoutMethodWithKey(ctx, cCheckoutMethod.CheckoutMethodCd, cCheckoutMethod.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cCheckoutMethodRepository.UpdateCCheckoutMethod(ctx, cCheckoutMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cCheckoutMethodService) DeleteCCheckoutMethod(ctx *context.Context, checkoutMethodCd int, languageCd int) *errordef.LogicError {
    cCheckoutMethodRepository := ss.cCheckoutMethodRepository
    results, err := cCheckoutMethodRepository.GetCCheckoutMethodWithKey(ctx, checkoutMethodCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cCheckoutMethodRepository.DeleteCCheckoutMethod(ctx, checkoutMethodCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cCheckoutMethodService) GetCCheckoutMethodWithKey(ctx *context.Context, checkoutMethodCd int,languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
    cCheckoutMethodRepository := ss.cCheckoutMethodRepository
    cCheckoutMethod, err := cCheckoutMethodRepository.GetCCheckoutMethodWithKey(ctx, checkoutMethodCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCheckoutMethod, nil
}

func (ss cCheckoutMethodService) GetCCheckoutMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
    cCheckoutMethodRepository := ss.cCheckoutMethodRepository
    cCheckoutMethod, err := cCheckoutMethodRepository.GetCCheckoutMethodWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCheckoutMethod, nil
}

