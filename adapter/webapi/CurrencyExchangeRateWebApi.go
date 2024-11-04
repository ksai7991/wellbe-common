package webapi

import (
    "net/http"
    
    application "wellbe-common/application"

    constants "wellbe-common/share/commonsettings/constants"
    messages "wellbe-common/share/messages"
    model "wellbe-common/domain/model"
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"
)


type CurrencyExchangeRateWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type currencyExchangeRateWebApi struct {
    currencyExchangeRateApplication application.CurrencyExchangeRateApplication
}

func NewCurrencyExchangeRateWebApi(la application.CurrencyExchangeRateApplication) CurrencyExchangeRateWebApi {
    return &currencyExchangeRateWebApi{
        currencyExchangeRateApplication :la,
    }
}


func (la currencyExchangeRateWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.POST("/cud/currency_exchange_rate/create", la.CreateCurrencyExchangeRate())
    r.POST("/cud/currency_exchange_rate/update", la.UpdateCurrencyExchangeRate())
    r.POST("/cud/currency_exchange_rate/delete", la.DeleteCurrencyExchangeRate())
    r.GET("/currency_exchange_rates/key", la.GetCurrencyExchangeRateWithKey())
    r.GET("/currency_exchange_rates/paire_name", la.GetCurrencyExchangeRateWithPaireName())
    r.GET("/currency_exchange_rates/base", la.GetCurrencyExchangeRateWithBase())

    return r
}

type CurrencyExchangeRateEntity struct {
    BaseCurrencyCd string `json:"base_currency_cd"`
    TargetCurrencyCd string `json:"target_currency_cd"`
    PaireName string `json:"paire_name"`
    Rate string `json:"rate"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la currencyExchangeRateWebApi)CreateCurrencyExchangeRate() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        requestBody := CurrencyExchangeRateEntity{}
        c.Bind(&requestBody)
        if len(requestBody.BaseCurrencyCd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "base_currency_cd")})
            return
        }
        if len(requestBody.TargetCurrencyCd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "target_currency_cd")})
            return
        }
        if len(requestBody.PaireName) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "paire_name")})
            return
        }
        if len(requestBody.Rate) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "rate")})
            return
        }
        if len(requestBody.PaireName) >= 6 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_OUT_OF_LENGTH, "6", "paire_name")})
            return
        }
        currencyExchangeRate := model.CurrencyExchangeRate{}
        currencyExchangeRate.BaseCurrencyCd, _ = strconv.Atoi(requestBody.BaseCurrencyCd)
        currencyExchangeRate.TargetCurrencyCd, _ = strconv.Atoi(requestBody.TargetCurrencyCd)

        currencyExchangeRate.Rate,_ = strconv.ParseFloat(requestBody.Rate, 64)




        result, err := la.currencyExchangeRateApplication.CreateCurrencyExchangeRate(&ctx, &currencyExchangeRate)
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
            "currency_exchange_rate": result,
        })
    }
}

func (la currencyExchangeRateWebApi)UpdateCurrencyExchangeRate() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        requestBody := CurrencyExchangeRateEntity{}
        c.Bind(&requestBody)
        if len(requestBody.BaseCurrencyCd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "base_currency_cd")})
            return
        }
        if len(requestBody.TargetCurrencyCd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "target_currency_cd")})
            return
        }
        if len(requestBody.PaireName) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "paire_name")})
            return
        }
        if len(requestBody.Rate) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "rate")})
            return
        }
        if len(requestBody.PaireName) >= 6 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_OUT_OF_LENGTH, "6", "paire_name")})
            return
        }
        currencyExchangeRate := model.CurrencyExchangeRate{}
        currencyExchangeRate.BaseCurrencyCd, _ = strconv.Atoi(requestBody.BaseCurrencyCd)
        currencyExchangeRate.TargetCurrencyCd, _ = strconv.Atoi(requestBody.TargetCurrencyCd)

        currencyExchangeRate.Rate,_ = strconv.ParseFloat(requestBody.Rate, 64)




        result, err := la.currencyExchangeRateApplication.UpdateCurrencyExchangeRate(&ctx, &currencyExchangeRate)
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
            "currency_exchange_rate": result,
        })
    }
}

func (la currencyExchangeRateWebApi)DeleteCurrencyExchangeRate() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        requestBody := CurrencyExchangeRateEntity{}
        c.Bind(&requestBody)
        if len(requestBody.BaseCurrencyCd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "base_currency_cd")})
            return
        }
        if len(requestBody.TargetCurrencyCd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "target_currency_cd")})
            return
        }
        currencyExchangeRate := model.CurrencyExchangeRate{}
        currencyExchangeRate.BaseCurrencyCd, _ = strconv.Atoi(requestBody.BaseCurrencyCd)
        currencyExchangeRate.TargetCurrencyCd, _ = strconv.Atoi(requestBody.TargetCurrencyCd)

        currencyExchangeRate.Rate,_ = strconv.ParseFloat(requestBody.Rate, 64)




        err := la.currencyExchangeRateApplication.DeleteCurrencyExchangeRate(&ctx, currencyExchangeRate.BaseCurrencyCd, currencyExchangeRate.TargetCurrencyCd)
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

func (la currencyExchangeRateWebApi)GetCurrencyExchangeRateWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        baseCurrencyCd_bind := c.Query("base_currency_cd")
        targetCurrencyCd_bind := c.Query("target_currency_cd")
        if len(baseCurrencyCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "base_currency_cd")})
            return
        }
        if len(targetCurrencyCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "target_currency_cd")})
            return
        }
        baseCurrencyCd, _ := strconv.Atoi(baseCurrencyCd_bind)
        targetCurrencyCd, _ := strconv.Atoi(targetCurrencyCd_bind)
        results, err := la.currencyExchangeRateApplication.GetCurrencyExchangeRateWithKey(&ctx, baseCurrencyCd,targetCurrencyCd)
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
            "currency_exchange_rates": results,
        })
    }
}

func (la currencyExchangeRateWebApi)GetCurrencyExchangeRateWithPaireName() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        paireName_bind := c.Query("paire_name")
        if len(paireName_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "paire_name")})
            return
        }
        paireName := paireName_bind
        results, err := la.currencyExchangeRateApplication.GetCurrencyExchangeRateWithPaireName(&ctx, paireName)
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
            "currency_exchange_rates": results,
        })
    }
}

func (la currencyExchangeRateWebApi)GetCurrencyExchangeRateWithBase() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        baseCurrencyCd_bind := c.Query("base_currency_cd")
        if len(baseCurrencyCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "base_currency_cd")})
            return
        }
        baseCurrencyCd, _ := strconv.Atoi(baseCurrencyCd_bind)
        results, err := la.currencyExchangeRateApplication.GetCurrencyExchangeRateWithBase(&ctx, baseCurrencyCd)
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
            "currency_exchange_rates": results,
        })
    }
}
