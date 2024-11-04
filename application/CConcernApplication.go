package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CConcernApplication interface {
    GetCConcernWithKey(*context.Context, int,int) ([]*model.CConcern, *errordef.LogicError)
    GetCConcernWithLanguageCd(*context.Context, int) ([]*model.CConcern, *errordef.LogicError)
}

type cConcernApplication struct {
    cConcernService service.CConcernService
    transaction repository.Transaction
}

func NewCConcernApplication(ls service.CConcernService, tr repository.Transaction) CConcernApplication {
    return &cConcernApplication{
        cConcernService :ls,
        transaction :tr,
    }
}

func (sa cConcernApplication) GetCConcernWithKey(ctx *context.Context, concernCd int,languageCd int) ([]*model.CConcern, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cConcernService.GetCConcernWithKey(ctx, concernCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cConcernApplication) GetCConcernWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CConcern, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cConcernService.GetCConcernWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
