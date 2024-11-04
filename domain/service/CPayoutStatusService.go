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

type CPayoutStatusService interface {
    CreateCPayoutStatus(*context.Context, *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError)
    UpdateCPayoutStatus(*context.Context, *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError)
    GetCPayoutStatusWithKey(*context.Context, int) ([]*model.CPayoutStatus, *errordef.LogicError)
    GetCPayoutStatusWithLanguageCd(*context.Context, int) ([]*model.CPayoutStatus, *errordef.LogicError)
    DeleteCPayoutStatus(*context.Context, int) *errordef.LogicError
}

type cPayoutStatusService struct {
    cPayoutStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCPayoutStatusService(pr repository.Repository, nu number.NumberUtil) CPayoutStatusService {
    return &cPayoutStatusService{
        cPayoutStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cPayoutStatusService) CreateCPayoutStatus(ctx *context.Context, cPayoutStatus *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError) {
    cPayoutStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutStatus.CreateFunction = "CreateCPayoutStatus"
    cPayoutStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutStatus.UpdateFunction = "CreateCPayoutStatus"
    cPayoutStatusRepository := ss.cPayoutStatusRepository
    created, err := cPayoutStatusRepository.CreateCPayoutStatus(ctx, cPayoutStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cPayoutStatusService) UpdateCPayoutStatus(ctx *context.Context, cPayoutStatus *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError) {
    cPayoutStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutStatus.UpdateFunction = "UpdateCPayoutStatus"
    cPayoutStatusRepository := ss.cPayoutStatusRepository
    results, err := cPayoutStatusRepository.GetCPayoutStatusWithKey(ctx, cPayoutStatus.PayoutStatusCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cPayoutStatusRepository.UpdateCPayoutStatus(ctx, cPayoutStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cPayoutStatusService) DeleteCPayoutStatus(ctx *context.Context, payoutStatusCd int) *errordef.LogicError {
    cPayoutStatusRepository := ss.cPayoutStatusRepository
    results, err := cPayoutStatusRepository.GetCPayoutStatusWithKey(ctx, payoutStatusCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cPayoutStatusRepository.DeleteCPayoutStatus(ctx, payoutStatusCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cPayoutStatusService) GetCPayoutStatusWithKey(ctx *context.Context, payoutStatusCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
    cPayoutStatusRepository := ss.cPayoutStatusRepository
    cPayoutStatus, err := cPayoutStatusRepository.GetCPayoutStatusWithKey(ctx, payoutStatusCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cPayoutStatus, nil
}

func (ss cPayoutStatusService) GetCPayoutStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
    cPayoutStatusRepository := ss.cPayoutStatusRepository
    cPayoutStatus, err := cPayoutStatusRepository.GetCPayoutStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cPayoutStatus, nil
}

