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


type CPayoutItemCategoryWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cPayoutItemCategoryWebApi struct {
    cPayoutItemCategoryApplication application.CPayoutItemCategoryApplication
}

func NewCPayoutItemCategoryWebApi(la application.CPayoutItemCategoryApplication) CPayoutItemCategoryWebApi {
    return &cPayoutItemCategoryWebApi{
        cPayoutItemCategoryApplication :la,
    }
}


func (la cPayoutItemCategoryWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_payout_item_categorys/key", la.GetCPayoutItemCategoryWithKey())
    r.GET("/c_payout_item_categorys/language_cd", la.GetCPayoutItemCategoryWithLanguageCd())

    return r
}

type CPayoutItemCategoryEntity struct {
    PayoutItemCategoryCd string `json:"payout_item_category_cd"`
    LanguageCd string `json:"language_cd"`
    PayoutItemCategoryName string `json:"payout_item_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cPayoutItemCategoryWebApi)GetCPayoutItemCategoryWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        payoutItemCategoryCd_bind := c.Query("payout_item_category_cd")
        if len(payoutItemCategoryCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "payout_item_category_cd")})
            return
        }
        payoutItemCategoryCd, _ := strconv.Atoi(payoutItemCategoryCd_bind)
        results, err := la.cPayoutItemCategoryApplication.GetCPayoutItemCategoryWithKey(&ctx, payoutItemCategoryCd)
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
            "c_payout_item_categorys": results,
        })
    }
}

func (la cPayoutItemCategoryWebApi)GetCPayoutItemCategoryWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cPayoutItemCategoryApplication.GetCPayoutItemCategoryWithLanguageCd(&ctx, languageCd)
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
            "c_payout_item_categorys": results,
        })
    }
}
