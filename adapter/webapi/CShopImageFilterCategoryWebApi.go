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


type CShopImageFilterCategoryWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cShopImageFilterCategoryWebApi struct {
    cShopImageFilterCategoryApplication application.CShopImageFilterCategoryApplication
}

func NewCShopImageFilterCategoryWebApi(la application.CShopImageFilterCategoryApplication) CShopImageFilterCategoryWebApi {
    return &cShopImageFilterCategoryWebApi{
        cShopImageFilterCategoryApplication :la,
    }
}


func (la cShopImageFilterCategoryWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_shop_image_filter_categorys/key", la.GetCShopImageFilterCategoryWithKey())
    r.GET("/c_shop_image_filter_categorys/language_cd", la.GetCShopImageFilterCategoryWithLanguageCd())

    return r
}

type CShopImageFilterCategoryEntity struct {
    ShopImageFilterCategoryCd string `json:"shop_image_filter_category_cd"`
    LanguageCd string `json:"language_cd"`
    ShopImageFilterCategoryName string `json:"shop_image_filter_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cShopImageFilterCategoryWebApi)GetCShopImageFilterCategoryWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        shopImageFilterCategoryCd_bind := c.Query("shop_image_filter_category_cd")
        languageCd_bind := c.Query("language_cd")
        if len(shopImageFilterCategoryCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "shop_image_filter_category_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        shopImageFilterCategoryCd, _ := strconv.Atoi(shopImageFilterCategoryCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cShopImageFilterCategoryApplication.GetCShopImageFilterCategoryWithKey(&ctx, shopImageFilterCategoryCd,languageCd)
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
            "c_shop_image_filter_categorys": results,
        })
    }
}

func (la cShopImageFilterCategoryWebApi)GetCShopImageFilterCategoryWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cShopImageFilterCategoryApplication.GetCShopImageFilterCategoryWithLanguageCd(&ctx, languageCd)
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
            "c_shop_image_filter_categorys": results,
        })
    }
}
