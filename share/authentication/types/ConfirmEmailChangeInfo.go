package types

type ConfirmEmailChangeInfo struct {
	CurrentEmailAddress string
	AccessToken         string
	ConfirmationCode    string
}