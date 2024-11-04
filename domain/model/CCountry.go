package model

type CCountry struct {
    CountryCd int `json:"country_cd"`
    LanguageCd int `json:"language_cd"`
    CountryName string `json:"country_name"`
    CountryCdIso string `json:"country_cd_iso"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
