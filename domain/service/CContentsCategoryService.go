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

type CContentsCategoryService interface {
    CreateCContentsCategory(*context.Context, *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError)
    UpdateCContentsCategory(*context.Context, *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError)
    GetCContentsCategoryWithKey(*context.Context, int,int) ([]*model.CContentsCategory, *errordef.LogicError)
    GetCContentsCategoryWithLanguageCd(*context.Context, int) ([]*model.CContentsCategory, *errordef.LogicError)
    DeleteCContentsCategory(*context.Context, int, int) *errordef.LogicError
}

type cContentsCategoryService struct {
    cContentsCategoryRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCContentsCategoryService(pr repository.Repository, nu number.NumberUtil) CContentsCategoryService {
    return &cContentsCategoryService{
        cContentsCategoryRepository :pr,
        numberUtil :nu,
    }
}

func (ss cContentsCategoryService) CreateCContentsCategory(ctx *context.Context, cContentsCategory *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError) {
    cContentsCategory.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsCategory.CreateFunction = "CreateCContentsCategory"
    cContentsCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsCategory.UpdateFunction = "CreateCContentsCategory"
    cContentsCategoryRepository := ss.cContentsCategoryRepository
    created, err := cContentsCategoryRepository.CreateCContentsCategory(ctx, cContentsCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cContentsCategoryService) UpdateCContentsCategory(ctx *context.Context, cContentsCategory *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError) {
    cContentsCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cContentsCategory.UpdateFunction = "UpdateCContentsCategory"
    cContentsCategoryRepository := ss.cContentsCategoryRepository
    results, err := cContentsCategoryRepository.GetCContentsCategoryWithKey(ctx, cContentsCategory.ContentsCategoryCd, cContentsCategory.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cContentsCategoryRepository.UpdateCContentsCategory(ctx, cContentsCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cContentsCategoryService) DeleteCContentsCategory(ctx *context.Context, contentsCategoryCd int, languageCd int) *errordef.LogicError {
    cContentsCategoryRepository := ss.cContentsCategoryRepository
    results, err := cContentsCategoryRepository.GetCContentsCategoryWithKey(ctx, contentsCategoryCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cContentsCategoryRepository.DeleteCContentsCategory(ctx, contentsCategoryCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cContentsCategoryService) GetCContentsCategoryWithKey(ctx *context.Context, contentsCategoryCd int,languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
    cContentsCategoryRepository := ss.cContentsCategoryRepository
    cContentsCategory, err := cContentsCategoryRepository.GetCContentsCategoryWithKey(ctx, contentsCategoryCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsCategory, nil
}

func (ss cContentsCategoryService) GetCContentsCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
    cContentsCategoryRepository := ss.cContentsCategoryRepository
    cContentsCategory, err := cContentsCategoryRepository.GetCContentsCategoryWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cContentsCategory, nil
}

