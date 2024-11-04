package model

type CCheckoutTiming struct {
    CheckoutTimingCd int `json:"checkout_timing_cd"`
    LanguageCd int `json:"language_cd"`
    CheckoutTimingName string `json:"checkout_timing_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
