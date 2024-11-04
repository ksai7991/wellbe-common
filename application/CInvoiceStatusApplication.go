package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CInvoiceStatusApplication interface {
    GetCInvoiceStatusWithKey(*context.Context, int,int) ([]*model.CInvoiceStatus, *errordef.LogicError)
    GetCInvoiceStatusWithLanguageCd(*context.Context, int) ([]*model.CInvoiceStatus, *errordef.LogicError)
}

type cInvoiceStatusApplication struct {
    cInvoiceStatusService service.CInvoiceStatusService
    transaction repository.Transaction
}

func NewCInvoiceStatusApplication(ls service.CInvoiceStatusService, tr repository.Transaction) CInvoiceStatusApplication {
    return &cInvoiceStatusApplication{
        cInvoiceStatusService :ls,
        transaction :tr,
    }
}

func (sa cInvoiceStatusApplication) GetCInvoiceStatusWithKey(ctx *context.Context, invoiceStatusCd int,languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cInvoiceStatusService.GetCInvoiceStatusWithKey(ctx, invoiceStatusCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cInvoiceStatusApplication) GetCInvoiceStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cInvoiceStatusService.GetCInvoiceStatusWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
