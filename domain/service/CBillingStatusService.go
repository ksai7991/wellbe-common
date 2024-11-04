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

type CBillingStatusService interface {
    CreateCBillingStatus(*context.Context, *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError)
    UpdateCBillingStatus(*context.Context, *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError)
    GetCBillingStatusWithKey(*context.Context, int,int) ([]*model.CBillingStatus, *errordef.LogicError)
    GetCBillingStatusWithLanguageCd(*context.Context, int) ([]*model.CBillingStatus, *errordef.LogicError)
    DeleteCBillingStatus(*context.Context, int, int) *errordef.LogicError
}

type cBillingStatusService struct {
    cBillingStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCBillingStatusService(pr repository.Repository, nu number.NumberUtil) CBillingStatusService {
    return &cBillingStatusService{
        cBillingStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cBillingStatusService) CreateCBillingStatus(ctx *context.Context, cBillingStatus *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError) {
    cBillingStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBillingStatus.CreateFunction = "CreateCBillingStatus"
    cBillingStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBillingStatus.UpdateFunction = "CreateCBillingStatus"
    cBillingStatusRepository := ss.cBillingStatusRepository
    created, err := cBillingStatusRepository.CreateCBillingStatus(ctx, cBillingStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cBillingStatusService) UpdateCBillingStatus(ctx *context.Context, cBillingStatus *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError) {
    cBillingStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBillingStatus.UpdateFunction = "UpdateCBillingStatus"
    cBillingStatusRepository := ss.cBillingStatusRepository
    results, err := cBillingStatusRepository.GetCBillingStatusWithKey(ctx, cBillingStatus.BillingStatusCd, cBillingStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cBillingStatusRepository.UpdateCBillingStatus(ctx, cBillingStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cBillingStatusService) DeleteCBillingStatus(ctx *context.Context, billingStatusCd int, languageCd int) *errordef.LogicError {
    cBillingStatusRepository := ss.cBillingStatusRepository
    results, err := cBillingStatusRepository.GetCBillingStatusWithKey(ctx, billingStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cBillingStatusRepository.DeleteCBillingStatus(ctx, billingStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cBillingStatusService) GetCBillingStatusWithKey(ctx *context.Context, billingStatusCd int,languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
    cBillingStatusRepository := ss.cBillingStatusRepository
    cBillingStatus, err := cBillingStatusRepository.GetCBillingStatusWithKey(ctx, billingStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBillingStatus, nil
}

func (ss cBillingStatusService) GetCBillingStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
    cBillingStatusRepository := ss.cBillingStatusRepository
    cBillingStatus, err := cBillingStatusRepository.GetCBillingStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBillingStatus, nil
}

