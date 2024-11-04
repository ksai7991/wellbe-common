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

type COrderTypeService interface {
    CreateCOrderType(*context.Context, *model.COrderType) (*model.COrderType, *errordef.LogicError)
    UpdateCOrderType(*context.Context, *model.COrderType) (*model.COrderType, *errordef.LogicError)
    GetCOrderTypeWithKey(*context.Context, int) ([]*model.COrderType, *errordef.LogicError)
    GetCOrderTypeWithLanguageCd(*context.Context, int) ([]*model.COrderType, *errordef.LogicError)
    DeleteCOrderType(*context.Context, int) *errordef.LogicError
}

type cOrderTypeService struct {
    cOrderTypeRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCOrderTypeService(pr repository.Repository, nu number.NumberUtil) COrderTypeService {
    return &cOrderTypeService{
        cOrderTypeRepository :pr,
        numberUtil :nu,
    }
}

func (ss cOrderTypeService) CreateCOrderType(ctx *context.Context, cOrderType *model.COrderType) (*model.COrderType, *errordef.LogicError) {
    cOrderType.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cOrderType.CreateFunction = "CreateCOrderType"
    cOrderType.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cOrderType.UpdateFunction = "CreateCOrderType"
    cOrderTypeRepository := ss.cOrderTypeRepository
    created, err := cOrderTypeRepository.CreateCOrderType(ctx, cOrderType)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cOrderTypeService) UpdateCOrderType(ctx *context.Context, cOrderType *model.COrderType) (*model.COrderType, *errordef.LogicError) {
    cOrderType.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cOrderType.UpdateFunction = "UpdateCOrderType"
    cOrderTypeRepository := ss.cOrderTypeRepository
    results, err := cOrderTypeRepository.GetCOrderTypeWithKey(ctx, cOrderType.OrderTypeCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cOrderTypeRepository.UpdateCOrderType(ctx, cOrderType)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cOrderTypeService) DeleteCOrderType(ctx *context.Context, orderTypeCd int) *errordef.LogicError {
    cOrderTypeRepository := ss.cOrderTypeRepository
    results, err := cOrderTypeRepository.GetCOrderTypeWithKey(ctx, orderTypeCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cOrderTypeRepository.DeleteCOrderType(ctx, orderTypeCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cOrderTypeService) GetCOrderTypeWithKey(ctx *context.Context, orderTypeCd int) ([]*model.COrderType, *errordef.LogicError) {
    cOrderTypeRepository := ss.cOrderTypeRepository
    cOrderType, err := cOrderTypeRepository.GetCOrderTypeWithKey(ctx, orderTypeCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cOrderType, nil
}

func (ss cOrderTypeService) GetCOrderTypeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.COrderType, *errordef.LogicError) {
    cOrderTypeRepository := ss.cOrderTypeRepository
    cOrderType, err := cOrderTypeRepository.GetCOrderTypeWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cOrderType, nil
}

