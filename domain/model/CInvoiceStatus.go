package model

type CInvoiceStatus struct {
    InvoiceStatusCd int `json:"invoice_status_cd"`
    LanguageCd int `json:"language_cd"`
    InvoiceStatusName string `json:"invoice_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
