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

type CPayoutItemCategoryService interface {
    CreateCPayoutItemCategory(*context.Context, *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError)
    UpdateCPayoutItemCategory(*context.Context, *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError)
    GetCPayoutItemCategoryWithKey(*context.Context, int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
    GetCPayoutItemCategoryWithLanguageCd(*context.Context, int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
    DeleteCPayoutItemCategory(*context.Context, int) *errordef.LogicError
}

type cPayoutItemCategoryService struct {
    cPayoutItemCategoryRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCPayoutItemCategoryService(pr repository.Repository, nu number.NumberUtil) CPayoutItemCategoryService {
    return &cPayoutItemCategoryService{
        cPayoutItemCategoryRepository :pr,
        numberUtil :nu,
    }
}

func (ss cPayoutItemCategoryService) CreateCPayoutItemCategory(ctx *context.Context, cPayoutItemCategory *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError) {
    cPayoutItemCategory.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutItemCategory.CreateFunction = "CreateCPayoutItemCategory"
    cPayoutItemCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutItemCategory.UpdateFunction = "CreateCPayoutItemCategory"
    cPayoutItemCategoryRepository := ss.cPayoutItemCategoryRepository
    created, err := cPayoutItemCategoryRepository.CreateCPayoutItemCategory(ctx, cPayoutItemCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cPayoutItemCategoryService) UpdateCPayoutItemCategory(ctx *context.Context, cPayoutItemCategory *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError) {
    cPayoutItemCategory.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cPayoutItemCategory.UpdateFunction = "UpdateCPayoutItemCategory"
    cPayoutItemCategoryRepository := ss.cPayoutItemCategoryRepository
    results, err := cPayoutItemCategoryRepository.GetCPayoutItemCategoryWithKey(ctx, cPayoutItemCategory.PayoutItemCategoryCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cPayoutItemCategoryRepository.UpdateCPayoutItemCategory(ctx, cPayoutItemCategory)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cPayoutItemCategoryService) DeleteCPayoutItemCategory(ctx *context.Context, payoutItemCategoryCd int) *errordef.LogicError {
    cPayoutItemCategoryRepository := ss.cPayoutItemCategoryRepository
    results, err := cPayoutItemCategoryRepository.GetCPayoutItemCategoryWithKey(ctx, payoutItemCategoryCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cPayoutItemCategoryRepository.DeleteCPayoutItemCategory(ctx, payoutItemCategoryCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cPayoutItemCategoryService) GetCPayoutItemCategoryWithKey(ctx *context.Context, payoutItemCategoryCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    cPayoutItemCategoryRepository := ss.cPayoutItemCategoryRepository
    cPayoutItemCategory, err := cPayoutItemCategoryRepository.GetCPayoutItemCategoryWithKey(ctx, payoutItemCategoryCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cPayoutItemCategory, nil
}

func (ss cPayoutItemCategoryService) GetCPayoutItemCategoryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    cPayoutItemCategoryRepository := ss.cPayoutItemCategoryRepository
    cPayoutItemCategory, err := cPayoutItemCategoryRepository.GetCPayoutItemCategoryWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cPayoutItemCategory, nil
}

