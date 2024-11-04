package model

type CShopContractPlanItem struct {
    ShopContractPlanItemCd int `json:"shop_contract_plan_item_cd"`
    LanguageCd int `json:"language_cd"`
    ShopContractPlanName string `json:"shop_contract_plan_name"`
    Unit string `json:"unit"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
