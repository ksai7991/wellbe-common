package apiclient

import (
	"context"
	"wellbe-common/domain/model"
	errordef "wellbe-common/share/errordef"
)
type CurrencyExchangeRateApiclient interface {
    GetRate(*context.Context, string) (*model.CurrencyExchangeRateApi, *errordef.LogicError)
}
