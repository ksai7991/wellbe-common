package model

type DefaultFeeMaster struct {
    Id string `json:"id"`
    FeeRate float64 `json:"fee_rate"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
