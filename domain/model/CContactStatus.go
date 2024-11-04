package model

type CContactStatus struct {
    ContactStatusCd int `json:"contact_status_cd"`
    LanguageCd int `json:"language_cd"`
    ContactStatusName string `json:"contact_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
