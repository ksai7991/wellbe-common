package model

type COrderType struct {
    OrderTypeCd int `json:"order_type_cd"`
    LanguageCd int `json:"language_cd"`
    OrderTypeName string `json:"order_type_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
