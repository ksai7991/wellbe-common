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

type CRecommendLabelService interface {
    CreateCRecommendLabel(*context.Context, *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError)
    UpdateCRecommendLabel(*context.Context, *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError)
    GetCRecommendLabelWithKey(*context.Context, int,int) ([]*model.CRecommendLabel, *errordef.LogicError)
    GetCRecommendLabelWithLanguageCd(*context.Context, int) ([]*model.CRecommendLabel, *errordef.LogicError)
    DeleteCRecommendLabel(*context.Context, int, int) *errordef.LogicError
}

type cRecommendLabelService struct {
    cRecommendLabelRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCRecommendLabelService(pr repository.Repository, nu number.NumberUtil) CRecommendLabelService {
    return &cRecommendLabelService{
        cRecommendLabelRepository :pr,
        numberUtil :nu,
    }
}

func (ss cRecommendLabelService) CreateCRecommendLabel(ctx *context.Context, cRecommendLabel *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError) {
    cRecommendLabel.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cRecommendLabel.CreateFunction = "CreateCRecommendLabel"
    cRecommendLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cRecommendLabel.UpdateFunction = "CreateCRecommendLabel"
    cRecommendLabelRepository := ss.cRecommendLabelRepository
    created, err := cRecommendLabelRepository.CreateCRecommendLabel(ctx, cRecommendLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cRecommendLabelService) UpdateCRecommendLabel(ctx *context.Context, cRecommendLabel *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError) {
    cRecommendLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cRecommendLabel.UpdateFunction = "UpdateCRecommendLabel"
    cRecommendLabelRepository := ss.cRecommendLabelRepository
    results, err := cRecommendLabelRepository.GetCRecommendLabelWithKey(ctx, cRecommendLabel.RecommendLabelCd, cRecommendLabel.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cRecommendLabelRepository.UpdateCRecommendLabel(ctx, cRecommendLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cRecommendLabelService) DeleteCRecommendLabel(ctx *context.Context, recommendLabelCd int, languageCd int) *errordef.LogicError {
    cRecommendLabelRepository := ss.cRecommendLabelRepository
    results, err := cRecommendLabelRepository.GetCRecommendLabelWithKey(ctx, recommendLabelCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cRecommendLabelRepository.DeleteCRecommendLabel(ctx, recommendLabelCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cRecommendLabelService) GetCRecommendLabelWithKey(ctx *context.Context, recommendLabelCd int,languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
    cRecommendLabelRepository := ss.cRecommendLabelRepository
    cRecommendLabel, err := cRecommendLabelRepository.GetCRecommendLabelWithKey(ctx, recommendLabelCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cRecommendLabel, nil
}

func (ss cRecommendLabelService) GetCRecommendLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
    cRecommendLabelRepository := ss.cRecommendLabelRepository
    cRecommendLabel, err := cRecommendLabelRepository.GetCRecommendLabelWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cRecommendLabel, nil
}

