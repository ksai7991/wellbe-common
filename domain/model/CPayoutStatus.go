package model

type CPayoutStatus struct {
    PayoutStatusCd int `json:"payout_status_cd"`
    LanguageCd int `json:"language_cd"`
    PayoutStatusName string `json:"payout_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
