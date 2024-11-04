package model

type CState struct {
    StateCd int `json:"state_cd"`
    LanguageCd int `json:"language_cd"`
    CountryCd int `json:"country_cd"`
    StateName string `json:"state_name"`
    StateCdIso string `json:"state_cd_iso"`
    TimezoneIana string `json:"timezone_iana"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
