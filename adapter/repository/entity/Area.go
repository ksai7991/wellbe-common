package entity

type Area struct {
	AreaCd            int    `json:"area_cd"`
	AreaName          string `json:"area_name"`
	SearchAreaNameSeo string `json:"search_area_name_seo"`
	StateCd           int    `json:"state_cd"`
	StateName         string `json:"state_name"`
	StateCdIso        string `json:"state_cd_iso"`
	CountryCd         int    `json:"country_cd"`
	CountryName       string `json:"country_name"`
	CountryCdIso      string `json:"country_cd_iso"`
}