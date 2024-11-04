package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CAccountWithdrawalReasonApplication interface {
    GetCAccountWithdrawalReasonWithKey(*context.Context, int,int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
    GetCAccountWithdrawalReasonWithLanguageCd(*context.Context, int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
}

type cAccountWithdrawalReasonApplication struct {
    cAccountWithdrawalReasonService service.CAccountWithdrawalReasonService
    transaction repository.Transaction
}

func NewCAccountWithdrawalReasonApplication(ls service.CAccountWithdrawalReasonService, tr repository.Transaction) CAccountWithdrawalReasonApplication {
    return &cAccountWithdrawalReasonApplication{
        cAccountWithdrawalReasonService :ls,
        transaction :tr,
    }
}

func (sa cAccountWithdrawalReasonApplication) GetCAccountWithdrawalReasonWithKey(ctx *context.Context, accountWithdrawalReasonCd int,languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAccountWithdrawalReasonService.GetCAccountWithdrawalReasonWithKey(ctx, accountWithdrawalReasonCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cAccountWithdrawalReasonApplication) GetCAccountWithdrawalReasonWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cAccountWithdrawalReasonService.GetCAccountWithdrawalReasonWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
