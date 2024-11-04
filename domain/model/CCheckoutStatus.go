package model

type CCheckoutStatus struct {
    CheckoutStatusCd int `json:"checkout_status_cd"`
    LanguageCd int `json:"language_cd"`
    CheckoutStatusName string `json:"checkout_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
