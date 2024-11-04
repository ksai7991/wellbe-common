package model

type CAccountWithdrawalReason struct {
    AccountWithdrawalReasonCd int `json:"account_withdrawal_reason_cd"`
    LanguageCd int `json:"language_cd"`
    AccountWithdrawalReasonName string `json:"account_withdrawal_reason_name"`
    AccountWithdrawalReasonAbbreviation string `json:"account_withdrawal_reason_abbreviation"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
