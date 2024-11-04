package model

type CContentsCategory struct {
    ContentsCategoryCd int `json:"contents_category_cd"`
    LanguageCd int `json:"language_cd"`
    ContentsCategoryName string `json:"contents_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
