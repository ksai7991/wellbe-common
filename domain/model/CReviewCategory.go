package model

type CReviewCategory struct {
    ReviewCategoryCd int `json:"review_category_cd"`
    LanguageCd int `json:"language_cd"`
    ReviewCategoryName string `json:"review_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
