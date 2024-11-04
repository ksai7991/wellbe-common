package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CurrencyForPaymentRepository interface {
    CreateCurrencyForPayment(*context.Context, *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError)
    UpdateCurrencyForPayment(*context.Context, *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError)
    DeleteCurrencyForPayment(*context.Context, int) *errordef.LogicError
    GetCurrencyForPaymentWithKey(*context.Context, int) ([]*model.CurrencyForPayment, *errordef.LogicError)
}
