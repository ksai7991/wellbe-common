package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CCurrencyRepository interface {
    CreateCCurrency(*context.Context, *model.CCurrency) (*model.CCurrency, *errordef.LogicError)
    UpdateCCurrency(*context.Context, *model.CCurrency) (*model.CCurrency, *errordef.LogicError)
    DeleteCCurrency(*context.Context, int, int) *errordef.LogicError
    GetCCurrencyWithKey(*context.Context, int,int) ([]*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithCurrencyCdIso(*context.Context, string,int) ([]*model.CCurrency, *errordef.LogicError)
    GetCCurrencyWithLanguageCd(*context.Context, int) ([]*model.CCurrency, *errordef.LogicError)
}
