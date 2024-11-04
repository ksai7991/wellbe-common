package directory

import (
	"flag"
	log "wellbe-common/share/log"
)

func GetRootDir() string {
	logger := log.GetLogger()
	defer logger.Sync()
    var myPath string
    flag.StringVar(&myPath, "my-path", "/", "Provide project path as an absolute path")
    flag.Parse()
	logger.Debug(myPath)
	return myPath
}