package mail

import (
	"strings"
	"wellbe-common/settings"
	model "wellbe-common/share/commonmodel"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func Send(mail *model.Mail) (*errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	err := sesSend(constants.NOREPLY_MAIL_ADDRESS, mail.ToAddress, mail.SubjectText, mail.BodyText)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_MAIERERROR}
	}
	return nil
}

func sesSend(from string, to string, title string, body string) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    bodyFormat := removeEscape(body)
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
    svc := ses.New(sess)
    input := &ses.SendEmailInput{
    Destination: &ses.Destination{
        ToAddresses: []*string{
        aws.String(to),
        },
    },
    Message: &ses.Message{
        Body: &ses.Body{
        Text: &ses.Content{
            Charset: aws.String("UTF-8"),
            Data:    aws.String(bodyFormat),
        },
        },
        Subject: &ses.Content{
        Charset: aws.String("UTF-8"),
        Data:    aws.String(title),
        },
    },
    Source: aws.String(from),
    }
    _, err := svc.SendEmail(input)
    if err != nil {
        logger.Error(err.Error())
        return &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_MAIERERROR, Msg: err.Error()}
    }
    return nil
}

func removeEscape(msg string) string {
    return strings.Replace(msg, "\\r\\n", "\r\n", -1)
}