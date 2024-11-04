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

type CTellCountryService interface {
    CreateCTellCountry(*context.Context, *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError)
    UpdateCTellCountry(*context.Context, *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError)
    GetCTellCountryWithKey(*context.Context, int,int) ([]*model.CTellCountry, *errordef.LogicError)
    GetCTellCountryWithLanguageCd(*context.Context, int) ([]*model.CTellCountry, *errordef.LogicError)
    DeleteCTellCountry(*context.Context, int, int) *errordef.LogicError
}

type cTellCountryService struct {
    cTellCountryRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCTellCountryService(pr repository.Repository, nu number.NumberUtil) CTellCountryService {
    return &cTellCountryService{
        cTellCountryRepository :pr,
        numberUtil :nu,
    }
}

func (ss cTellCountryService) CreateCTellCountry(ctx *context.Context, cTellCountry *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError) {
    cTellCountry.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTellCountry.CreateFunction = "CreateCTellCountry"
    cTellCountry.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTellCountry.UpdateFunction = "CreateCTellCountry"
    cTellCountryRepository := ss.cTellCountryRepository
    created, err := cTellCountryRepository.CreateCTellCountry(ctx, cTellCountry)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cTellCountryService) UpdateCTellCountry(ctx *context.Context, cTellCountry *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError) {
    cTellCountry.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTellCountry.UpdateFunction = "UpdateCTellCountry"
    cTellCountryRepository := ss.cTellCountryRepository
    results, err := cTellCountryRepository.GetCTellCountryWithKey(ctx, cTellCountry.LanguageCd, cTellCountry.TellCountryCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cTellCountryRepository.UpdateCTellCountry(ctx, cTellCountry)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cTellCountryService) DeleteCTellCountry(ctx *context.Context, languageCd int, tellCountryCd int) *errordef.LogicError {
    cTellCountryRepository := ss.cTellCountryRepository
    results, err := cTellCountryRepository.GetCTellCountryWithKey(ctx, languageCd, tellCountryCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cTellCountryRepository.DeleteCTellCountry(ctx, languageCd, tellCountryCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cTellCountryService) GetCTellCountryWithKey(ctx *context.Context, languageCd int,tellCountryCd int) ([]*model.CTellCountry, *errordef.LogicError) {
    cTellCountryRepository := ss.cTellCountryRepository
    cTellCountry, err := cTellCountryRepository.GetCTellCountryWithKey(ctx, languageCd,tellCountryCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cTellCountry, nil
}

func (ss cTellCountryService) GetCTellCountryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTellCountry, *errordef.LogicError) {
    cTellCountryRepository := ss.cTellCountryRepository
    cTellCountry, err := cTellCountryRepository.GetCTellCountryWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cTellCountry, nil
}

