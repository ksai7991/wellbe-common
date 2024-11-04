package webapi

import (
	"net/http"
	"regexp"

	"fmt"
	"wellbe-common/adapter/repository/query"
	"wellbe-common/domain/repository"
	constants "wellbe-common/share/commonsettings/constants"
	messages "wellbe-common/share/messages"

	"github.com/gin-gonic/gin"
)


type QueryCAreaWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type queryCAreaWebApi struct {
    query *query.Query
    repository repository.Repository
    transaction repository.Transaction
}

func NewQueryCAreaWebApi(q *query.Query, r repository.Repository, tr repository.Transaction) QueryCAreaWebApi {
    return &queryCAreaWebApi{
        query :q,
        repository: r,
        transaction :tr,
    }
}


func (la queryCAreaWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/query/c_area", la.GetCArea())

    return r
}

func (la queryCAreaWebApi)GetCArea() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        defer la.transaction.Rollback(&ctx)
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT && key != constants.API_KEY_API {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        language_cd := c.Query("language_cd")
        if len(language_cd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }

        area_cd := c.Query("area_cd")
        if len(area_cd) > 0 &&  !regexp.MustCompile(`^[+-]?\d+$`).MatchString(area_cd) {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT, "area_cd", `^[+-]?\d+$`)})
            return
        }
        state_cd := c.Query("state_cd")
        if len(state_cd) > 0 &&  !regexp.MustCompile(`^[+-]?\d+$`).MatchString(state_cd) {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT, "state_cd", `^[+-]?\d+$`)})
            return
        }
        country_cd := c.Query("country_cd")
        if len(country_cd) > 0 &&  !regexp.MustCompile(`^[+-]?\d+$`).MatchString(country_cd) {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT, "country_cd", `^[+-]?\d+$`)})
            return
        }

        result, err := la.query.QueryArea(&ctx, language_cd, country_cd, state_cd, area_cd)
        
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
            "c_area": result,
        })
        la.transaction.Commit(&ctx)
    }
}
