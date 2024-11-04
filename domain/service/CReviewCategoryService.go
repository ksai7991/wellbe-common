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

type CReviewCategoryService interface {
    CreateCReviewCategory(*context.Context, *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError)
    UpdateCReviewCategory(*context.Context, *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError)
    GetCReviewCategoryWithKey(*context.Context, int,int) ([]*model.CReviewCategory, *errordef.LogicError)
    GetCReviewCategoryWithLanguageCd(*context.Context, int) ([]*model.CReviewCategory, *errordef.LogicError)
    DeleteCReviewCategory(*context.Context, int, int) *errordef.LogicError
}

type cReviewCategoryService struct {
    cReviewCategoryRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCReviewCategoryService(pr repository.Repository, nu number.NumberUtil) CReviewCategoryService {
    return &cReviewCategoryService{
        cReviewCategoryRepository :pr,
        numberUtil :nu,
    }
}

func (ss cReviewCategoryService) CreateCReviewCategory(ctx *context.Context, cReviewCategory *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError) {
    cReviewCategory.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewCategory.CreateFunction = "CreateCReviewCategory"
    cReviewCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewCategory.UpdateFunction = "CreateCReviewCategory"
    cReviewCategoryRepository := ss.cReviewCategoryRepository
    created, err := cReviewCategoryRepository.CreateCReviewCategory(ctx, cReviewCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cReviewCategoryService) UpdateCReviewCategory(ctx *context.Context, cReviewCategory *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError) {
    cReviewCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cReviewCategory.UpdateFunction = "UpdateCReviewCategory"
    cReviewCategoryRepository := ss.cReviewCategoryRepository
    results, err := cReviewCategoryRepository.GetCReviewCategoryWithKey(ctx, cReviewCategory.ReviewCategoryCd, cReviewCategory.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cReviewCategoryRepository.UpdateCReviewCategory(ctx, cReviewCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cReviewCategoryService) DeleteCReviewCategory(ctx *context.Context, reviewCategoryCd int, languageCd int) *errordef.LogicError {
    cReviewCategoryRepository := ss.cReviewCategoryRepository
    results, err := cReviewCategoryRepository.GetCReviewCategoryWithKey(ctx, reviewCategoryCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cReviewCategoryRepository.DeleteCReviewCategory(ctx, reviewCategoryCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cReviewCategoryService) GetCReviewCategoryWithKey(ctx *context.Context, reviewCategoryCd int,languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
    cReviewCategoryRepository := ss.cReviewCategoryRepository
    cReviewCategory, err := cReviewCategoryRepository.GetCReviewCategoryWithKey(ctx, reviewCategoryCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cReviewCategory, nil
}

func (ss cReviewCategoryService) GetCReviewCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
    cReviewCategoryRepository := ss.cReviewCategoryRepository
    cReviewCategory, err := cReviewCategoryRepository.GetCReviewCategoryWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cReviewCategory, nil
}

