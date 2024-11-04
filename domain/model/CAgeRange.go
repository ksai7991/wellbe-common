package model

type CAgeRange struct {
    AgeRangeCd int `json:"age_range_cd"`
    LanguageCd int `json:"language_cd"`
    AgeRangeGender string `json:"age_range_gender"`
    AgeRangeFrom int `json:"age_range_from"`
    AgeRangeTo int `json:"age_range_to"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
