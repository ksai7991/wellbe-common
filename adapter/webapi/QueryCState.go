package webapi

import (
	"net/http"

	"fmt"
	"wellbe-common/adapter/repository/query"
	"wellbe-common/domain/repository"
	constants "wellbe-common/share/commonsettings/constants"
	messages "wellbe-common/share/messages"

	"github.com/gin-gonic/gin"
)


type QueryCStateWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type queryCStateWebApi struct {
    query *query.Query
    repository repository.Repository
    transaction repository.Transaction
}

func NewQueryCStateWebApi(q *query.Query, r repository.Repository, tr repository.Transaction) QueryCStateWebApi {
    return &queryCStateWebApi{
        query :q,
        repository: r,
        transaction :tr,
    }
}


func (la queryCStateWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/query/c_state", la.GetCState())

    return r
}

func (la queryCStateWebApi)GetCState() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        defer la.transaction.Rollback(&ctx)
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT && key != constants.API_KEY_API {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        state_name := c.Query("state_name")
        if len(state_name) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "state_name")})
            return
        }

        result, err := la.query.QueryCState(&ctx, state_name)
        
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
            "c_states": result,
        })
        la.transaction.Commit(&ctx)
    }
}
