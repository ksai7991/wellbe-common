package util

import (
	"wellbe-common/share/commonmodel"
)

func GetLanguageCdChar(languages []*commonmodel.CLanguage, languageCd int) string {
	for _, v := range languages {
		if v.LanguageCd == languageCd {
			return v.LanguageCharCd
		}
	}

	return ""
}

func GetLanguagesExceptSource(languages []*commonmodel.CLanguage, languageCd int) ([]*commonmodel.CLanguage) {
	results := []*commonmodel.CLanguage{}
	for _, v := range languages {
		if v.LanguageCd != languageCd {
			results = append(results, v)
		}
	}

	return results
}
