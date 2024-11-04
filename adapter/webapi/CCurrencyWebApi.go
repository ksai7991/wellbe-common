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


type CCurrencyWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cCurrencyWebApi struct {
    cCurrencyApplication application.CCurrencyApplication
}

func NewCCurrencyWebApi(la application.CCurrencyApplication) CCurrencyWebApi {
    return &cCurrencyWebApi{
        cCurrencyApplication :la,
    }
}


func (la cCurrencyWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_currencys/key", la.GetCCurrencyWithKey())
    r.GET("/c_currencys/currency_cd_iso", la.GetCCurrencyWithCurrencyCdIso())
    r.GET("/c_currencys/language_cd", la.GetCCurrencyWithLanguageCd())

    return r
}

type CCurrencyEntity struct {
    CurrencyCd string `json:"currency_cd"`
    LanguageCd string `json:"language_cd"`
    CurrencyName string `json:"currency_name"`
    CurrencyCdIso string `json:"currency_cd_iso"`
    SignificantDigit string `json:"significant_digit"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cCurrencyWebApi)GetCCurrencyWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        currencyCd_bind := c.Query("currency_cd")
        languageCd_bind := c.Query("language_cd")
        if len(currencyCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "currency_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        currencyCd, _ := strconv.Atoi(currencyCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cCurrencyApplication.GetCCurrencyWithKey(&ctx, currencyCd,languageCd)
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
            "c_currencys": results,
        })
    }
}

func (la cCurrencyWebApi)GetCCurrencyWithCurrencyCdIso() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        currencyCdIso_bind := c.Query("currency_cd_iso")
        significantDigit_bind := c.Query("significant_digit")
        if len(currencyCdIso_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "currency_cd_iso")})
            return
        }
        if len(significantDigit_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "significant_digit")})
            return
        }
        currencyCdIso := currencyCdIso_bind
        significantDigit, _ := strconv.Atoi(significantDigit_bind)
        results, err := la.cCurrencyApplication.GetCCurrencyWithCurrencyCdIso(&ctx, currencyCdIso,significantDigit)
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
            "c_currencys": results,
        })
    }
}

func (la cCurrencyWebApi)GetCCurrencyWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cCurrencyApplication.GetCCurrencyWithLanguageCd(&ctx, languageCd)
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
            "c_currencys": results,
        })
    }
}
