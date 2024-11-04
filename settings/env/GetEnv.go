package env

import (
	"fmt"
	"os"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"github.com/joho/godotenv"
)
func EnvLoad(defaultFolder string) *errordef.LogicError {
	logger := log.GetLogger()
	defer logger.Sync()
	if len(os.Getenv(constants.ENV_KEYWORD)) == 0 {
		os.Setenv(constants.ENV_KEYWORD, "local")
	}
	err := godotenv.Load(fmt.Sprintf(defaultFolder+"settings/env/common-%s.env", os.Getenv(constants.ENV_KEYWORD)))
	if err != nil {
		logger.Error("Error loading env target env is " + os.Getenv(constants.ENV_KEYWORD))
		return &errordef.LogicError{Msg: "Error loading env target env is " + os.Getenv(constants.ENV_KEYWORD), Code: constants.LOGIC_ERROR_CODE_ENVERROR}
	}
	return nil
}