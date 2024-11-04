package webapi

import (
    "net/http"
    
    application "wellbe-common/application"

    constants "wellbe-common/share/commonsettings/constants"
    messages "wellbe-common/share/messages"
    model "wellbe-common/domain/model"
    "fmt"
    "strconv"
    "regexp"
    "github.com/gin-gonic/gin"
)


type DefaultFeeMasterWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type defaultFeeMasterWebApi struct {
    defaultFeeMasterApplication application.DefaultFeeMasterApplication
}

func NewDefaultFeeMasterWebApi(la application.DefaultFeeMasterApplication) DefaultFeeMasterWebApi {
    return &defaultFeeMasterWebApi{
        defaultFeeMasterApplication :la,
    }
}


func (la defaultFeeMasterWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.POST("/cud/default_fee_master/create", la.CreateDefaultFeeMaster())
    r.POST("/cud/default_fee_master/update", la.UpdateDefaultFeeMaster())
    r.POST("/cud/default_fee_master/delete", la.DeleteDefaultFeeMaster())

    return r
}

type DefaultFeeMasterEntity struct {
    Id string `json:"id"`
    FeeRate string `json:"fee_rate"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la defaultFeeMasterWebApi)CreateDefaultFeeMaster() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        requestBody := DefaultFeeMasterEntity{}
        c.Bind(&requestBody)
        if len(requestBody.Id) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "id")})
            return
        }
        if len(requestBody.FeeRate) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "fee_rate")})
            return
        }
        re := regexp.MustCompile(`^[+-]?\d+(?:\.\d+)?$`)
        if len(requestBody.FeeRate) > 0 && re.MatchString(requestBody.FeeRate) == false {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT, "fee_rate", `^[+-]?\d+(?:\.\d+)?$`)})
            return
        }
        defaultFeeMaster := model.DefaultFeeMaster{}

        defaultFeeMaster.FeeRate,_ = strconv.ParseFloat(requestBody.FeeRate, 64)




        result, err := la.defaultFeeMasterApplication.CreateDefaultFeeMaster(&ctx, &defaultFeeMaster)
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
            "default_fee_master": result,
        })
    }
}

func (la defaultFeeMasterWebApi)UpdateDefaultFeeMaster() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        requestBody := DefaultFeeMasterEntity{}
        c.Bind(&requestBody)
        if len(requestBody.Id) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "id")})
            return
        }
        if len(requestBody.FeeRate) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "fee_rate")})
            return
        }
        re := regexp.MustCompile(`^[+-]?\d+(?:\.\d+)?$`)
        if len(requestBody.FeeRate) > 0 && re.MatchString(requestBody.FeeRate) == false {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT, "fee_rate", `^[+-]?\d+(?:\.\d+)?$`)})
            return
        }
        defaultFeeMaster := model.DefaultFeeMaster{}

        defaultFeeMaster.FeeRate,_ = strconv.ParseFloat(requestBody.FeeRate, 64)




        result, err := la.defaultFeeMasterApplication.UpdateDefaultFeeMaster(&ctx, &defaultFeeMaster)
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
            "default_fee_master": result,
        })
    }
}

func (la defaultFeeMasterWebApi)DeleteDefaultFeeMaster() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        requestBody := DefaultFeeMasterEntity{}
        c.Bind(&requestBody)
        if len(requestBody.Id) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "id")})
            return
        }
        defaultFeeMaster := model.DefaultFeeMaster{}

        defaultFeeMaster.FeeRate,_ = strconv.ParseFloat(requestBody.FeeRate, 64)




        err := la.defaultFeeMasterApplication.DeleteDefaultFeeMaster(&ctx, defaultFeeMaster.Id)
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
