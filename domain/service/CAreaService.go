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

type CAreaService interface {
    CreateCArea(*context.Context, *model.CArea) (*model.CArea, *errordef.LogicError)
    UpdateCArea(*context.Context, *model.CArea) (*model.CArea, *errordef.LogicError)
    GetCAreaWithKey(*context.Context, int,int) ([]*model.CArea, *errordef.LogicError)
    GetCAreaWithLanguageCd(*context.Context, int) ([]*model.CArea, *errordef.LogicError)
    GetCAreaWithStateCd(*context.Context, int) ([]*model.CArea, *errordef.LogicError)
    DeleteCArea(*context.Context, int, int) *errordef.LogicError
}

type cAreaService struct {
    cAreaRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCAreaService(pr repository.Repository, nu number.NumberUtil) CAreaService {
    return &cAreaService{
        cAreaRepository :pr,
        numberUtil :nu,
    }
}

func (ss cAreaService) CreateCArea(ctx *context.Context, cArea *model.CArea) (*model.CArea, *errordef.LogicError) {
    cArea.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cArea.CreateFunction = "CreateCArea"
    cArea.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cArea.UpdateFunction = "CreateCArea"
    cAreaRepository := ss.cAreaRepository
    created, err := cAreaRepository.CreateCArea(ctx, cArea)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cAreaService) UpdateCArea(ctx *context.Context, cArea *model.CArea) (*model.CArea, *errordef.LogicError) {
    cArea.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cArea.UpdateFunction = "UpdateCArea"
    cAreaRepository := ss.cAreaRepository
    results, err := cAreaRepository.GetCAreaWithKey(ctx, cArea.LanguageCd, cArea.AreaCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cAreaRepository.UpdateCArea(ctx, cArea)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cAreaService) DeleteCArea(ctx *context.Context, languageCd int, areaCd int) *errordef.LogicError {
    cAreaRepository := ss.cAreaRepository
    results, err := cAreaRepository.GetCAreaWithKey(ctx, languageCd, areaCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cAreaRepository.DeleteCArea(ctx, languageCd, areaCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cAreaService) GetCAreaWithKey(ctx *context.Context, languageCd int,areaCd int) ([]*model.CArea, *errordef.LogicError) {
    cAreaRepository := ss.cAreaRepository
    cArea, err := cAreaRepository.GetCAreaWithKey(ctx, languageCd,areaCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cArea, nil
}

func (ss cAreaService) GetCAreaWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CArea, *errordef.LogicError) {
    cAreaRepository := ss.cAreaRepository
    cArea, err := cAreaRepository.GetCAreaWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cArea, nil
}

func (ss cAreaService) GetCAreaWithStateCd(ctx *context.Context, stateCd int) ([]*model.CArea, *errordef.LogicError) {
    cAreaRepository := ss.cAreaRepository
    cArea, err := cAreaRepository.GetCAreaWithStateCd(ctx, stateCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cArea, nil
}

