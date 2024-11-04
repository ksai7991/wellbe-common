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

type CShopImageFilterCategoryService interface {
    CreateCShopImageFilterCategory(*context.Context, *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError)
    UpdateCShopImageFilterCategory(*context.Context, *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError)
    GetCShopImageFilterCategoryWithKey(*context.Context, int,int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
    GetCShopImageFilterCategoryWithLanguageCd(*context.Context, int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
    DeleteCShopImageFilterCategory(*context.Context, int, int) *errordef.LogicError
}

type cShopImageFilterCategoryService struct {
    cShopImageFilterCategoryRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCShopImageFilterCategoryService(pr repository.Repository, nu number.NumberUtil) CShopImageFilterCategoryService {
    return &cShopImageFilterCategoryService{
        cShopImageFilterCategoryRepository :pr,
        numberUtil :nu,
    }
}

func (ss cShopImageFilterCategoryService) CreateCShopImageFilterCategory(ctx *context.Context, cShopImageFilterCategory *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError) {
    cShopImageFilterCategory.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopImageFilterCategory.CreateFunction = "CreateCShopImageFilterCategory"
    cShopImageFilterCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopImageFilterCategory.UpdateFunction = "CreateCShopImageFilterCategory"
    cShopImageFilterCategoryRepository := ss.cShopImageFilterCategoryRepository
    created, err := cShopImageFilterCategoryRepository.CreateCShopImageFilterCategory(ctx, cShopImageFilterCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cShopImageFilterCategoryService) UpdateCShopImageFilterCategory(ctx *context.Context, cShopImageFilterCategory *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError) {
    cShopImageFilterCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopImageFilterCategory.UpdateFunction = "UpdateCShopImageFilterCategory"
    cShopImageFilterCategoryRepository := ss.cShopImageFilterCategoryRepository
    results, err := cShopImageFilterCategoryRepository.GetCShopImageFilterCategoryWithKey(ctx, cShopImageFilterCategory.ShopImageFilterCategoryCd, cShopImageFilterCategory.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cShopImageFilterCategoryRepository.UpdateCShopImageFilterCategory(ctx, cShopImageFilterCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cShopImageFilterCategoryService) DeleteCShopImageFilterCategory(ctx *context.Context, shopImageFilterCategoryCd int, languageCd int) *errordef.LogicError {
    cShopImageFilterCategoryRepository := ss.cShopImageFilterCategoryRepository
    results, err := cShopImageFilterCategoryRepository.GetCShopImageFilterCategoryWithKey(ctx, shopImageFilterCategoryCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cShopImageFilterCategoryRepository.DeleteCShopImageFilterCategory(ctx, shopImageFilterCategoryCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cShopImageFilterCategoryService) GetCShopImageFilterCategoryWithKey(ctx *context.Context, shopImageFilterCategoryCd int,languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    cShopImageFilterCategoryRepository := ss.cShopImageFilterCategoryRepository
    cShopImageFilterCategory, err := cShopImageFilterCategoryRepository.GetCShopImageFilterCategoryWithKey(ctx, shopImageFilterCategoryCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopImageFilterCategory, nil
}

func (ss cShopImageFilterCategoryService) GetCShopImageFilterCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    cShopImageFilterCategoryRepository := ss.cShopImageFilterCategoryRepository
    cShopImageFilterCategory, err := cShopImageFilterCategoryRepository.GetCShopImageFilterCategoryWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopImageFilterCategory, nil
}

