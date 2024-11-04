package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CPayoutStatusApplication interface {
    GetCPayoutStatusWithKey(*context.Context, int) ([]*model.CPayoutStatus, *errordef.LogicError)
    GetCPayoutStatusWithLanguageCd(*context.Context, int) ([]*model.CPayoutStatus, *errordef.LogicError)
}

type cPayoutStatusApplication struct {
    cPayoutStatusService service.CPayoutStatusService
    transaction repository.Transaction
}

func NewCPayoutStatusApplication(ls service.CPayoutStatusService, tr repository.Transaction) CPayoutStatusApplication {
    return &cPayoutStatusApplication{
        cPayoutStatusService :ls,
        transaction :tr,
    }
}

func (sa cPayoutStatusApplication) GetCPayoutStatusWithKey(ctx *context.Context, payoutStatusCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cPayoutStatusService.GetCPayoutStatusWithKey(ctx, payoutStatusCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cPayoutStatusApplication) GetCPayoutStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cPayoutStatusService.GetCPayoutStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
