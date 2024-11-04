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


type CCheckoutMethodWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cCheckoutMethodWebApi struct {
    cCheckoutMethodApplication application.CCheckoutMethodApplication
}

func NewCCheckoutMethodWebApi(la application.CCheckoutMethodApplication) CCheckoutMethodWebApi {
    return &cCheckoutMethodWebApi{
        cCheckoutMethodApplication :la,
    }
}


func (la cCheckoutMethodWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_checkout_methods/key", la.GetCCheckoutMethodWithKey())
    r.GET("/c_checkout_methods/language_cd", la.GetCCheckoutMethodWithLanguageCd())

    return r
}

type CCheckoutMethodEntity struct {
    CheckoutMethodCd string `json:"checkout_method_cd"`
    LanguageCd string `json:"language_cd"`
    CheckoutMethodName string `json:"checkout_method_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cCheckoutMethodWebApi)GetCCheckoutMethodWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        checkoutMethodCd_bind := c.Query("checkout_method_cd")
        languageCd_bind := c.Query("language_cd")
        if len(checkoutMethodCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "checkout_method_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        checkoutMethodCd, _ := strconv.Atoi(checkoutMethodCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cCheckoutMethodApplication.GetCCheckoutMethodWithKey(&ctx, checkoutMethodCd,languageCd)
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
            "c_checkout_methods": results,
        })
    }
}

func (la cCheckoutMethodWebApi)GetCCheckoutMethodWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cCheckoutMethodApplication.GetCCheckoutMethodWithLanguageCd(&ctx, languageCd)
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
            "c_checkout_methods": results,
        })
    }
}
