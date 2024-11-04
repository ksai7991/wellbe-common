package model

type CRecommendLabel struct {
    RecommendLabelCd int `json:"recommend_label_cd"`
    LanguageCd int `json:"language_cd"`
    RecommendLabelName string `json:"recommend_label_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
