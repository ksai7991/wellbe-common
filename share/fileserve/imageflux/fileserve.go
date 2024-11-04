package imageflux

import (
	"strings"
	"wellbe-common/settings"
	constants "wellbe-common/share/commonsettings/constants"
)

  func GetImageFluxImagePublishedPath(fullpath string, width string) string {
    imgixDomain := settings.GetImgixDomainName()
    path := GetKey(fullpath)
    
    return constants.HTTPS_PROTOCOL + imgixDomain + "/w="+width+",a=0,f=webp:auto/" + path
  }

  func GetKey(fullpath string) string {
    splFullPath := strings.Split(fullpath, "/")
    return splFullPath[len(splFullPath)-1]
  }