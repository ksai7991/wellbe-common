package cognito

import (
	"context"
	"fmt"
	"time"
	"wellbe-common/settings"
	types "wellbe-common/share/authentication/types"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
	"wellbe-common/share/messages"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/dgrijalva/jwt-go"
	jwk "github.com/lestrrat-go/jwx/jwk"
)

func SignUp(signUpInfo types.SignUpInfo) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	ua := &cognito.AttributeType {
			Name: aws.String("email"),
			Value: aws.String(signUpInfo.EmailAddress),
	}
	cognitoClient := cognito.New(sess)
	user := &cognito.SignUpInput{
		Username: aws.String(signUpInfo.UserName),
		Password: aws.String(signUpInfo.Password),
		ClientId: aws.String(signUpInfo.ClientId),
		UserAttributes: []*cognito.AttributeType{
				ua,
		},
	}

	_, err := cognitoClient.SignUp(user)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_SIGNUP}
	}
	
	return nil
}

func ConfirmSignUp(confirmSignUpInfo types.ConfirmSignUpInfo) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	user := &cognito.ConfirmSignUpInput{
		ConfirmationCode: aws.String(confirmSignUpInfo.ConfirmationCode),
		Username:         aws.String(confirmSignUpInfo.UserName),
		ClientId:         aws.String(confirmSignUpInfo.ClientId),
	}

	_, err := cognitoClient.ConfirmSignUp(user)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOCIG_ERROR_CODE_CANNOT_CONFIRM_SIGNUP}
	}
	
	return nil
}

func Login(loginInfo types.LoginInfo) (*types.TokenInfo, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	params := map[string]*string{
		"USERNAME": aws.String(loginInfo.UserName),
		"PASSWORD": aws.String(loginInfo.Password),
	}

	authTry := &cognito.AdminInitiateAuthInput{
		AuthFlow:       aws.String("ADMIN_USER_PASSWORD_AUTH"),
		AuthParameters: params,
		ClientId:       aws.String(loginInfo.ClientId),
		UserPoolId:     aws.String(loginInfo.UserPoolId),
	}

	req, resp := cognitoClient.AdminInitiateAuthRequest(authTry)
	err := req.Send()
	if err != nil {
		return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_LOGIN}
	}
	tokenInfo := types.TokenInfo{}
	tokenInfo.IdToken = *resp.AuthenticationResult.IdToken
	tokenInfo.AccessToken = *resp.AuthenticationResult.AccessToken
	tokenInfo.RefleshToken = *resp.AuthenticationResult.RefreshToken
	return &tokenInfo, nil
}

func Refresh(loginInfo types.LoginInfo, tokenInfo *types.TokenInfo) (*types.TokenInfo, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	params := map[string]*string{
		"REFRESH_TOKEN": aws.String(tokenInfo.RefleshToken),
	}

	authTry := &cognito.AdminInitiateAuthInput{
		AuthFlow:       aws.String("REFRESH_TOKEN_AUTH"),
		AuthParameters: params,
		ClientId:       aws.String(loginInfo.ClientId),
		UserPoolId:     aws.String(loginInfo.UserPoolId),
	}

	req, resp := cognitoClient.AdminInitiateAuthRequest(authTry)
	err := req.Send()
	if err != nil {
		logger.Error(err.Error())
		return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_LOGIN}
	}
	tokenInfo.IdToken = *resp.AuthenticationResult.IdToken
	tokenInfo.AccessToken = *resp.AuthenticationResult.AccessToken
	return tokenInfo, nil
}

