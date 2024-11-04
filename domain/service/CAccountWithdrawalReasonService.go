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

type CAccountWithdrawalReasonService interface {
    CreateCAccountWithdrawalReason(*context.Context, *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError)
    UpdateCAccountWithdrawalReason(*context.Context, *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError)
    GetCAccountWithdrawalReasonWithKey(*context.Context, int,int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
    GetCAccountWithdrawalReasonWithLanguageCd(*context.Context, int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
    DeleteCAccountWithdrawalReason(*context.Context, int, int) *errordef.LogicError
}

type cAccountWithdrawalReasonService struct {
    cAccountWithdrawalReasonRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCAccountWithdrawalReasonService(pr repository.Repository, nu number.NumberUtil) CAccountWithdrawalReasonService {
    return &cAccountWithdrawalReasonService{
        cAccountWithdrawalReasonRepository :pr,
        numberUtil :nu,
    }
}

func (ss cAccountWithdrawalReasonService) CreateCAccountWithdrawalReason(ctx *context.Context, cAccountWithdrawalReason *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError) {
    cAccountWithdrawalReason.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cAccountWithdrawalReason.CreateFunction = "CreateCAccountWithdrawalReason"
    cAccountWithdrawalReason.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cAccountWithdrawalReason.UpdateFunction = "CreateCAccountWithdrawalReason"
    cAccountWithdrawalReasonRepository := ss.cAccountWithdrawalReasonRepository
    created, err := cAccountWithdrawalReasonRepository.CreateCAccountWithdrawalReason(ctx, cAccountWithdrawalReason)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cAccountWithdrawalReasonService) UpdateCAccountWithdrawalReason(ctx *context.Context, cAccountWithdrawalReason *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError) {
    cAccountWithdrawalReason.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cAccountWithdrawalReason.UpdateFunction = "UpdateCAccountWithdrawalReason"
    cAccountWithdrawalReasonRepository := ss.cAccountWithdrawalReasonRepository
    results, err := cAccountWithdrawalReasonRepository.GetCAccountWithdrawalReasonWithKey(ctx, cAccountWithdrawalReason.AccountWithdrawalReasonCd, cAccountWithdrawalReason.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cAccountWithdrawalReasonRepository.UpdateCAccountWithdrawalReason(ctx, cAccountWithdrawalReason)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cAccountWithdrawalReasonService) DeleteCAccountWithdrawalReason(ctx *context.Context, accountWithdrawalReasonCd int, languageCd int) *errordef.LogicError {
    cAccountWithdrawalReasonRepository := ss.cAccountWithdrawalReasonRepository
    results, err := cAccountWithdrawalReasonRepository.GetCAccountWithdrawalReasonWithKey(ctx, accountWithdrawalReasonCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cAccountWithdrawalReasonRepository.DeleteCAccountWithdrawalReason(ctx, accountWithdrawalReasonCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cAccountWithdrawalReasonService) GetCAccountWithdrawalReasonWithKey(ctx *context.Context, accountWithdrawalReasonCd int,languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    cAccountWithdrawalReasonRepository := ss.cAccountWithdrawalReasonRepository
    cAccountWithdrawalReason, err := cAccountWithdrawalReasonRepository.GetCAccountWithdrawalReasonWithKey(ctx, accountWithdrawalReasonCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cAccountWithdrawalReason, nil
}

func (ss cAccountWithdrawalReasonService) GetCAccountWithdrawalReasonWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    cAccountWithdrawalReasonRepository := ss.cAccountWithdrawalReasonRepository
    cAccountWithdrawalReason, err := cAccountWithdrawalReasonRepository.GetCAccountWithdrawalReasonWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cAccountWithdrawalReason, nil
}

