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


type CShopMaintenanceLabelWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cShopMaintenanceLabelWebApi struct {
    cShopMaintenanceLabelApplication application.CShopMaintenanceLabelApplication
}

func NewCShopMaintenanceLabelWebApi(la application.CShopMaintenanceLabelApplication) CShopMaintenanceLabelWebApi {
    return &cShopMaintenanceLabelWebApi{
        cShopMaintenanceLabelApplication :la,
    }
}


func (la cShopMaintenanceLabelWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_shop_maintenance_labels/key", la.GetCShopMaintenanceLabelWithKey())
    r.GET("/c_shop_maintenance_labels/language_cd", la.GetCShopMaintenanceLabelWithLanguageCd())

    return r
}

type CShopMaintenanceLabelEntity struct {
    ShopMaintenanceLabelCd string `json:"shop_maintenance_label_cd"`
    LanguageCd string `json:"language_cd"`
    ShopMaintenanceLabelName string `json:"shop_maintenance_label_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cShopMaintenanceLabelWebApi)GetCShopMaintenanceLabelWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        shopMaintenanceLabelCd_bind := c.Query("shop_maintenance_label_cd")
        languageCd_bind := c.Query("language_cd")
        if len(shopMaintenanceLabelCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "shop_maintenance_label_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        shopMaintenanceLabelCd, _ := strconv.Atoi(shopMaintenanceLabelCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cShopMaintenanceLabelApplication.GetCShopMaintenanceLabelWithKey(&ctx, shopMaintenanceLabelCd,languageCd)
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
            "c_shop_maintenance_labels": results,
        })
    }
}

func (la cShopMaintenanceLabelWebApi)GetCShopMaintenanceLabelWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cShopMaintenanceLabelApplication.GetCShopMaintenanceLabelWithLanguageCd(&ctx, languageCd)
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
            "c_shop_maintenance_labels": results,
        })
    }
}
