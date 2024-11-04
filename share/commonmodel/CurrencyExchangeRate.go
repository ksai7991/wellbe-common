package commonmodel

type CurrencyExchangeRate struct {
	BaseCurrencyCd   int     `json:"base_currency_cd"`
	TargetCurrencyCd int     `json:"target_currency_cd"`
	PaireName        string  `json:"paire_name"`
	Rate             float64 `json:"rate"`
	CreateDatetime   string  `json:"create_datetime"`
	CreateFunction   string  `json:"create_function"`
	UpdateDatetime   string  `json:"update_datetime"`
	UpdateFunction   string  `json:"update_function"`
}
