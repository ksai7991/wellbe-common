package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CurrencyExchangeRateRepository interface {
    CreateCurrencyExchangeRate(*context.Context, *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    UpdateCurrencyExchangeRate(*context.Context, *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    DeleteCurrencyExchangeRate(*context.Context, int, int) *errordef.LogicError
    GetCurrencyExchangeRateWithKey(*context.Context, int,int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithPaireName(*context.Context, string) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
    GetCurrencyExchangeRateWithBase(*context.Context, int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
}
