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

type CBookingStatusService interface {
    CreateCBookingStatus(*context.Context, *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError)
    UpdateCBookingStatus(*context.Context, *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError)
    GetCBookingStatusWithKey(*context.Context, int,int) ([]*model.CBookingStatus, *errordef.LogicError)
    GetCBookingStatusWithLanguageCd(*context.Context, int) ([]*model.CBookingStatus, *errordef.LogicError)
    DeleteCBookingStatus(*context.Context, int, int) *errordef.LogicError
}

type cBookingStatusService struct {
    cBookingStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCBookingStatusService(pr repository.Repository, nu number.NumberUtil) CBookingStatusService {
    return &cBookingStatusService{
        cBookingStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cBookingStatusService) CreateCBookingStatus(ctx *context.Context, cBookingStatus *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError) {
    cBookingStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingStatus.CreateFunction = "CreateCBookingStatus"
    cBookingStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingStatus.UpdateFunction = "CreateCBookingStatus"
    cBookingStatusRepository := ss.cBookingStatusRepository
    created, err := cBookingStatusRepository.CreateCBookingStatus(ctx, cBookingStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cBookingStatusService) UpdateCBookingStatus(ctx *context.Context, cBookingStatus *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError) {
    cBookingStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBookingStatus.UpdateFunction = "UpdateCBookingStatus"
    cBookingStatusRepository := ss.cBookingStatusRepository
    results, err := cBookingStatusRepository.GetCBookingStatusWithKey(ctx, cBookingStatus.BookingStatusCd, cBookingStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cBookingStatusRepository.UpdateCBookingStatus(ctx, cBookingStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cBookingStatusService) DeleteCBookingStatus(ctx *context.Context, bookingStatusCd int, languageCd int) *errordef.LogicError {
    cBookingStatusRepository := ss.cBookingStatusRepository
    results, err := cBookingStatusRepository.GetCBookingStatusWithKey(ctx, bookingStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cBookingStatusRepository.DeleteCBookingStatus(ctx, bookingStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cBookingStatusService) GetCBookingStatusWithKey(ctx *context.Context, bookingStatusCd int,languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
    cBookingStatusRepository := ss.cBookingStatusRepository
    cBookingStatus, err := cBookingStatusRepository.GetCBookingStatusWithKey(ctx, bookingStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBookingStatus, nil
}

func (ss cBookingStatusService) GetCBookingStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
    cBookingStatusRepository := ss.cBookingStatusRepository
    cBookingStatus, err := cBookingStatusRepository.GetCBookingStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBookingStatus, nil
}

