package authentication

import (
	cognito "wellbe-common/share/authentication/cognito"
	types "wellbe-common/share/authentication/types"
	errordef "wellbe-common/share/errordef"
)

func SignUp(signUpInfo types.SignUpInfo) (*errordef.LogicError) {
	return cognito.SignUp(signUpInfo)
}

func ConfirmSignUp(confirmSignUpInfo types.ConfirmSignUpInfo) (*errordef.LogicError) {
	return cognito.ConfirmSignUp(confirmSignUpInfo)
}

func Login(loginInfo types.LoginInfo) (*types.TokenInfo, *errordef.LogicError) {
	return cognito.Login(loginInfo)
}

func Refresh(loginInfo types.LoginInfo, tokenInfo *types.TokenInfo) (*types.TokenInfo, *errordef.LogicError) {
	return cognito.Refresh(loginInfo, tokenInfo)
}

func RequestEmailUpdate(emailChangeInfo types.EmailChangeInfo) (*errordef.LogicError) {
	return cognito.RequestEmailUpdate(emailChangeInfo)
}

func ConfirmEmailUpdate(confirmEmailChangeInfo types.ConfirmEmailChangeInfo) (*errordef.LogicError) {
	return cognito.ConfirmEmailUpdate(confirmEmailChangeInfo)
}

func ChangePassword(changePasswordInfo types.ChangePasswordInfo) (*errordef.LogicError) {
	return cognito.ChangePassword(changePasswordInfo)
}

func Verify(tokenInfo *types.TokenInfo, userPoolId string, clientId string) (*types.UserInfo ,*errordef.LogicError){
	return cognito.Verify(tokenInfo, userPoolId, clientId)
}

func DeleteAccount(deleteUserInfo types.DeleteUserInfo) (*errordef.LogicError){
	return cognito.DeleteAccount(deleteUserInfo)
}