package model

type CWeekday struct {
    WeekdayCd int `json:"weekday_cd"`
    LanguageCd int `json:"language_cd"`
    WeekdayName string `json:"weekday_name"`
    WeekdayAbbreviation string `json:"weekday_abbreviation"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
