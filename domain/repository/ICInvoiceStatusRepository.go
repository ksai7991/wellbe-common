package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CInvoiceStatusRepository interface {
    CreateCInvoiceStatus(*context.Context, *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError)
    UpdateCInvoiceStatus(*context.Context, *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError)
    DeleteCInvoiceStatus(*context.Context, int, int) *errordef.LogicError
    GetCInvoiceStatusWithKey(*context.Context, int,int) ([]*model.CInvoiceStatus, *errordef.LogicError)
    GetCInvoiceStatusWithLanguageCd(*context.Context, int) ([]*model.CInvoiceStatus, *errordef.LogicError)
}
