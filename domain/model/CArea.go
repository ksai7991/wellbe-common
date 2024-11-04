package model

type CArea struct {
    LanguageCd int `json:"language_cd"`
    AreaCd int `json:"area_cd"`
    StateCd int `json:"state_cd"`
    AreaName string `json:"area_name"`
    SearchAreaNameSeo string `json:"search_area_name_seo"`
    WestLongitude float64 `json:"west_longitude"`
    EastLongitude float64 `json:"east_longitude"`
    NorthLatitude float64 `json:"north_latitude"`
    SouthLatitude float64 `json:"south_latitude"`
    SummaryAreaFlg bool `json:"summary_area_flg"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}
