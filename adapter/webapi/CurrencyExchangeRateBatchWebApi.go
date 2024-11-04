package webapi

import (
	"net/http"

	application "wellbe-common/application"

	constants "wellbe-common/share/commonsettings/constants"
	messages "wellbe-common/share/messages"

	"github.com/gin-gonic/gin"
)


type CurrencyExchangeRateBatchWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type currencyExchangeRateBatchWebApi struct {
    currencyExchangeRateApplication application.CurrencyExchangeRateBatchApplication
}

func NewCurrencyExchangeRateBatchWebApi(la application.CurrencyExchangeRateBatchApplication) CurrencyExchangeRateBatchWebApi {
    return &currencyExchangeRateBatchWebApi{
        currencyExchangeRateApplication :la,
    }
}


func (la currencyExchangeRateBatchWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.POST("/common/batch/exchange_rate/update", la.StoreCurrencyExchangeRate())

    return r
}

func (la currencyExchangeRateBatchWebApi)StoreCurrencyExchangeRate() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_BATCH {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }
        
        err := la.currencyExchangeRateApplication.StoreCurrencyExchangeRate(&ctx)
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
        })
    }
}