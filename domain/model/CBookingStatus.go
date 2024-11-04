package model

type CBookingStatus struct {
    BookingStatusCd int `json:"booking_status_cd"`
    LanguageCd int `json:"language_cd"`
    BookingStatusName string `json:"booking_status_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
