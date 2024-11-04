package model

type CBillingStatus struct {
    BillingStatusCd int `json:"billing_status_cd"`
    LanguageCd int `json:"language_cd"`
    BillingStatusName string `json:"billing_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
