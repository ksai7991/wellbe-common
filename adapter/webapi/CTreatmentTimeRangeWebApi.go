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


type CTreatmentTimeRangeWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type cTreatmentTimeRangeWebApi struct {
    cTreatmentTimeRangeApplication application.CTreatmentTimeRangeApplication
}

func NewCTreatmentTimeRangeWebApi(la application.CTreatmentTimeRangeApplication) CTreatmentTimeRangeWebApi {
    return &cTreatmentTimeRangeWebApi{
        cTreatmentTimeRangeApplication :la,
    }
}


func (la cTreatmentTimeRangeWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/c_treatment_time_ranges/key", la.GetCTreatmentTimeRangeWithKey())
    r.GET("/c_treatment_time_ranges/language_cd", la.GetCTreatmentTimeRangeWithLanguageCd())

    return r
}

type CTreatmentTimeRangeEntity struct {
    TreatmentTimeCd string `json:"treatment_time_cd"`
    LanguageCd string `json:"language_cd"`
    TreatmentTimeName string `json:"treatment_time_name"`
    MinTime string `json:"min_time"`
    MaxTime string `json:"max_time"`
    CreateDatetime string `json:"create_datetime"`
    CreateFunction string `json:"create_function"`
    UpdateDatetime string `json:"update_datetime"`
    UpdateFunction string `json:"update_function"`
}

func (la cTreatmentTimeRangeWebApi)GetCTreatmentTimeRangeWithKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        treatmentTimeCd_bind := c.Query("treatment_time_cd")
        languageCd_bind := c.Query("language_cd")
        if len(treatmentTimeCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "treatment_time_cd")})
            return
        }
        if len(languageCd_bind) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        treatmentTimeCd, _ := strconv.Atoi(treatmentTimeCd_bind)
        languageCd, _ := strconv.Atoi(languageCd_bind)
        results, err := la.cTreatmentTimeRangeApplication.GetCTreatmentTimeRangeWithKey(&ctx, treatmentTimeCd,languageCd)
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
            "c_treatment_time_ranges": results,
        })
    }
}

func (la cTreatmentTimeRangeWebApi)GetCTreatmentTimeRangeWithLanguageCd() gin.HandlerFunc {
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
        results, err := la.cTreatmentTimeRangeApplication.GetCTreatmentTimeRangeWithLanguageCd(&ctx, languageCd)
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
            "c_treatment_time_ranges": results,
        })
    }
}
