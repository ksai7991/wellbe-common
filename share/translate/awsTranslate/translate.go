package translate

import (
	"wellbe-common/settings"
	"wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	"wellbe-common/share/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

func GetTranslate(text string, src string, target string) (string, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()

    sess := session.Must(session.NewSession(configure))
    trs := translate.New(sess)

    result, err := trs.Text(&translate.TextInput{
        SourceLanguageCode: aws.String(src),
        TargetLanguageCode: aws.String(target),
        TerminologyNames: aws.StringSlice([]string{constants.TRANSLATE_CUSTOM, constants.TRANSLATE_SHOP}),
        Text:               aws.String(text),
    })
    if err != nil {
        logger.Error(err.Error())
        return "", &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_TRANSLATE_ERROR}
    }

	return *result.TranslatedText, nil
}

func ImportTerminology(name string, body []byte) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()

    sess := session.Must(session.NewSession(configure))
    trs := translate.New(sess)

    _, err := trs.ImportTerminology(&translate.ImportTerminologyInput{
        Name: aws.String(name),
        MergeStrategy: aws.String(constants.TRANSLATE_MERGE_STRATEGY),
        TerminologyData: &translate.TerminologyData{
            Directionality: aws.String(constants.TRANSLATE_DIRECTIONALITY_MULTI),
            Format: aws.String(constants.TRANSLATE_FILE_FORMAT_CSV),
            File: body,
        },
    })
    if err != nil {
        logger.Error(err.Error())
        return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_TRANSLATE_ERROR}
    }

	return nil
}