func RequestEmailUpdate(emailChangeInfo types.EmailChangeInfo) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	updateInput := &cognito.UpdateUserAttributesInput{
		AccessToken: aws.String(emailChangeInfo.AccessToken),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(emailChangeInfo.RequestedEmailAddress),
			},
		},
	}
	_, err := cognitoClient.UpdateUserAttributes(updateInput)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_CHANGE_EMAIL}
	}
	adminUpdateInput := &cognito.AdminUpdateUserAttributesInput{
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(emailChangeInfo.CurrentEmailAddress),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
		UserPoolId: aws.String(emailChangeInfo.UserPoolId),
		Username:   aws.String(emailChangeInfo.CognitoUserName),
	}
	_, err = cognitoClient.AdminUpdateUserAttributes(adminUpdateInput)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_CHANGE_EMAIL}
	}
	return nil
}

func ConfirmEmailUpdate(confirmEmailChangeInfo types.ConfirmEmailChangeInfo) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	verifyInput := &cognito.VerifyUserAttributeInput{
		AccessToken:   aws.String(confirmEmailChangeInfo.AccessToken),
		AttributeName: aws.String("email"),
		Code:          aws.String(confirmEmailChangeInfo.ConfirmationCode),
	}
	_, err := cognitoClient.VerifyUserAttribute(verifyInput)
	if err != nil {
		logger.Error(err.Error())
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_CHANGE_EMAIL}
	}
	return nil
}

func ChangePassword(changePasswordInfo types.ChangePasswordInfo) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	params := &cognito.ChangePasswordInput{
			AccessToken:      aws.String(changePasswordInfo.AccessToken),
			PreviousPassword: aws.String(changePasswordInfo.PreviousPassword),
			ProposedPassword: aws.String(changePasswordInfo.ProposedPassword),
	}

	_, err := cognitoClient.ChangePassword(params)
	if err != nil {
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_CHANGE_PASSWORD}
	}
	return nil
}

func Verify(tokenInfo *types.TokenInfo, userPoolId string, clientId string) (*types.UserInfo ,*errordef.LogicError){
    logger := log.GetLogger()
    defer logger.Sync()
    regionName := settings.GetAWSRegionName()
    keySet, err := jwk.Fetch(context.Background(), fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", regionName, userPoolId))
    if err != nil {
        logger.Error(err.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code:constants.LOGIC_ERROR_CODE_JWK_NOTEXISTS}
    }

	// idプロバイダーとの時間ずれを調整
	time.Sleep(1 * time.Second)
    token, err := jwt.Parse(tokenInfo.IdToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            errInternal := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            return nil, errInternal
        }
        kid, ok := token.Header["kid"].(string)
        if !ok {
            errInternal := fmt.Errorf("kid header not found")
            return nil, errInternal
        }
        keys, ok := keySet.LookupKeyID(kid);
        if !ok {
            errInternal := fmt.Errorf("key with specified kid is not present in jwks")
            return nil, errInternal
        }
        var publickey interface{}
        err = keys.Raw(&publickey)
        if err != nil {
            errInternal := fmt.Errorf("could not parse pubkey")
            return nil, errInternal
        }
        return publickey, nil
    })
    
    if token == nil {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_PLEASE_LOGIN, Code:constants.LOCIG_ERROR_CODE_TOKEN_PARSE_ERROR}
    }
    
    claims, ok := token.Claims.(jwt.MapClaims);
    if ok && token.Valid {
    } else {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_PLEASE_LOGIN, Code:constants.LOCIG_ERROR_CODE_TOKEN_PARSE_ERROR}
    }

	userInfo := &types.UserInfo{}
	userInfo.Email = claims["email"].(string)
	userInfo.UserName = claims["cognito:username"].(string)

    return userInfo, nil
}

func DeleteAccount(deleteUserInfo types.DeleteUserInfo) (*errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    configure := settings.GetAWSConfigure()
    sess := session.Must(session.NewSession(configure))
	cognitoClient := cognito.New(sess)
	params := &cognito.DeleteUserInput{
			AccessToken:      aws.String(deleteUserInfo.AccessToken),
	}

	_, err := cognitoClient.DeleteUser(params)
	if err != nil {
		return &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_CANNOT_CHANGE_PASSWORD}
	}
	return nil
}