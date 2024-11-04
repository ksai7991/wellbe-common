package util

import "fmt"

func GetMapUrl(latitude string, longitude string) string {
	if latitude != "" && longitude != "" {
		return fmt.Sprintf("http://maps.google.com/maps?q=%v,%v", latitude, longitude)
	} else {
		return ""
	}
}