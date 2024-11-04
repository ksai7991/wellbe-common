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

type CCheckoutTimingService interface {
    CreateCCheckoutTiming(*context.Context, *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError)
    UpdateCCheckoutTiming(*context.Context, *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError)
    GetCCheckoutTimingWithKey(*context.Context, int,int) ([]*model.CCheckoutTiming, *errordef.LogicError)
    GetCCheckoutTimingWithLanguageCd(*context.Context, int) ([]*model.CCheckoutTiming, *errordef.LogicError)
    DeleteCCheckoutTiming(*context.Context, int, int) *errordef.LogicError
}

type cCheckoutTimingService struct {
    cCheckoutTimingRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCCheckoutTimingService(pr repository.Repository, nu number.NumberUtil) CCheckoutTimingService {
    return &cCheckoutTimingService{
        cCheckoutTimingRepository :pr,
        numberUtil :nu,
    }
}

func (ss cCheckoutTimingService) CreateCCheckoutTiming(ctx *context.Context, cCheckoutTiming *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError) {
    cCheckoutTiming.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutTiming.CreateFunction = "CreateCCheckoutTiming"
    cCheckoutTiming.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutTiming.UpdateFunction = "CreateCCheckoutTiming"
    cCheckoutTimingRepository := ss.cCheckoutTimingRepository
    created, err := cCheckoutTimingRepository.CreateCCheckoutTiming(ctx, cCheckoutTiming)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cCheckoutTimingService) UpdateCCheckoutTiming(ctx *context.Context, cCheckoutTiming *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError) {
    cCheckoutTiming.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutTiming.UpdateFunction = "UpdateCCheckoutTiming"
    cCheckoutTimingRepository := ss.cCheckoutTimingRepository
    results, err := cCheckoutTimingRepository.GetCCheckoutTimingWithKey(ctx, cCheckoutTiming.CheckoutTimingCd, cCheckoutTiming.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cCheckoutTimingRepository.UpdateCCheckoutTiming(ctx, cCheckoutTiming)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cCheckoutTimingService) DeleteCCheckoutTiming(ctx *context.Context, checkoutTimingCd int, languageCd int) *errordef.LogicError {
    cCheckoutTimingRepository := ss.cCheckoutTimingRepository
    results, err := cCheckoutTimingRepository.GetCCheckoutTimingWithKey(ctx, checkoutTimingCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cCheckoutTimingRepository.DeleteCCheckoutTiming(ctx, checkoutTimingCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cCheckoutTimingService) GetCCheckoutTimingWithKey(ctx *context.Context, checkoutTimingCd int,languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
    cCheckoutTimingRepository := ss.cCheckoutTimingRepository
    cCheckoutTiming, err := cCheckoutTimingRepository.GetCCheckoutTimingWithKey(ctx, checkoutTimingCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCheckoutTiming, nil
}

func (ss cCheckoutTimingService) GetCCheckoutTimingWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
    cCheckoutTimingRepository := ss.cCheckoutTimingRepository
    cCheckoutTiming, err := cCheckoutTimingRepository.GetCCheckoutTimingWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCheckoutTiming, nil
}

