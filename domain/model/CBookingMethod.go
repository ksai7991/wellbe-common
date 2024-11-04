package model

type CBookingMethod struct {
    BookingMethodCd int `json:"booking_method_cd"`
    LanguageCd int `json:"language_cd"`
    BookingMethodName string `json:"booking_method_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
