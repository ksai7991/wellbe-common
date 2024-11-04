package model

type CMenuLabel struct {
    MenuLabelCd int `json:"menu_label_cd"`
    LanguageCd int `json:"language_cd"`
    MenuLabelName string `json:"menu_label_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}