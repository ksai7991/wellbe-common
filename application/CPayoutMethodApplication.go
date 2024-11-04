package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CPayoutMethodApplication interface {
    GetCPayoutMethodWithKey(*context.Context, int) ([]*model.CPayoutMethod, *errordef.LogicError)
    GetCPayoutMethodWithLanguageCd(*context.Context, int) ([]*model.CPayoutMethod, *errordef.LogicError)
}

type cPayoutMethodApplication struct {
    cPayoutMethodService service.CPayoutMethodService
    transaction repository.Transaction
}

func NewCPayoutMethodApplication(ls service.CPayoutMethodService, tr repository.Transaction) CPayoutMethodApplication {
    return &cPayoutMethodApplication{
        cPayoutMethodService :ls,
        transaction :tr,
    }
}

func (sa cPayoutMethodApplication) GetCPayoutMethodWithKey(ctx *context.Context, payoutMethodCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cPayoutMethodService.GetCPayoutMethodWithKey(ctx, payoutMethodCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cPayoutMethodApplication) GetCPayoutMethodWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cPayoutMethodService.GetCPayoutMethodWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
