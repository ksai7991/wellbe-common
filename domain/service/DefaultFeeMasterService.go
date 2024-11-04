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

type DefaultFeeMasterService interface {
    CreateDefaultFeeMaster(*context.Context, *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    UpdateDefaultFeeMaster(*context.Context, *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    GetDefaultFeeMasterWithKey(*context.Context, string) ([]*model.DefaultFeeMaster, *errordef.LogicError)
    DeleteDefaultFeeMaster(*context.Context, string) *errordef.LogicError
}

type defaultFeeMasterService struct {
    defaultFeeMasterRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewDefaultFeeMasterService(pr repository.Repository, nu number.NumberUtil) DefaultFeeMasterService {
    return &defaultFeeMasterService{
        defaultFeeMasterRepository :pr,
        numberUtil :nu,
    }
}

func (ss defaultFeeMasterService) CreateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
    defaultFeeMaster.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    defaultFeeMaster.CreateFunction = "CreateDefaultFeeMaster"
    defaultFeeMaster.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    defaultFeeMaster.UpdateFunction = "CreateDefaultFeeMaster"
    defaultFeeMasterRepository := ss.defaultFeeMasterRepository
    created, err := defaultFeeMasterRepository.CreateDefaultFeeMaster(ctx, defaultFeeMaster)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss defaultFeeMasterService) UpdateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
    defaultFeeMaster.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    defaultFeeMaster.UpdateFunction = "UpdateDefaultFeeMaster"
    defaultFeeMasterRepository := ss.defaultFeeMasterRepository
    results, err := defaultFeeMasterRepository.GetDefaultFeeMasterWithKey(ctx, defaultFeeMaster.Id)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := defaultFeeMasterRepository.UpdateDefaultFeeMaster(ctx, defaultFeeMaster)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss defaultFeeMasterService) DeleteDefaultFeeMaster(ctx *context.Context, id string) *errordef.LogicError {
    defaultFeeMasterRepository := ss.defaultFeeMasterRepository
    results, err := defaultFeeMasterRepository.GetDefaultFeeMasterWithKey(ctx, id)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = defaultFeeMasterRepository.DeleteDefaultFeeMaster(ctx, id)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss defaultFeeMasterService) GetDefaultFeeMasterWithKey(ctx *context.Context, id string) ([]*model.DefaultFeeMaster, *errordef.LogicError) {
    defaultFeeMasterRepository := ss.defaultFeeMasterRepository
    defaultFeeMaster, err := defaultFeeMasterRepository.GetDefaultFeeMasterWithKey(ctx, id)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return defaultFeeMaster, nil
}

