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


type CurrencyForPaymentWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type currencyForPaymentWebApi struct {
    currencyForPaymentApplication application.CurrencyForPaymentApplication
}

func NewCurrencyForPaymentWebApi(la application.CurrencyForPaymentApplication) CurrencyForPaymentWebApi {
    return &currencyForPaymentWebApi{
        currencyForPaymentApplication :la,
    }
}


func (la currencyForPaymentWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/currency_for_payments/key", la.GetCurrencyForPaymentWithKey())

    return r
}

type CurrencyForPaymentEntity struct {
    CurrencyCd string `json:"currency_cd"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la currencyForPaymentWebApi)GetCurrencyForPaymentWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        currencyCd_bind := c.Query("currency_cd")
        if len(currencyCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "currency_cd")})
            return
        }
        currencyCd, _ := strconv.Atoi(currencyCd_bind)
        results, err := la.currencyForPaymentApplication.GetCurrencyForPaymentWithKey(&ctx, currencyCd)
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
            "currency_for_payments": results,
        })
    }
}
