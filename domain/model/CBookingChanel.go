package model

type CBookingChanel struct {
    BookingChanelCd int `json:"booking_chanel_cd"`
    LanguageCd int `json:"language_cd"`
    BookingChanelName string `json:"booking_chanel_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
