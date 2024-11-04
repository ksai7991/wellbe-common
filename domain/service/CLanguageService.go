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

type CLanguageService interface {
    CreateCLanguage(*context.Context, *model.CLanguage) (*model.CLanguage, *errordef.LogicError)
    UpdateCLanguage(*context.Context, *model.CLanguage) (*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithKey(*context.Context, int) ([]*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithLanguageCharCd(*context.Context, string) ([]*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithFilterCol(*context.Context, string,string) ([]*model.CLanguage, *errordef.LogicError)
    DeleteCLanguage(*context.Context, int) *errordef.LogicError
}

type cLanguageService struct {
    cLanguageRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCLanguageService(pr repository.Repository, nu number.NumberUtil) CLanguageService {
    return &cLanguageService{
        cLanguageRepository :pr,
        numberUtil :nu,
    }
}

func (ss cLanguageService) CreateCLanguage(ctx *context.Context, cLanguage *model.CLanguage) (*model.CLanguage, *errordef.LogicError) {
    cLanguage.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cLanguage.CreateFunction = "CreateCLanguage"
    cLanguage.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cLanguage.UpdateFunction = "CreateCLanguage"
    cLanguageRepository := ss.cLanguageRepository
    created, err := cLanguageRepository.CreateCLanguage(ctx, cLanguage)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cLanguageService) UpdateCLanguage(ctx *context.Context, cLanguage *model.CLanguage) (*model.CLanguage, *errordef.LogicError) {
    cLanguage.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cLanguage.UpdateFunction = "UpdateCLanguage"
    cLanguageRepository := ss.cLanguageRepository
    results, err := cLanguageRepository.GetCLanguageWithKey(ctx, cLanguage.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cLanguageRepository.UpdateCLanguage(ctx, cLanguage)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cLanguageService) DeleteCLanguage(ctx *context.Context, languageCd int) *errordef.LogicError {
    cLanguageRepository := ss.cLanguageRepository
    results, err := cLanguageRepository.GetCLanguageWithKey(ctx, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cLanguageRepository.DeleteCLanguage(ctx, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cLanguageService) GetCLanguageWithKey(ctx *context.Context, languageCd int) ([]*model.CLanguage, *errordef.LogicError) {
    cLanguageRepository := ss.cLanguageRepository
    cLanguage, err := cLanguageRepository.GetCLanguageWithKey(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cLanguage, nil
}

func (ss cLanguageService) GetCLanguageWithLanguageCharCd(ctx *context.Context, languageCharCd string) ([]*model.CLanguage, *errordef.LogicError) {
    cLanguageRepository := ss.cLanguageRepository
    cLanguage, err := cLanguageRepository.GetCLanguageWithLanguageCharCd(ctx, languageCharCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cLanguage, nil
}

func (ss cLanguageService) GetCLanguageWithFilterCol(ctx *context.Context, languageCharCd string,languageName string) ([]*model.CLanguage, *errordef.LogicError) {
    cLanguageRepository := ss.cLanguageRepository
    cLanguage, err := cLanguageRepository.GetCLanguageWithFilterCol(ctx, languageCharCd,languageName)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cLanguage, nil
}

