package model

type CTreatmentTimeRange struct {
    TreatmentTimeCd int `json:"treatment_time_cd"`
    LanguageCd int `json:"language_cd"`
    TreatmentTimeName string `json:"treatment_time_name"`
    MinTime int `json:"min_time"`
    MaxTime int `json:"max_time"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
