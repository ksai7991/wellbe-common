package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CMailTemplateApplication interface {
    GetCMailTemplateWithKey(*context.Context, int,int) ([]*model.CMailTemplate, *errordef.LogicError)
    GetCMailTemplateWithLanguageCd(*context.Context, int) ([]*model.CMailTemplate, *errordef.LogicError)
}

type cMailTemplateApplication struct {
    cMailTemplateService service.CMailTemplateService
    transaction repository.Transaction
}

func NewCMailTemplateApplication(ls service.CMailTemplateService, tr repository.Transaction) CMailTemplateApplication {
    return &cMailTemplateApplication{
        cMailTemplateService :ls,
        transaction :tr,
    }
}

func (sa cMailTemplateApplication) GetCMailTemplateWithKey(ctx *context.Context, mailTemplateCd int,languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cMailTemplateService.GetCMailTemplateWithKey(ctx, mailTemplateCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cMailTemplateApplication) GetCMailTemplateWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cMailTemplateService.GetCMailTemplateWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
