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

type CReviewStatusService interface {
    CreateCReviewStatus(*context.Context, *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError)
    UpdateCReviewStatus(*context.Context, *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError)
    GetCReviewStatusWithKey(*context.Context, int,int) ([]*model.CReviewStatus, *errordef.LogicError)
    GetCReviewStatusWithLanguageCd(*context.Context, int) ([]*model.CReviewStatus, *errordef.LogicError)
    DeleteCReviewStatus(*context.Context, int, int) *errordef.LogicError
}

type cReviewStatusService struct {
    cReviewStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCReviewStatusService(pr repository.Repository, nu number.NumberUtil) CReviewStatusService {
    return &cReviewStatusService{
        cReviewStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cReviewStatusService) CreateCReviewStatus(ctx *context.Context, cReviewStatus *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError) {
    cReviewStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewStatus.CreateFunction = "CreateCReviewStatus"
    cReviewStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewStatus.UpdateFunction = "CreateCReviewStatus"
    cReviewStatusRepository := ss.cReviewStatusRepository
    created, err := cReviewStatusRepository.CreateCReviewStatus(ctx, cReviewStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cReviewStatusService) UpdateCReviewStatus(ctx *context.Context, cReviewStatus *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError) {
    cReviewStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewStatus.UpdateFunction = "UpdateCReviewStatus"
    cReviewStatusRepository := ss.cReviewStatusRepository
    results, err := cReviewStatusRepository.GetCReviewStatusWithKey(ctx, cReviewStatus.ReviewStatusCd, cReviewStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cReviewStatusRepository.UpdateCReviewStatus(ctx, cReviewStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cReviewStatusService) DeleteCReviewStatus(ctx *context.Context, reviewStatusCd int, languageCd int) *errordef.LogicError {
    cReviewStatusRepository := ss.cReviewStatusRepository
    results, err := cReviewStatusRepository.GetCReviewStatusWithKey(ctx, reviewStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cReviewStatusRepository.DeleteCReviewStatus(ctx, reviewStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cReviewStatusService) GetCReviewStatusWithKey(ctx *context.Context, reviewStatusCd int,languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
    cReviewStatusRepository := ss.cReviewStatusRepository
    cReviewStatus, err := cReviewStatusRepository.GetCReviewStatusWithKey(ctx, reviewStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cReviewStatus, nil
}

func (ss cReviewStatusService) GetCReviewStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
    cReviewStatusRepository := ss.cReviewStatusRepository
    cReviewStatus, err := cReviewStatusRepository.GetCReviewStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cReviewStatus, nil
}

