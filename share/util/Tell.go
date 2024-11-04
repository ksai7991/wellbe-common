package util

import (
	"regexp"
	"strings"
)

func ConvertToInternationalPhoneNumber(phoneNumber string, countryCode string) string {
if len(phoneNumber) == 0 {
		return ""
	}
	if len(countryCode) == 0 {
		return phoneNumber
	}
	// 電話番号に含まれる不要な文字を削除します
	regex := regexp.MustCompile(`\D`)
	cleanedPhoneNumber := regex.ReplaceAllString(phoneNumber, "")

	// 先頭のゼロを削除します
	trimmedPhoneNumber := strings.TrimLeft(cleanedPhoneNumber, "0")

	// 国際電話番号の形式に変換します
	internationalPhoneNumber := "+" + countryCode + trimmedPhoneNumber

	return internationalPhoneNumber
}
