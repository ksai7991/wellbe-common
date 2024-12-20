package webapi

import (
    "net/http"
    
    application "wellbe-common/application"

    constants "wellbe-common/share/commonsettings/constants"
    messages "wellbe-common/share/messages"
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
)


type CAccountWithdrawalReasonWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cAccountWithdrawalReasonWebApi struct {
    cAccountWithdrawalReasonApplication application.CAccountWithdrawalReasonApplication
}

func NewCAccountWithdrawalReasonWebApi(la application.CAccountWithdrawalReasonApplication) CAccountWithdrawalReasonWebApi {
    return &cAccountWithdrawalReasonWebApi{
        cAccountWithdrawalReasonApplication :la,
    }
}


func (la cAccountWithdrawalReasonWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_account_withdrawal_reasons/key", la.GetCAccountWithdrawalReasonWithKey())
    r.GET("/c_account_withdrawal_reasons/language_cd", la.GetCAccountWithdrawalReasonWithLanguageCd())

    return r
}

type CAccountWithdrawalReasonEntity struct {
    AccountWithdrawalReasonCd string `json:"account_withdrawal_reason_cd"`
    LanguageCd string `json:"language_cd"`
    AccountWithdrawalReasonName string `json:"account_withdrawal_reason_name"`
    AccountWithdrawalReasonAbbreviation string `json:"account_withdrawal_reason_abbreviation"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cAccountWithdrawalReasonWebApi)GetCAccountWithdrawalReasonWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        accountWithdrawalReasonCd_bind := c.Query("account_withdrawal_reason_cd")
        languageCd_bind := c.Query("language_cd")
        if len(accountWithdrawalReasonCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "account_withdrawal_reason_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        accountWithdrawalReasonCd, _ := strconv.Atoi(accountWithdrawalReasonCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cAccountWithdrawalReasonApplication.GetCAccountWithdrawalReasonWithKey(&ctx, accountWithdrawalReasonCd,languageCd)
        if err != nil {
            if err.Code >= 900 {
                c.JSON(http.StatusInternalServerError, gin.H{
                    constants.WEBAPI_RESPONSE_KEYWORD_STATUS:constants.LOGIC_ERROR_CODE_SEVERERROR,
                    constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: messages.MESSAGE_EN_SERVER_ERROR,
                })
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{
                    constants.WEBAPI_RESPONSE_KEYWORD_STATUS: err.Code,
                    constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: err.Msg,
                })
            }
            return
        }

        c.JSON(http.StatusOK, gin.H{
            constants.WEBAPI_RESPONSE_KEYWORD_STATUS:constants.LOGIC_ERROR_CODE_SUCCESS,
            constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: messages.MESSAGE_EN_SUCCESS,
            "c_account_withdrawal_reasons": results,
        })
    }
}

func (la cAccountWithdrawalReasonWebApi)GetCAccountWithdrawalReasonWithLanguageCd() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        languageCd_bind := c.Query("language_cd")
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cAccountWithdrawalReasonApplication.GetCAccountWithdrawalReasonWithLanguageCd(&ctx, languageCd)
        if err != nil {
            if err.Code >= 900 {
                c.JSON(http.StatusInternalServerError, gin.H{
                    constants.WEBAPI_RESPONSE_KEYWORD_STATUS:constants.LOGIC_ERROR_CODE_SEVERERROR,
                    constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: messages.MESSAGE_EN_SERVER_ERROR,
                })
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{
                    constants.WEBAPI_RESPONSE_KEYWORD_STATUS: err.Code,
                    constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: err.Msg,
                })
            }
            return
        }

        c.JSON(http.StatusOK, gin.H{
            constants.WEBAPI_RESPONSE_KEYWORD_STATUS:constants.LOGIC_ERROR_CODE_SUCCESS,
            constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: messages.MESSAGE_EN_SUCCESS,
            "c_account_withdrawal_reasons": results,
        })
    }
}
