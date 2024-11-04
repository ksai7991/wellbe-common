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

type CBillingContentService interface {
    CreateCBillingContent(*context.Context, *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError)
    UpdateCBillingContent(*context.Context, *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError)
    GetCBillingContentWithKey(*context.Context, int,int) ([]*model.CBillingContent, *errordef.LogicError)
    GetCBillingContentWithLanguageCd(*context.Context, int) ([]*model.CBillingContent, *errordef.LogicError)
    DeleteCBillingContent(*context.Context, int, int) *errordef.LogicError
}

type cBillingContentService struct {
    cBillingContentRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCBillingContentService(pr repository.Repository, nu number.NumberUtil) CBillingContentService {
    return &cBillingContentService{
        cBillingContentRepository :pr,
        numberUtil :nu,
    }
}

func (ss cBillingContentService) CreateCBillingContent(ctx *context.Context, cBillingContent *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError) {
    cBillingContent.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBillingContent.CreateFunction = "CreateCBillingContent"
    cBillingContent.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBillingContent.UpdateFunction = "CreateCBillingContent"
    cBillingContentRepository := ss.cBillingContentRepository
    created, err := cBillingContentRepository.CreateCBillingContent(ctx, cBillingContent)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cBillingContentService) UpdateCBillingContent(ctx *context.Context, cBillingContent *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError) {
    cBillingContent.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cBillingContent.UpdateFunction = "UpdateCBillingContent"
    cBillingContentRepository := ss.cBillingContentRepository
    results, err := cBillingContentRepository.GetCBillingContentWithKey(ctx, cBillingContent.BillingContentCd, cBillingContent.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cBillingContentRepository.UpdateCBillingContent(ctx, cBillingContent)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cBillingContentService) DeleteCBillingContent(ctx *context.Context, billingContentCd int, languageCd int) *errordef.LogicError {
    cBillingContentRepository := ss.cBillingContentRepository
    results, err := cBillingContentRepository.GetCBillingContentWithKey(ctx, billingContentCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cBillingContentRepository.DeleteCBillingContent(ctx, billingContentCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cBillingContentService) GetCBillingContentWithKey(ctx *context.Context, billingContentCd int,languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
    cBillingContentRepository := ss.cBillingContentRepository
    cBillingContent, err := cBillingContentRepository.GetCBillingContentWithKey(ctx, billingContentCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBillingContent, nil
}

func (ss cBillingContentService) GetCBillingContentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
    cBillingContentRepository := ss.cBillingContentRepository
    cBillingContent, err := cBillingContentRepository.GetCBillingContentWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cBillingContent, nil
}

