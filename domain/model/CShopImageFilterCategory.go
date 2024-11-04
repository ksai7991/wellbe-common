package model

type CShopImageFilterCategory struct {
    ShopImageFilterCategoryCd int `json:"shop_image_filter_category_cd"`
    LanguageCd int `json:"language_cd"`
    ShopImageFilterCategoryName string `json:"shop_image_filter_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
