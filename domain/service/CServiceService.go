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

type CServiceService interface {
    CreateCService(*context.Context, *model.CService) (*model.CService, *errordef.LogicError)
    UpdateCService(*context.Context, *model.CService) (*model.CService, *errordef.LogicError)
    GetCServiceWithKey(*context.Context, int,int) ([]*model.CService, *errordef.LogicError)
    GetCServiceWithLanguageCd(*context.Context, int) ([]*model.CService, *errordef.LogicError)
    DeleteCService(*context.Context, int, int) *errordef.LogicError
}

type cServiceService struct {
    cServiceRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCServiceService(pr repository.Repository, nu number.NumberUtil) CServiceService {
    return &cServiceService{
        cServiceRepository :pr,
        numberUtil :nu,
    }
}

func (ss cServiceService) CreateCService(ctx *context.Context, cService *model.CService) (*model.CService, *errordef.LogicError) {
    cService.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cService.CreateFunction = "CreateCService"
    cService.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cService.UpdateFunction = "CreateCService"
    cServiceRepository := ss.cServiceRepository
    created, err := cServiceRepository.CreateCService(ctx, cService)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cServiceService) UpdateCService(ctx *context.Context, cService *model.CService) (*model.CService, *errordef.LogicError) {
    cService.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cService.UpdateFunction = "UpdateCService"
    cServiceRepository := ss.cServiceRepository
    results, err := cServiceRepository.GetCServiceWithKey(ctx, cService.ServiceCd, cService.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cServiceRepository.UpdateCService(ctx, cService)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cServiceService) DeleteCService(ctx *context.Context, serviceCd int, languageCd int) *errordef.LogicError {
    cServiceRepository := ss.cServiceRepository
    results, err := cServiceRepository.GetCServiceWithKey(ctx, serviceCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cServiceRepository.DeleteCService(ctx, serviceCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cServiceService) GetCServiceWithKey(ctx *context.Context, serviceCd int,languageCd int) ([]*model.CService, *errordef.LogicError) {
    cServiceRepository := ss.cServiceRepository
    cService, err := cServiceRepository.GetCServiceWithKey(ctx, serviceCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cService, nil
}

func (ss cServiceService) GetCServiceWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CService, *errordef.LogicError) {
    cServiceRepository := ss.cServiceRepository
    cService, err := cServiceRepository.GetCServiceWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cService, nil
}

