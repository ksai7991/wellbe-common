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

type CMailTemplateService interface {
    CreateCMailTemplate(*context.Context, *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError)
    UpdateCMailTemplate(*context.Context, *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError)
    GetCMailTemplateWithKey(*context.Context, int,int) ([]*model.CMailTemplate, *errordef.LogicError)
    GetCMailTemplateWithLanguageCd(*context.Context, int) ([]*model.CMailTemplate, *errordef.LogicError)
    DeleteCMailTemplate(*context.Context, int, int) *errordef.LogicError
}

type cMailTemplateService struct {
    cMailTemplateRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCMailTemplateService(pr repository.Repository, nu number.NumberUtil) CMailTemplateService {
    return &cMailTemplateService{
        cMailTemplateRepository :pr,
        numberUtil :nu,
    }
}

func (ss cMailTemplateService) CreateCMailTemplate(ctx *context.Context, cMailTemplate *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError) {
    cMailTemplate.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cMailTemplate.CreateFunction = "CreateCMailTemplate"
    cMailTemplate.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cMailTemplate.UpdateFunction = "CreateCMailTemplate"
    cMailTemplateRepository := ss.cMailTemplateRepository
    created, err := cMailTemplateRepository.CreateCMailTemplate(ctx, cMailTemplate)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cMailTemplateService) UpdateCMailTemplate(ctx *context.Context, cMailTemplate *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError) {
    cMailTemplate.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cMailTemplate.UpdateFunction = "UpdateCMailTemplate"
    cMailTemplateRepository := ss.cMailTemplateRepository
    results, err := cMailTemplateRepository.GetCMailTemplateWithKey(ctx, cMailTemplate.MailTemplateCd, cMailTemplate.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cMailTemplateRepository.UpdateCMailTemplate(ctx, cMailTemplate)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cMailTemplateService) DeleteCMailTemplate(ctx *context.Context, mailTemplateCd int, languageCd int) *errordef.LogicError {
    cMailTemplateRepository := ss.cMailTemplateRepository
    results, err := cMailTemplateRepository.GetCMailTemplateWithKey(ctx, mailTemplateCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cMailTemplateRepository.DeleteCMailTemplate(ctx, mailTemplateCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cMailTemplateService) GetCMailTemplateWithKey(ctx *context.Context, mailTemplateCd int,languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
    cMailTemplateRepository := ss.cMailTemplateRepository
    cMailTemplate, err := cMailTemplateRepository.GetCMailTemplateWithKey(ctx, mailTemplateCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cMailTemplate, nil
}

func (ss cMailTemplateService) GetCMailTemplateWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
    cMailTemplateRepository := ss.cMailTemplateRepository
    cMailTemplate, err := cMailTemplateRepository.GetCMailTemplateWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cMailTemplate, nil
}

