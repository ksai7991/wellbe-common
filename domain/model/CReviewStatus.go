package model

type CReviewStatus struct {
    ReviewStatusCd int `json:"review_status_cd"`
    LanguageCd int `json:"language_cd"`
    ReviewStatusName string `json:"review_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
