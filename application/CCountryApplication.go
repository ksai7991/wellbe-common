package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CCountryApplication interface {
    GetCCountryWithKey(*context.Context, int,int) ([]*model.CCountry, *errordef.LogicError)
    GetCCountryWithCountryCdIso(*context.Context, string) ([]*model.CCountry, *errordef.LogicError)
    GetCCountryWithLanguageCd(*context.Context, int) ([]*model.CCountry, *errordef.LogicError)
}

type cCountryApplication struct {
    cCountryService service.CCountryService
    transaction repository.Transaction
}

func NewCCountryApplication(ls service.CCountryService, tr repository.Transaction) CCountryApplication {
    return &cCountryApplication{
        cCountryService :ls,
        transaction :tr,
    }
}

func (sa cCountryApplication) GetCCountryWithKey(ctx *context.Context, countryCd int,languageCd int) ([]*model.CCountry, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCountryService.GetCCountryWithKey(ctx, countryCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCountryApplication) GetCCountryWithCountryCdIso(ctx *context.Context, countryCdIso string) ([]*model.CCountry, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCountryService.GetCCountryWithCountryCdIso(ctx, countryCdIso)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cCountryApplication) GetCCountryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCountry, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cCountryService.GetCCountryWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
