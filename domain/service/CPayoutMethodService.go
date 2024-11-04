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

type CPayoutMethodService interface {
    CreateCPayoutMethod(*context.Context, *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError)
    UpdateCPayoutMethod(*context.Context, *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError)
    GetCPayoutMethodWithKey(*context.Context, int) ([]*model.CPayoutMethod, *errordef.LogicError)
    GetCPayoutMethodWithLanguageCd(*context.Context, int) ([]*model.CPayoutMethod, *errordef.LogicError)
    DeleteCPayoutMethod(*context.Context, int) *errordef.LogicError
}

type cPayoutMethodService struct {
    cPayoutMethodRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCPayoutMethodService(pr repository.Repository, nu number.NumberUtil) CPayoutMethodService {
    return &cPayoutMethodService{
        cPayoutMethodRepository :pr,
        numberUtil :nu,
    }
}

func (ss cPayoutMethodService) CreateCPayoutMethod(ctx *context.Context, cPayoutMethod *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError) {
    cPayoutMethod.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutMethod.CreateFunction = "CreateCPayoutMethod"
    cPayoutMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutMethod.UpdateFunction = "CreateCPayoutMethod"
    cPayoutMethodRepository := ss.cPayoutMethodRepository
    created, err := cPayoutMethodRepository.CreateCPayoutMethod(ctx, cPayoutMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cPayoutMethodService) UpdateCPayoutMethod(ctx *context.Context, cPayoutMethod *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError) {
    cPayoutMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutMethod.UpdateFunction = "UpdateCPayoutMethod"
    cPayoutMethodRepository := ss.cPayoutMethodRepository
    results, err := cPayoutMethodRepository.GetCPayoutMethodWithKey(ctx, cPayoutMethod.PayoutMethodCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cPayoutMethodRepository.UpdateCPayoutMethod(ctx, cPayoutMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cPayoutMethodService) DeleteCPayoutMethod(ctx *context.Context, payoutMethodCd int) *errordef.LogicError {
    cPayoutMethodRepository := ss.cPayoutMethodRepository
    results, err := cPayoutMethodRepository.GetCPayoutMethodWithKey(ctx, payoutMethodCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cPayoutMethodRepository.DeleteCPayoutMethod(ctx, payoutMethodCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cPayoutMethodService) GetCPayoutMethodWithKey(ctx *context.Context, payoutMethodCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
    cPayoutMethodRepository := ss.cPayoutMethodRepository
    cPayoutMethod, err := cPayoutMethodRepository.GetCPayoutMethodWithKey(ctx, payoutMethodCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cPayoutMethod, nil
}

func (ss cPayoutMethodService) GetCPayoutMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
    cPayoutMethodRepository := ss.cPayoutMethodRepository
    cPayoutMethod, err := cPayoutMethodRepository.GetCPayoutMethodWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cPayoutMethod, nil
}

