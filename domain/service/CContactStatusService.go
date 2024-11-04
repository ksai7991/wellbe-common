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

type CContactStatusService interface {
    CreateCContactStatus(*context.Context, *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError)
    UpdateCContactStatus(*context.Context, *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError)
    GetCContactStatusWithKey(*context.Context, int,int) ([]*model.CContactStatus, *errordef.LogicError)
    GetCContactStatusWithLanguageCd(*context.Context, int) ([]*model.CContactStatus, *errordef.LogicError)
    DeleteCContactStatus(*context.Context, int, int) *errordef.LogicError
}

type cContactStatusService struct {
    cContactStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCContactStatusService(pr repository.Repository, nu number.NumberUtil) CContactStatusService {
    return &cContactStatusService{
        cContactStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cContactStatusService) CreateCContactStatus(ctx *context.Context, cContactStatus *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError) {
    cContactStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContactStatus.CreateFunction = "CreateCContactStatus"
    cContactStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContactStatus.UpdateFunction = "CreateCContactStatus"
    cContactStatusRepository := ss.cContactStatusRepository
    created, err := cContactStatusRepository.CreateCContactStatus(ctx, cContactStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cContactStatusService) UpdateCContactStatus(ctx *context.Context, cContactStatus *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError) {
    cContactStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContactStatus.UpdateFunction = "UpdateCContactStatus"
    cContactStatusRepository := ss.cContactStatusRepository
    results, err := cContactStatusRepository.GetCContactStatusWithKey(ctx, cContactStatus.ContactStatusCd, cContactStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cContactStatusRepository.UpdateCContactStatus(ctx, cContactStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cContactStatusService) DeleteCContactStatus(ctx *context.Context, contactStatusCd int, languageCd int) *errordef.LogicError {
    cContactStatusRepository := ss.cContactStatusRepository
    results, err := cContactStatusRepository.GetCContactStatusWithKey(ctx, contactStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cContactStatusRepository.DeleteCContactStatus(ctx, contactStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cContactStatusService) GetCContactStatusWithKey(ctx *context.Context, contactStatusCd int,languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
    cContactStatusRepository := ss.cContactStatusRepository
    cContactStatus, err := cContactStatusRepository.GetCContactStatusWithKey(ctx, contactStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContactStatus, nil
}

func (ss cContactStatusService) GetCContactStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
    cContactStatusRepository := ss.cContactStatusRepository
    cContactStatus, err := cContactStatusRepository.GetCContactStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContactStatus, nil
}

