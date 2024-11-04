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

type CReviewContentService interface {
    CreateCReviewContent(*context.Context, *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError)
    UpdateCReviewContent(*context.Context, *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError)
    GetCReviewContentWithKey(*context.Context, int,int) ([]*model.CReviewContent, *errordef.LogicError)
    GetCReviewContentWithLanguageCd(*context.Context, int) ([]*model.CReviewContent, *errordef.LogicError)
    DeleteCReviewContent(*context.Context, int, int) *errordef.LogicError
}

type cReviewContentService struct {
    cReviewContentRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCReviewContentService(pr repository.Repository, nu number.NumberUtil) CReviewContentService {
    return &cReviewContentService{
        cReviewContentRepository :pr,
        numberUtil :nu,
    }
}

func (ss cReviewContentService) CreateCReviewContent(ctx *context.Context, cReviewContent *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError) {
    cReviewContent.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewContent.CreateFunction = "CreateCReviewContent"
    cReviewContent.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewContent.UpdateFunction = "CreateCReviewContent"
    cReviewContentRepository := ss.cReviewContentRepository
    created, err := cReviewContentRepository.CreateCReviewContent(ctx, cReviewContent)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cReviewContentService) UpdateCReviewContent(ctx *context.Context, cReviewContent *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError) {
    cReviewContent.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewContent.UpdateFunction = "UpdateCReviewContent"
    cReviewContentRepository := ss.cReviewContentRepository
    results, err := cReviewContentRepository.GetCReviewContentWithKey(ctx, cReviewContent.ReviewContentCd, cReviewContent.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cReviewContentRepository.UpdateCReviewContent(ctx, cReviewContent)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cReviewContentService) DeleteCReviewContent(ctx *context.Context, reviewContentCd int, languageCd int) *errordef.LogicError {
    cReviewContentRepository := ss.cReviewContentRepository
    results, err := cReviewContentRepository.GetCReviewContentWithKey(ctx, reviewContentCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cReviewContentRepository.DeleteCReviewContent(ctx, reviewContentCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cReviewContentService) GetCReviewContentWithKey(ctx *context.Context, reviewContentCd int,languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
    cReviewContentRepository := ss.cReviewContentRepository
    cReviewContent, err := cReviewContentRepository.GetCReviewContentWithKey(ctx, reviewContentCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cReviewContent, nil
}

func (ss cReviewContentService) GetCReviewContentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
    cReviewContentRepository := ss.cReviewContentRepository
    cReviewContent, err := cReviewContentRepository.GetCReviewContentWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cReviewContent, nil
}

