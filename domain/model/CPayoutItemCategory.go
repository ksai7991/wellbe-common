package model

type CPayoutItemCategory struct {
    PayoutItemCategoryCd int `json:"payout_item_category_cd"`
    LanguageCd int `json:"language_cd"`
    PayoutItemCategoryName string `json:"payout_item_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
