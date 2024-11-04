package model

type CConcern struct {
    ConcernCd int `json:"concern_cd"`
    LanguageCd int `json:"language_cd"`
    ConcernName string `json:"concern_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
