package model

type CReviewContent struct {
    ReviewContentCd int `json:"review_content_cd"`
    LanguageCd int `json:"language_cd"`
    ReviewCategoryCd int `json:"review_category_cd"`
    ReviewContentName string `json:"review_content_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
