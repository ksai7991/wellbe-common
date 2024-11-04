package recaptcha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"wellbe-common/settings"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	"wellbe-common/share/log"
	messages "wellbe-common/share/messages"
)

type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func ReCaptchaVerify(token string) (*errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"
	secret := settings.GetRecaptchSecretKey()
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {secret}, "response": {token}})
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_SEVERERROR, Msg: messages.MESSAGE_EN_SERVER_ERROR}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_SEVERERROR, Msg: messages.MESSAGE_EN_SERVER_ERROR}
	}
	var r RecaptchaResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_SEVERERROR, Msg: messages.MESSAGE_EN_SERVER_ERROR}
	}
	if r.Success == false {
		return &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_RECAPTCHA_FAIL, Msg: fmt.Sprintf(messages.MESSAGE_EN_RECAPTCHA_FAIL, r.ErrorCodes)}
	} else{
		return nil
	}
}
