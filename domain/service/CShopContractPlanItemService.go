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

type CShopContractPlanItemService interface {
    CreateCShopContractPlanItem(*context.Context, *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError)
    UpdateCShopContractPlanItem(*context.Context, *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError)
    GetCShopContractPlanItemWithKey(*context.Context, int,int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
    GetCShopContractPlanItemWithLanguageCd(*context.Context, int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
    DeleteCShopContractPlanItem(*context.Context, int, int) *errordef.LogicError
}

type cShopContractPlanItemService struct {
    cShopContractPlanItemRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCShopContractPlanItemService(pr repository.Repository, nu number.NumberUtil) CShopContractPlanItemService {
    return &cShopContractPlanItemService{
        cShopContractPlanItemRepository :pr,
        numberUtil :nu,
    }
}

func (ss cShopContractPlanItemService) CreateCShopContractPlanItem(ctx *context.Context, cShopContractPlanItem *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError) {
    cShopContractPlanItem.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopContractPlanItem.CreateFunction = "CreateCShopContractPlanItem"
    cShopContractPlanItem.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopContractPlanItem.UpdateFunction = "CreateCShopContractPlanItem"
    cShopContractPlanItemRepository := ss.cShopContractPlanItemRepository
    created, err := cShopContractPlanItemRepository.CreateCShopContractPlanItem(ctx, cShopContractPlanItem)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cShopContractPlanItemService) UpdateCShopContractPlanItem(ctx *context.Context, cShopContractPlanItem *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError) {
    cShopContractPlanItem.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopContractPlanItem.UpdateFunction = "UpdateCShopContractPlanItem"
    cShopContractPlanItemRepository := ss.cShopContractPlanItemRepository
    results, err := cShopContractPlanItemRepository.GetCShopContractPlanItemWithKey(ctx, cShopContractPlanItem.ShopContractPlanItemCd, cShopContractPlanItem.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cShopContractPlanItemRepository.UpdateCShopContractPlanItem(ctx, cShopContractPlanItem)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cShopContractPlanItemService) DeleteCShopContractPlanItem(ctx *context.Context, shopContractPlanItemCd int, languageCd int) *errordef.LogicError {
    cShopContractPlanItemRepository := ss.cShopContractPlanItemRepository
    results, err := cShopContractPlanItemRepository.GetCShopContractPlanItemWithKey(ctx, shopContractPlanItemCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cShopContractPlanItemRepository.DeleteCShopContractPlanItem(ctx, shopContractPlanItemCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cShopContractPlanItemService) GetCShopContractPlanItemWithKey(ctx *context.Context, shopContractPlanItemCd int,languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    cShopContractPlanItemRepository := ss.cShopContractPlanItemRepository
    cShopContractPlanItem, err := cShopContractPlanItemRepository.GetCShopContractPlanItemWithKey(ctx, shopContractPlanItemCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopContractPlanItem, nil
}

func (ss cShopContractPlanItemService) GetCShopContractPlanItemWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    cShopContractPlanItemRepository := ss.cShopContractPlanItemRepository
    cShopContractPlanItem, err := cShopContractPlanItemRepository.GetCShopContractPlanItemWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopContractPlanItem, nil
}

