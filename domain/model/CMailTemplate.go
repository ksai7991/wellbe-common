package model

type CMailTemplate struct {
    MailTemplateCd int `json:"mail_template_cd"`
    LanguageCd int `json:"language_cd"`
    Subject string `json:"subject"`
    Body string `json:"body"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
