package model

type CShopStatus struct {
    ShopStatusCd int `json:"shop_status_cd"`
    LanguageCd int `json:"language_cd"`
    ShopStatusName string `json:"shop_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
