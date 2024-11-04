package model

type CLanguage struct {
    LanguageCd int `json:"language_cd"`
    LanguageCharCd string `json:"language_char_cd"`
    LanguageName string `json:"language_name"`
    SortNumber int `json:"sort_number"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
