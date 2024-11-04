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

type CShopMaintenanceLabelService interface {
    CreateCShopMaintenanceLabel(*context.Context, *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError)
    UpdateCShopMaintenanceLabel(*context.Context, *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError)
    GetCShopMaintenanceLabelWithKey(*context.Context, int,int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
    GetCShopMaintenanceLabelWithLanguageCd(*context.Context, int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
    DeleteCShopMaintenanceLabel(*context.Context, int, int) *errordef.LogicError
}

type cShopMaintenanceLabelService struct {
    cShopMaintenanceLabelRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCShopMaintenanceLabelService(pr repository.Repository, nu number.NumberUtil) CShopMaintenanceLabelService {
    return &cShopMaintenanceLabelService{
        cShopMaintenanceLabelRepository :pr,
        numberUtil :nu,
    }
}

func (ss cShopMaintenanceLabelService) CreateCShopMaintenanceLabel(ctx *context.Context, cShopMaintenanceLabel *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError) {
    cShopMaintenanceLabel.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopMaintenanceLabel.CreateFunction = "CreateCShopMaintenanceLabel"
    cShopMaintenanceLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopMaintenanceLabel.UpdateFunction = "CreateCShopMaintenanceLabel"
    cShopMaintenanceLabelRepository := ss.cShopMaintenanceLabelRepository
    created, err := cShopMaintenanceLabelRepository.CreateCShopMaintenanceLabel(ctx, cShopMaintenanceLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cShopMaintenanceLabelService) UpdateCShopMaintenanceLabel(ctx *context.Context, cShopMaintenanceLabel *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError) {
    cShopMaintenanceLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopMaintenanceLabel.UpdateFunction = "UpdateCShopMaintenanceLabel"
    cShopMaintenanceLabelRepository := ss.cShopMaintenanceLabelRepository
    results, err := cShopMaintenanceLabelRepository.GetCShopMaintenanceLabelWithKey(ctx, cShopMaintenanceLabel.ShopMaintenanceLabelCd, cShopMaintenanceLabel.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cShopMaintenanceLabelRepository.UpdateCShopMaintenanceLabel(ctx, cShopMaintenanceLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cShopMaintenanceLabelService) DeleteCShopMaintenanceLabel(ctx *context.Context, shopMaintenanceLabelCd int, languageCd int) *errordef.LogicError {
    cShopMaintenanceLabelRepository := ss.cShopMaintenanceLabelRepository
    results, err := cShopMaintenanceLabelRepository.GetCShopMaintenanceLabelWithKey(ctx, shopMaintenanceLabelCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cShopMaintenanceLabelRepository.DeleteCShopMaintenanceLabel(ctx, shopMaintenanceLabelCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cShopMaintenanceLabelService) GetCShopMaintenanceLabelWithKey(ctx *context.Context, shopMaintenanceLabelCd int,languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    cShopMaintenanceLabelRepository := ss.cShopMaintenanceLabelRepository
    cShopMaintenanceLabel, err := cShopMaintenanceLabelRepository.GetCShopMaintenanceLabelWithKey(ctx, shopMaintenanceLabelCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopMaintenanceLabel, nil
}

func (ss cShopMaintenanceLabelService) GetCShopMaintenanceLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    cShopMaintenanceLabelRepository := ss.cShopMaintenanceLabelRepository
    cShopMaintenanceLabel, err := cShopMaintenanceLabelRepository.GetCShopMaintenanceLabelWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopMaintenanceLabel, nil
}

