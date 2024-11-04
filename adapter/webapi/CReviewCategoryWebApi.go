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


type CReviewCategoryWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cReviewCategoryWebApi struct {
    cReviewCategoryApplication application.CReviewCategoryApplication
}

func NewCReviewCategoryWebApi(la application.CReviewCategoryApplication) CReviewCategoryWebApi {
    return &cReviewCategoryWebApi{
        cReviewCategoryApplication :la,
    }
}


func (la cReviewCategoryWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_review_categorys/key", la.GetCReviewCategoryWithKey())
    r.GET("/c_review_categorys/language_cd", la.GetCReviewCategoryWithLanguageCd())

    return r
}

type CReviewCategoryEntity struct {
    ReviewCategoryCd string `json:"review_category_cd"`
    LanguageCd string `json:"language_cd"`
    ReviewCategoryName string `json:"review_category_name"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cReviewCategoryWebApi)GetCReviewCategoryWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        reviewCategoryCd_bind := c.Query("review_category_cd")
        languageCd_bind := c.Query("language_cd")
        if len(reviewCategoryCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "review_category_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        reviewCategoryCd, _ := strconv.Atoi(reviewCategoryCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cReviewCategoryApplication.GetCReviewCategoryWithKey(&ctx, reviewCategoryCd,languageCd)
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
            "c_review_categorys": results,
        })
    }
}

func (la cReviewCategoryWebApi)GetCReviewCategoryWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cReviewCategoryApplication.GetCReviewCategoryWithLanguageCd(&ctx, languageCd)
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
            "c_review_categorys": results,
        })
    }
}
