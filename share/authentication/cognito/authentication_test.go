package cognito

import (
	"testing"
	env "wellbe-common/settings/env"
	"wellbe-common/share/authentication/types"

	"github.com/stretchr/testify/assert"
)

func TestRequestEmailUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
	token, _ := Login(types.LoginInfo{
		EmailAddress: "k.sasaki.tky@gmail.com",
		UserName: "22e33188-7c2d-4cb9-bfd8-05036c925fc0",
		Password: "!QAZ2wsx",
		ClientId: "30onbqb2lljvruntkl68049hqh",
		UserPoolId: "ap-northeast-1_7pTSRkG56",
	})
    err := RequestEmailUpdate(types.EmailChangeInfo{
		UserPoolId: "ap-northeast-1_7pTSRkG56",
		CurrentEmailAddress: "k.sasaki.tky@gmail.com",
		RequestedEmailAddress: "kentaro_sasaki@realize-wellbe.com",
		AccessToken: token.AccessToken,
	})
    assert.Nil(t, err)
}
func TestConfirmEmailUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
	token, _ := Login(types.LoginInfo{
		EmailAddress: "k.sasaki.tky@gmail.com",
		UserName: "22e33188-7c2d-4cb9-bfd8-05036c925fc0",
		Password: "!QAZ2wsx",
		ClientId: "30onbqb2lljvruntkl68049hqh",
		UserPoolId: "ap-northeast-1_7pTSRkG56",
	})
	err := ConfirmEmailUpdate(types.ConfirmEmailChangeInfo{
		CurrentEmailAddress: "k.sasaki.tky@gmail.com",
		AccessToken: token.AccessToken,
		ConfirmationCode: "829555", // need to update
	})
    assert.Nil(t, err)
}

func TestChangePassword(t *testing.T) {
    env.EnvLoad("./../../../")
	token, _ := Login(types.LoginInfo{
		EmailAddress: "k.sasaki.tky@gmail.com",
		UserName: "22e33188-7c2d-4cb9-bfd8-05036c925fc0",
		Password: "!QAZ3edc",
		ClientId: "30onbqb2lljvruntkl68049hqh",
		UserPoolId: "ap-northeast-1_7pTSRkG56",
	})
    err := ChangePassword(types.ChangePasswordInfo{
		PreviousPassword: "!QAZ3ed",
		ProposedPassword: "!QAZ2wsx",
		AccessToken: token.AccessToken,
	})
    assert.Nil(t, err)
    ChangePassword(types.ChangePasswordInfo{
		PreviousPassword: "!QAZ2wsx",
		ProposedPassword: "!QAZ3edc",
		AccessToken: token.AccessToken,
	})
}

func TestChangePasswordInvalid(t *testing.T) {
    env.EnvLoad("./../../../")
	token, _ := Login(types.LoginInfo{
		EmailAddress: "k.sasaki.tky@gmail.com",
		UserName: "22e33188-7c2d-4cb9-bfd8-05036c925fc0",
		Password: "!QAZ2wsx",
		ClientId: "30onbqb2lljvruntkl68049hqh",
		UserPoolId: "ap-northeast-1_7pTSRkG56",
	})
    err := ChangePassword(types.ChangePasswordInfo{
		PreviousPassword: "!QAZ2wsx",
		ProposedPassword: "abc",
		AccessToken: token.AccessToken,
	})
    assert.NotNil(t, err)
}

func TestVerify(t *testing.T) {
    env.EnvLoad("./../../../")
	token, _ := Login(types.LoginInfo{
		EmailAddress: "k.sasaki.tky@gmail.com",
		UserName: "22e33188-7c2d-4cb9-bfd8-05036c925fc0",
		Password: "!QAZ2wsx",
		ClientId: "30onbqb2lljvruntkl68049hqh",
		UserPoolId: "ap-northeast-1_7pTSRkG56",
	})
	tokenInfo := &types.TokenInfo{
		AccessToken: token.AccessToken,
		IdToken: token.IdToken,
		RefleshToken: token.RefleshToken,
	}
	user, err := Verify(tokenInfo,
	"ap-northeast-1_7pTSRkG56",
	"30onbqb2lljvruntkl68049hqh")

	assert.Nil(t, err)
	assert.Equal(t, "22e33188-7c2d-4cb9-bfd8-05036c925fc0", user.UserName)
	assert.Equal(t, "k.sasaki.tky@gmail.com", user.Email)
	assert.NotEmpty(t, tokenInfo.AccessToken)
	assert.NotEqual(t, token.AccessToken, tokenInfo.AccessToken)
}