package translate

import (
	"wellbe-common/settings"
	errordef "wellbe-common/share/errordef"
	tr "wellbe-common/share/translate/awsTranslate"

	commonconstants "wellbe-common/share/commonsettings/constants"
)


func GetTranslate(text string, src string, target string) (string, *errordef.LogicError) {
	if settings.GetTranslateStab() == commonconstants.ENV_WELLBE_AWS_TRANSLATE_STAB_TRUE {
		return text, nil
	} else {
		return tr.GetTranslate(text, src, target)
	}
}

func ImportTerminology(name string, body []byte) (*errordef.LogicError) {
	return tr.ImportTerminology(name, body)
}