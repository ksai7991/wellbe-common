package model

type CPayoutMethod struct {
    PayoutMethodCd int `json:"payout_method_cd"`
    LanguageCd int `json:"language_cd"`
    PayoutMethodName string `json:"payout_method_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
