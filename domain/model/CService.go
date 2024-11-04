package model

type CService struct {
    ServiceCd int `json:"service_cd"`
    LanguageCd int `json:"language_cd"`
    ServiceName string `json:"service_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
