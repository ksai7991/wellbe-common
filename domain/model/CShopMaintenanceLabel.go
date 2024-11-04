package model

type CShopMaintenanceLabel struct {
    ShopMaintenanceLabelCd int `json:"shop_maintenance_label_cd"`
    LanguageCd int `json:"language_cd"`
    ShopMaintenanceLabelName string `json:"shop_maintenance_label_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
