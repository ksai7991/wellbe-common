package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CLanguageApplication interface {
    GetCLanguageWithKey(*context.Context, int) ([]*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithLanguageCharCd(*context.Context, string) ([]*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithFilterCol(*context.Context, string,string) ([]*model.CLanguage, *errordef.LogicError)
}

type cLanguageApplication struct {
    cLanguageService service.CLanguageService
    transaction repository.Transaction
}

func NewCLanguageApplication(ls service.CLanguageService, tr repository.Transaction) CLanguageApplication {
    return &cLanguageApplication{
        cLanguageService :ls,
        transaction :tr,
    }
}

func (sa cLanguageApplication) GetCLanguageWithKey(ctx *context.Context, languageCd int) ([]*model.CLanguage, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cLanguageService.GetCLanguageWithKey(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cLanguageApplication) GetCLanguageWithLanguageCharCd(ctx *context.Context, languageCharCd string) ([]*model.CLanguage, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cLanguageService.GetCLanguageWithLanguageCharCd(ctx, languageCharCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cLanguageApplication) GetCLanguageWithFilterCol(ctx *context.Context, languageCharCd string,languageName string) ([]*model.CLanguage, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cLanguageService.GetCLanguageWithFilterCol(ctx, languageCharCd,languageName)
    sa.transaction.Commit(ctx)
    return result, err
}
