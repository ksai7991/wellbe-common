package commonmodel

type CCurrency struct {
	CurrencyCd       int    `json:"currency_cd"`
	LanguageCd       int    `json:"language_cd"`
	CurrencyName     string `json:"currency_name"`
	CurrencyCdIso    string `json:"currency_cd_iso"`
	SignificantDigit int    `json:"significant_digit"`
	CreateDatetime   string `json:"create_datetime"`
	CreateFunction   string `json:"create_function"`
	UpdateDatetime   string `json:"update_datetime"`
	UpdateFunction   string `json:"update_function"`
}
