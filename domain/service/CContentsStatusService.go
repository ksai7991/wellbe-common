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

type CContentsStatusService interface {
    CreateCContentsStatus(*context.Context, *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError)
    UpdateCContentsStatus(*context.Context, *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError)
    GetCContentsStatusWithKey(*context.Context, int,int) ([]*model.CContentsStatus, *errordef.LogicError)
    GetCContentsStatusWithLanguageCd(*context.Context, int) ([]*model.CContentsStatus, *errordef.LogicError)
    DeleteCContentsStatus(*context.Context, int, int) *errordef.LogicError
}

type cContentsStatusService struct {
    cContentsStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCContentsStatusService(pr repository.Repository, nu number.NumberUtil) CContentsStatusService {
    return &cContentsStatusService{
        cContentsStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cContentsStatusService) CreateCContentsStatus(ctx *context.Context, cContentsStatus *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError) {
    cContentsStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsStatus.CreateFunction = "CreateCContentsStatus"
    cContentsStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsStatus.UpdateFunction = "CreateCContentsStatus"
    cContentsStatusRepository := ss.cContentsStatusRepository
    created, err := cContentsStatusRepository.CreateCContentsStatus(ctx, cContentsStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cContentsStatusService) UpdateCContentsStatus(ctx *context.Context, cContentsStatus *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError) {
    cContentsStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsStatus.UpdateFunction = "UpdateCContentsStatus"
    cContentsStatusRepository := ss.cContentsStatusRepository
    results, err := cContentsStatusRepository.GetCContentsStatusWithKey(ctx, cContentsStatus.ContentsStatusCd, cContentsStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cContentsStatusRepository.UpdateCContentsStatus(ctx, cContentsStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cContentsStatusService) DeleteCContentsStatus(ctx *context.Context, contentsStatusCd int, languageCd int) *errordef.LogicError {
    cContentsStatusRepository := ss.cContentsStatusRepository
    results, err := cContentsStatusRepository.GetCContentsStatusWithKey(ctx, contentsStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cContentsStatusRepository.DeleteCContentsStatus(ctx, contentsStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cContentsStatusService) GetCContentsStatusWithKey(ctx *context.Context, contentsStatusCd int,languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
    cContentsStatusRepository := ss.cContentsStatusRepository
    cContentsStatus, err := cContentsStatusRepository.GetCContentsStatusWithKey(ctx, contentsStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsStatus, nil
}

func (ss cContentsStatusService) GetCContentsStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
    cContentsStatusRepository := ss.cContentsStatusRepository
    cContentsStatus, err := cContentsStatusRepository.GetCContentsStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsStatus, nil
}

