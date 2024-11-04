package webapi

import (
	"net/http"
	"regexp"

	"fmt"
	"wellbe-common/adapter/repository/query"
	"wellbe-common/domain/repository"
	constants "wellbe-common/share/commonsettings/constants"
	"wellbe-common/share/datetime"
	messages "wellbe-common/share/messages"

	"github.com/gin-gonic/gin"
)


type QueryCAgeRangeWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type queryCAgeRangeWebApi struct {
    query *query.Query
    repository repository.Repository
    transaction repository.Transaction
}

func NewQueryCAgeRangeWebApi(q *query.Query, r repository.Repository, tr repository.Transaction) QueryCAgeRangeWebApi {
    return &queryCAgeRangeWebApi{
        query :q,
        repository: r,
        transaction :tr,
    }
}


func (la queryCAgeRangeWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/query/current_age_gender", la.GetCurrentAgeGender())

    return r
}

func (la queryCAgeRangeWebApi)GetCurrentAgeGender() gin.HandlerFunc {
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

        birth_date := c.Query("birth_date")
        if len(birth_date) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "birth_date")})
            return
        }
        if len(birth_date) > 0 &&  !regexp.MustCompile(`^[0-9]{4}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])$`).MatchString(birth_date) {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_VALUE_INVALID_FORMAT, "birth_date", `^[0-9]{4}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])$`)})
            return
        }

        age, errDate := datetime.CalculateAge(birth_date)
        if errDate != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                constants.WEBAPI_RESPONSE_KEYWORD_STATUS:constants.LOGIC_ERROR_CODE_SEVERERROR,
                constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: messages.MESSAGE_EN_SERVER_ERROR,
            })
        }
        result, err := la.query.QueryAgeRange(&ctx, language_cd, age)
        
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
            "c_age_range": result,
        })
        la.transaction.Commit(&ctx)
    }
}
