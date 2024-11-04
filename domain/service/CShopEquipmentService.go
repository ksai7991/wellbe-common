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

type CShopEquipmentService interface {
    CreateCShopEquipment(*context.Context, *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError)
    UpdateCShopEquipment(*context.Context, *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError)
    GetCShopEquipmentWithKey(*context.Context, int,int) ([]*model.CShopEquipment, *errordef.LogicError)
    GetCShopEquipmentWithLanguageCd(*context.Context, int) ([]*model.CShopEquipment, *errordef.LogicError)
    DeleteCShopEquipment(*context.Context, int, int) *errordef.LogicError
}

type cShopEquipmentService struct {
    cShopEquipmentRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCShopEquipmentService(pr repository.Repository, nu number.NumberUtil) CShopEquipmentService {
    return &cShopEquipmentService{
        cShopEquipmentRepository :pr,
        numberUtil :nu,
    }
}

func (ss cShopEquipmentService) CreateCShopEquipment(ctx *context.Context, cShopEquipment *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError) {
    cShopEquipment.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopEquipment.CreateFunction = "CreateCShopEquipment"
    cShopEquipment.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopEquipment.UpdateFunction = "CreateCShopEquipment"
    cShopEquipmentRepository := ss.cShopEquipmentRepository
    created, err := cShopEquipmentRepository.CreateCShopEquipment(ctx, cShopEquipment)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cShopEquipmentService) UpdateCShopEquipment(ctx *context.Context, cShopEquipment *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError) {
    cShopEquipment.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopEquipment.UpdateFunction = "UpdateCShopEquipment"
    cShopEquipmentRepository := ss.cShopEquipmentRepository
    results, err := cShopEquipmentRepository.GetCShopEquipmentWithKey(ctx, cShopEquipment.ShopEquipmentCd, cShopEquipment.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cShopEquipmentRepository.UpdateCShopEquipment(ctx, cShopEquipment)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cShopEquipmentService) DeleteCShopEquipment(ctx *context.Context, shopEquipmentCd int, languageCd int) *errordef.LogicError {
    cShopEquipmentRepository := ss.cShopEquipmentRepository
    results, err := cShopEquipmentRepository.GetCShopEquipmentWithKey(ctx, shopEquipmentCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cShopEquipmentRepository.DeleteCShopEquipment(ctx, shopEquipmentCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cShopEquipmentService) GetCShopEquipmentWithKey(ctx *context.Context, shopEquipmentCd int,languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
    cShopEquipmentRepository := ss.cShopEquipmentRepository
    cShopEquipment, err := cShopEquipmentRepository.GetCShopEquipmentWithKey(ctx, shopEquipmentCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopEquipment, nil
}

func (ss cShopEquipmentService) GetCShopEquipmentWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
    cShopEquipmentRepository := ss.cShopEquipmentRepository
    cShopEquipment, err := cShopEquipmentRepository.GetCShopEquipmentWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopEquipment, nil
}

