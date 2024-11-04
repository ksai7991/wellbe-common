package model

type CShopPaymentMethod struct {
    ShopPaymentMethodCd int `json:"shop_payment_method_cd"`
    LanguageCd int `json:"language_cd"`
    ShopPaymentName string `json:"shop_payment_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
