package model

type CGender struct {
    GenderCd int `json:"gender_cd"`
    LanguageCd int `json:"language_cd"`
    GenderName string `json:"gender_name"`
    GenderAbbreviation string `json:"gender_abbreviation"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
