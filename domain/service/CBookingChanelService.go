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

type CBookingChanelService interface {
    CreateCBookingChanel(*context.Context, *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError)
    UpdateCBookingChanel(*context.Context, *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError)
    GetCBookingChanelWithKey(*context.Context, int,int) ([]*model.CBookingChanel, *errordef.LogicError)
    GetCBookingChanelWithLanguageCd(*context.Context, int) ([]*model.CBookingChanel, *errordef.LogicError)
    DeleteCBookingChanel(*context.Context, int, int) *errordef.LogicError
}

type cBookingChanelService struct {
    cBookingChanelRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCBookingChanelService(pr repository.Repository, nu number.NumberUtil) CBookingChanelService {
    return &cBookingChanelService{
        cBookingChanelRepository :pr,
        numberUtil :nu,
    }
}

func (ss cBookingChanelService) CreateCBookingChanel(ctx *context.Context, cBookingChanel *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError) {
    cBookingChanel.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingChanel.CreateFunction = "CreateCBookingChanel"
    cBookingChanel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingChanel.UpdateFunction = "CreateCBookingChanel"
    cBookingChanelRepository := ss.cBookingChanelRepository
    created, err := cBookingChanelRepository.CreateCBookingChanel(ctx, cBookingChanel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cBookingChanelService) UpdateCBookingChanel(ctx *context.Context, cBookingChanel *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError) {
    cBookingChanel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingChanel.UpdateFunction = "UpdateCBookingChanel"
    cBookingChanelRepository := ss.cBookingChanelRepository
    results, err := cBookingChanelRepository.GetCBookingChanelWithKey(ctx, cBookingChanel.BookingChanelCd, cBookingChanel.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cBookingChanelRepository.UpdateCBookingChanel(ctx, cBookingChanel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cBookingChanelService) DeleteCBookingChanel(ctx *context.Context, bookingChanelCd int, languageCd int) *errordef.LogicError {
    cBookingChanelRepository := ss.cBookingChanelRepository
    results, err := cBookingChanelRepository.GetCBookingChanelWithKey(ctx, bookingChanelCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cBookingChanelRepository.DeleteCBookingChanel(ctx, bookingChanelCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cBookingChanelService) GetCBookingChanelWithKey(ctx *context.Context, bookingChanelCd int,languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
    cBookingChanelRepository := ss.cBookingChanelRepository
    cBookingChanel, err := cBookingChanelRepository.GetCBookingChanelWithKey(ctx, bookingChanelCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBookingChanel, nil
}

func (ss cBookingChanelService) GetCBookingChanelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
    cBookingChanelRepository := ss.cBookingChanelRepository
    cBookingChanel, err := cBookingChanelRepository.GetCBookingChanelWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBookingChanel, nil
}

