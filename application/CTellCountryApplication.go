package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CTellCountryApplication interface {
    GetCTellCountryWithKey(*context.Context, int,int) ([]*model.CTellCountry, *errordef.LogicError)
    GetCTellCountryWithLanguageCd(*context.Context, int) ([]*model.CTellCountry, *errordef.LogicError)
}

type cTellCountryApplication struct {
    cTellCountryService service.CTellCountryService
    transaction repository.Transaction
}

func NewCTellCountryApplication(ls service.CTellCountryService, tr repository.Transaction) CTellCountryApplication {
    return &cTellCountryApplication{
        cTellCountryService :ls,
        transaction :tr,
    }
}

func (sa cTellCountryApplication) GetCTellCountryWithKey(ctx *context.Context, languageCd int,tellCountryCd int) ([]*model.CTellCountry, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cTellCountryService.GetCTellCountryWithKey(ctx, languageCd,tellCountryCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cTellCountryApplication) GetCTellCountryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTellCountry, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cTellCountryService.GetCTellCountryWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
