package model

type CBillingContent struct {
    BillingContentCd int `json:"billing_content_cd"`
    LanguageCd int `json:"language_cd"`
    BillingContentName string `json:"billing_content_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
