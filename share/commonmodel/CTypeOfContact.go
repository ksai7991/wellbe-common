package commonmodel

type CTypeOfContact struct {
	TypeOfContactCd   int    `json:"type_of_contact_cd"`
	LanguageCd        int    `json:"language_cd"`
	TypeOfContactName string `json:"type_of_contact_name"`
	CreateDatetime    string `json:"create_datetime"`
	CreateFunction    string `json:"create_function"`
	UpdateDatetime    string `json:"update_datetime"`
	UpdateFunction    string `json:"update_function"`
}
