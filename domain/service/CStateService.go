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

type CStateService interface {
    CreateCState(*context.Context, *model.CState) (*model.CState, *errordef.LogicError)
    UpdateCState(*context.Context, *model.CState) (*model.CState, *errordef.LogicError)
    GetCStateWithKey(*context.Context, int,int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithStateCdIso(*context.Context, string) ([]*model.CState, *errordef.LogicError)
    GetCStateWithLanguageCd(*context.Context, int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithCountryCd(*context.Context, int,int) ([]*model.CState, *errordef.LogicError)
    GetCStateWithStateCd(*context.Context, int) ([]*model.CState, *errordef.LogicError)
    DeleteCState(*context.Context, int, int) *errordef.LogicError
}

type cStateService struct {
    cStateRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCStateService(pr repository.Repository, nu number.NumberUtil) CStateService {
    return &cStateService{
        cStateRepository :pr,
        numberUtil :nu,
    }
}

func (ss cStateService) CreateCState(ctx *context.Context, cState *model.CState) (*model.CState, *errordef.LogicError) {
    cState.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cState.CreateFunction = "CreateCState"
    cState.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cState.UpdateFunction = "CreateCState"
    cStateRepository := ss.cStateRepository
    created, err := cStateRepository.CreateCState(ctx, cState)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cStateService) UpdateCState(ctx *context.Context, cState *model.CState) (*model.CState, *errordef.LogicError) {
    cState.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cState.UpdateFunction = "UpdateCState"
    cStateRepository := ss.cStateRepository
    results, err := cStateRepository.GetCStateWithKey(ctx, cState.StateCd, cState.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cStateRepository.UpdateCState(ctx, cState)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cStateService) DeleteCState(ctx *context.Context, stateCd int, languageCd int) *errordef.LogicError {
    cStateRepository := ss.cStateRepository
    results, err := cStateRepository.GetCStateWithKey(ctx, stateCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cStateRepository.DeleteCState(ctx, stateCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cStateService) GetCStateWithKey(ctx *context.Context, stateCd int,languageCd int) ([]*model.CState, *errordef.LogicError) {
    cStateRepository := ss.cStateRepository
    cState, err := cStateRepository.GetCStateWithKey(ctx, stateCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cState, nil
}

func (ss cStateService) GetCStateWithStateCdIso(ctx *context.Context, stateCdIso string) ([]*model.CState, *errordef.LogicError) {
    cStateRepository := ss.cStateRepository
    cState, err := cStateRepository.GetCStateWithStateCdIso(ctx, stateCdIso)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cState, nil
}

func (ss cStateService) GetCStateWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CState, *errordef.LogicError) {
    cStateRepository := ss.cStateRepository
    cState, err := cStateRepository.GetCStateWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cState, nil
}

func (ss cStateService) GetCStateWithCountryCd(ctx *context.Context, languageCd int,countryCd int) ([]*model.CState, *errordef.LogicError) {
    cStateRepository := ss.cStateRepository
    cState, err := cStateRepository.GetCStateWithCountryCd(ctx, languageCd,countryCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cState, nil
}

func (ss cStateService) GetCStateWithStateCd(ctx *context.Context, stateCd int) ([]*model.CState, *errordef.LogicError) {
    cStateRepository := ss.cStateRepository
    cState, err := cStateRepository.GetCStateWithStateCd(ctx, stateCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cState, nil
}

