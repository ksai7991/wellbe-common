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


type CShopContractPlanItemWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cShopContractPlanItemWebApi struct {
    cShopContractPlanItemApplication application.CShopContractPlanItemApplication
}

func NewCShopContractPlanItemWebApi(la application.CShopContractPlanItemApplication) CShopContractPlanItemWebApi {
    return &cShopContractPlanItemWebApi{
        cShopContractPlanItemApplication :la,
    }
}


func (la cShopContractPlanItemWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_shop_contract_plan_items/key", la.GetCShopContractPlanItemWithKey())
    r.GET("/c_shop_contract_plan_items/language_cd", la.GetCShopContractPlanItemWithLanguageCd())

    return r
}

type CShopContractPlanItemEntity struct {
    ShopContractPlanItemCd string `json:"shop_contract_plan_item_cd"`
    LanguageCd string `json:"language_cd"`
    ShopContractPlanName string `json:"shop_contract_plan_name"`
    Unit string `json:"unit"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cShopContractPlanItemWebApi)GetCShopContractPlanItemWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        shopContractPlanItemCd_bind := c.Query("shop_contract_plan_item_cd")
        languageCd_bind := c.Query("language_cd")
        if len(shopContractPlanItemCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "shop_contract_plan_item_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        shopContractPlanItemCd, _ := strconv.Atoi(shopContractPlanItemCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cShopContractPlanItemApplication.GetCShopContractPlanItemWithKey(&ctx, shopContractPlanItemCd,languageCd)
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
            "c_shop_contract_plan_items": results,
        })
    }
}

func (la cShopContractPlanItemWebApi)GetCShopContractPlanItemWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cShopContractPlanItemApplication.GetCShopContractPlanItemWithLanguageCd(&ctx, languageCd)
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
            "c_shop_contract_plan_items": results,
        })
    }
}
