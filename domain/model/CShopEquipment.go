package model

type CShopEquipment struct {
    ShopEquipmentCd int `json:"shop_equipment_cd"`
    LanguageCd int `json:"language_cd"`
    ShopEquipmentName string `json:"shop_equipment_name"`
    UnitName string `json:"unit_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
