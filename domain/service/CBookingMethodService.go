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

type CBookingMethodService interface {
    CreateCBookingMethod(*context.Context, *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError)
    UpdateCBookingMethod(*context.Context, *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError)
    GetCBookingMethodWithKey(*context.Context, int,int) ([]*model.CBookingMethod, *errordef.LogicError)
    GetCBookingMethodWithLanguageCd(*context.Context, int) ([]*model.CBookingMethod, *errordef.LogicError)
    DeleteCBookingMethod(*context.Context, int, int) *errordef.LogicError
}

type cBookingMethodService struct {
    cBookingMethodRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCBookingMethodService(pr repository.Repository, nu number.NumberUtil) CBookingMethodService {
    return &cBookingMethodService{
        cBookingMethodRepository :pr,
        numberUtil :nu,
    }
}

func (ss cBookingMethodService) CreateCBookingMethod(ctx *context.Context, cBookingMethod *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError) {
    cBookingMethod.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingMethod.CreateFunction = "CreateCBookingMethod"
    cBookingMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingMethod.UpdateFunction = "CreateCBookingMethod"
    cBookingMethodRepository := ss.cBookingMethodRepository
    created, err := cBookingMethodRepository.CreateCBookingMethod(ctx, cBookingMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cBookingMethodService) UpdateCBookingMethod(ctx *context.Context, cBookingMethod *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError) {
    cBookingMethod.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingMethod.UpdateFunction = "UpdateCBookingMethod"
    cBookingMethodRepository := ss.cBookingMethodRepository
    results, err := cBookingMethodRepository.GetCBookingMethodWithKey(ctx, cBookingMethod.BookingMethodCd, cBookingMethod.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cBookingMethodRepository.UpdateCBookingMethod(ctx, cBookingMethod)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cBookingMethodService) DeleteCBookingMethod(ctx *context.Context, bookingMethodCd int, languageCd int) *errordef.LogicError {
    cBookingMethodRepository := ss.cBookingMethodRepository
    results, err := cBookingMethodRepository.GetCBookingMethodWithKey(ctx, bookingMethodCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cBookingMethodRepository.DeleteCBookingMethod(ctx, bookingMethodCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cBookingMethodService) GetCBookingMethodWithKey(ctx *context.Context, bookingMethodCd int,languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
    cBookingMethodRepository := ss.cBookingMethodRepository
    cBookingMethod, err := cBookingMethodRepository.GetCBookingMethodWithKey(ctx, bookingMethodCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBookingMethod, nil
}

func (ss cBookingMethodService) GetCBookingMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
    cBookingMethodRepository := ss.cBookingMethodRepository
    cBookingMethod, err := cBookingMethodRepository.GetCBookingMethodWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBookingMethod, nil
}

