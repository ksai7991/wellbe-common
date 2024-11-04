package model

type CCouponTargetAttr struct {
    CouponTargetAttrCd int `json:"coupon_target_attr_cd"`
    LanguageCd int `json:"language_cd"`
    CouponTargetAttrName string `json:"coupon_target_attr_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
