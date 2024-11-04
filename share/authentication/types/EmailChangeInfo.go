package types

type EmailChangeInfo struct {
	CurrentEmailAddress   string
	RequestedEmailAddress string
	AccessToken           string
	UserPoolId            string
	CognitoUserName       string
}