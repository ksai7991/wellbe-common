package model

type CContentsLabel struct {
    ContentsLabelCd int `json:"contents_label_cd"`
    LanguageCd int `json:"language_cd"`
    ContentsCategoryCd int `json:"contents_category_cd"`
    ContentsLabelName string `json:"contents_label_name"`
    ContentsLabelUrl string `json:"contents_label_url"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
