package model

type CCheckoutMethod struct {
    CheckoutMethodCd int `json:"checkout_method_cd"`
    LanguageCd int `json:"language_cd"`
    CheckoutMethodName string `json:"checkout_method_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
