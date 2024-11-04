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

type CCheckoutStatusService interface {
    CreateCCheckoutStatus(*context.Context, *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError)
    UpdateCCheckoutStatus(*context.Context, *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError)
    GetCCheckoutStatusWithKey(*context.Context, int,int) ([]*model.CCheckoutStatus, *errordef.LogicError)
    GetCCheckoutStatusWithLanguageCd(*context.Context, int) ([]*model.CCheckoutStatus, *errordef.LogicError)
    DeleteCCheckoutStatus(*context.Context, int, int) *errordef.LogicError
}

type cCheckoutStatusService struct {
    cCheckoutStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCCheckoutStatusService(pr repository.Repository, nu number.NumberUtil) CCheckoutStatusService {
    return &cCheckoutStatusService{
        cCheckoutStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cCheckoutStatusService) CreateCCheckoutStatus(ctx *context.Context, cCheckoutStatus *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError) {
    cCheckoutStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutStatus.CreateFunction = "CreateCCheckoutStatus"
    cCheckoutStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutStatus.UpdateFunction = "CreateCCheckoutStatus"
    cCheckoutStatusRepository := ss.cCheckoutStatusRepository
    created, err := cCheckoutStatusRepository.CreateCCheckoutStatus(ctx, cCheckoutStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cCheckoutStatusService) UpdateCCheckoutStatus(ctx *context.Context, cCheckoutStatus *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError) {
    cCheckoutStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCheckoutStatus.UpdateFunction = "UpdateCCheckoutStatus"
    cCheckoutStatusRepository := ss.cCheckoutStatusRepository
    results, err := cCheckoutStatusRepository.GetCCheckoutStatusWithKey(ctx, cCheckoutStatus.CheckoutStatusCd, cCheckoutStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cCheckoutStatusRepository.UpdateCCheckoutStatus(ctx, cCheckoutStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cCheckoutStatusService) DeleteCCheckoutStatus(ctx *context.Context, checkoutStatusCd int, languageCd int) *errordef.LogicError {
    cCheckoutStatusRepository := ss.cCheckoutStatusRepository
    results, err := cCheckoutStatusRepository.GetCCheckoutStatusWithKey(ctx, checkoutStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cCheckoutStatusRepository.DeleteCCheckoutStatus(ctx, checkoutStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cCheckoutStatusService) GetCCheckoutStatusWithKey(ctx *context.Context, checkoutStatusCd int,languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
    cCheckoutStatusRepository := ss.cCheckoutStatusRepository
    cCheckoutStatus, err := cCheckoutStatusRepository.GetCCheckoutStatusWithKey(ctx, checkoutStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCheckoutStatus, nil
}

func (ss cCheckoutStatusService) GetCCheckoutStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
    cCheckoutStatusRepository := ss.cCheckoutStatusRepository
    cCheckoutStatus, err := cCheckoutStatusRepository.GetCCheckoutStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCheckoutStatus, nil
}

