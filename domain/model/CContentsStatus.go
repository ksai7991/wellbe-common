package model

type CContentsStatus struct {
    ContentsStatusCd int `json:"contents_status_cd"`
    LanguageCd int `json:"language_cd"`
    ContentsStatusName string `json:"contents_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
