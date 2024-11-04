package model

type CTellCountry struct {
    LanguageCd int `json:"language_cd"`
    TellCountryCd int `json:"tell_country_cd"`
    CountryName string `json:"country_name"`
    CountryNo string `json:"country_no"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
