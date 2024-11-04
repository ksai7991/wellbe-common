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

type CContentsLabelService interface {
    CreateCContentsLabel(*context.Context, *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError)
    UpdateCContentsLabel(*context.Context, *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithKey(*context.Context, int,int) ([]*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithLanguageCd(*context.Context, int) ([]*model.CContentsLabel, *errordef.LogicError)
    GetCContentsLabelWithContentsCateogry(*context.Context, int,int) ([]*model.CContentsLabel, *errordef.LogicError)
    DeleteCContentsLabel(*context.Context, int, int) *errordef.LogicError
}

type cContentsLabelService struct {
    cContentsLabelRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCContentsLabelService(pr repository.Repository, nu number.NumberUtil) CContentsLabelService {
    return &cContentsLabelService{
        cContentsLabelRepository :pr,
        numberUtil :nu,
    }
}

func (ss cContentsLabelService) CreateCContentsLabel(ctx *context.Context, cContentsLabel *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError) {
    cContentsLabel.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsLabel.CreateFunction = "CreateCContentsLabel"
    cContentsLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsLabel.UpdateFunction = "CreateCContentsLabel"
    cContentsLabelRepository := ss.cContentsLabelRepository
    created, err := cContentsLabelRepository.CreateCContentsLabel(ctx, cContentsLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cContentsLabelService) UpdateCContentsLabel(ctx *context.Context, cContentsLabel *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError) {
    cContentsLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsLabel.UpdateFunction = "UpdateCContentsLabel"
    cContentsLabelRepository := ss.cContentsLabelRepository
    results, err := cContentsLabelRepository.GetCContentsLabelWithKey(ctx, cContentsLabel.ContentsLabelCd, cContentsLabel.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cContentsLabelRepository.UpdateCContentsLabel(ctx, cContentsLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cContentsLabelService) DeleteCContentsLabel(ctx *context.Context, contentsLabelCd int, languageCd int) *errordef.LogicError {
    cContentsLabelRepository := ss.cContentsLabelRepository
    results, err := cContentsLabelRepository.GetCContentsLabelWithKey(ctx, contentsLabelCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cContentsLabelRepository.DeleteCContentsLabel(ctx, contentsLabelCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cContentsLabelService) GetCContentsLabelWithKey(ctx *context.Context, contentsLabelCd int,languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    cContentsLabelRepository := ss.cContentsLabelRepository
    cContentsLabel, err := cContentsLabelRepository.GetCContentsLabelWithKey(ctx, contentsLabelCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsLabel, nil
}

func (ss cContentsLabelService) GetCContentsLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    cContentsLabelRepository := ss.cContentsLabelRepository
    cContentsLabel, err := cContentsLabelRepository.GetCContentsLabelWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsLabel, nil
}

func (ss cContentsLabelService) GetCContentsLabelWithContentsCateogry(ctx *context.Context, languageCd int,contentsCategoryCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
    cContentsLabelRepository := ss.cContentsLabelRepository
    cContentsLabel, err := cContentsLabelRepository.GetCContentsLabelWithContentsCateogry(ctx, languageCd,contentsCategoryCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsLabel, nil
}

