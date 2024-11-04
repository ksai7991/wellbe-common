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

type CCountryService interface {
    CreateCCountry(*context.Context, *model.CCountry) (*model.CCountry, *errordef.LogicError)
    UpdateCCountry(*context.Context, *model.CCountry) (*model.CCountry, *errordef.LogicError)
    GetCCountryWithKey(*context.Context, int,int) ([]*model.CCountry, *errordef.LogicError)
    GetCCountryWithCountryCdIso(*context.Context, string) ([]*model.CCountry, *errordef.LogicError)
    GetCCountryWithLanguageCd(*context.Context, int) ([]*model.CCountry, *errordef.LogicError)
    DeleteCCountry(*context.Context, int, int) *errordef.LogicError
}

type cCountryService struct {
    cCountryRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCCountryService(pr repository.Repository, nu number.NumberUtil) CCountryService {
    return &cCountryService{
        cCountryRepository :pr,
        numberUtil :nu,
    }
}

func (ss cCountryService) CreateCCountry(ctx *context.Context, cCountry *model.CCountry) (*model.CCountry, *errordef.LogicError) {
    cCountry.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCountry.CreateFunction = "CreateCCountry"
    cCountry.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCountry.UpdateFunction = "CreateCCountry"
    cCountryRepository := ss.cCountryRepository
    created, err := cCountryRepository.CreateCCountry(ctx, cCountry)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cCountryService) UpdateCCountry(ctx *context.Context, cCountry *model.CCountry) (*model.CCountry, *errordef.LogicError) {
    cCountry.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cCountry.UpdateFunction = "UpdateCCountry"
    cCountryRepository := ss.cCountryRepository
    results, err := cCountryRepository.GetCCountryWithKey(ctx, cCountry.CountryCd, cCountry.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cCountryRepository.UpdateCCountry(ctx, cCountry)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cCountryService) DeleteCCountry(ctx *context.Context, countryCd int, languageCd int) *errordef.LogicError {
    cCountryRepository := ss.cCountryRepository
    results, err := cCountryRepository.GetCCountryWithKey(ctx, countryCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cCountryRepository.DeleteCCountry(ctx, countryCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cCountryService) GetCCountryWithKey(ctx *context.Context, countryCd int,languageCd int) ([]*model.CCountry, *errordef.LogicError) {
    cCountryRepository := ss.cCountryRepository
    cCountry, err := cCountryRepository.GetCCountryWithKey(ctx, countryCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCountry, nil
}

func (ss cCountryService) GetCCountryWithCountryCdIso(ctx *context.Context, countryCdIso string) ([]*model.CCountry, *errordef.LogicError) {
    cCountryRepository := ss.cCountryRepository
    cCountry, err := cCountryRepository.GetCCountryWithCountryCdIso(ctx, countryCdIso)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCountry, nil
}

func (ss cCountryService) GetCCountryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCountry, *errordef.LogicError) {
    cCountryRepository := ss.cCountryRepository
    cCountry, err := cCountryRepository.GetCCountryWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cCountry, nil
}

