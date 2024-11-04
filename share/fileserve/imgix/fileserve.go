package s3

import (
	"strings"
	"wellbe-common/settings"
	constants "wellbe-common/share/commonsettings/constants"
)

  func GetImigxImagePublishedPath(fullpath string) string {
    imgixDomain := settings.GetImgixDomainName()
    path := GetKey(fullpath)
    
    return constants.HTTPS_PROTOCOL + imgixDomain + "/" + path
  }

  func GetKey(fullpath string) string {
    splFullPath := strings.Split(fullpath, "/")
    return splFullPath[len(splFullPath)-1]
  }