package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CGenderApplication interface {
    GetCGenderWithKey(*context.Context, int,int) ([]*model.CGender, *errordef.LogicError)
    GetCGenderWithLanguageCd(*context.Context, int) ([]*model.CGender, *errordef.LogicError)
}

type cGenderApplication struct {
    cGenderService service.CGenderService
    transaction repository.Transaction
}

func NewCGenderApplication(ls service.CGenderService, tr repository.Transaction) CGenderApplication {
    return &cGenderApplication{
        cGenderService :ls,
        transaction :tr,
    }
}

func (sa cGenderApplication) GetCGenderWithKey(ctx *context.Context, genderCd int,languageCd int) ([]*model.CGender, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cGenderService.GetCGenderWithKey(ctx, genderCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cGenderApplication) GetCGenderWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CGender, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cGenderService.GetCGenderWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
