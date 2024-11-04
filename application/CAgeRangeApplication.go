package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CAgeRangeApplication interface {
    GetCAgeRangeWithKey(*context.Context, int,int) ([]*model.CAgeRange, *errordef.LogicError)
    GetCAgeRangeWithLanguageCd(*context.Context, int) ([]*model.CAgeRange, *errordef.LogicError)
}

type cAgeRangeApplication struct {
    cAgeRangeService service.CAgeRangeService
    transaction repository.Transaction
}

func NewCAgeRangeApplication(ls service.CAgeRangeService, tr repository.Transaction) CAgeRangeApplication {
    return &cAgeRangeApplication{
        cAgeRangeService :ls,
        transaction :tr,
    }
}

func (sa cAgeRangeApplication) GetCAgeRangeWithKey(ctx *context.Context, ageRangeCd int,languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAgeRangeService.GetCAgeRangeWithKey(ctx, ageRangeCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cAgeRangeApplication) GetCAgeRangeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAgeRangeService.GetCAgeRangeWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
