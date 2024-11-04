package model

type CurrencyExchangeRateApi struct {
	BaseCurrencyCdIso string             `json:"base"`
	Rates             map[string]float64 `json:"rates"`
}