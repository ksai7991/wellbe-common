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


type QueryCurrencyForPaymentWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type queryCurrencyForPaymentWebApi struct {
    query *query.Query
    repository repository.Repository
    transaction repository.Transaction
}

func NewQueryCurrencyForPaymentWebApi(q *query.Query, r repository.Repository, tr repository.Transaction) QueryCurrencyForPaymentWebApi {
    return &queryCurrencyForPaymentWebApi{
        query :q,
        repository: r,
        transaction :tr,
    }
}


func (la queryCurrencyForPaymentWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/query/currency_for_payments", la.GetCurrencyForPayment())

    return r
}

func (la queryCurrencyForPaymentWebApi)GetCurrencyForPayment() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        defer la.transaction.Rollback(&ctx)
        key := c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)
        if key != constants.API_KEY_CLIENT {
            c.JSON(http.StatusUnauthorized, gin.H{})
            return
        }

        language_cd := c.Query("language_cd")
        if len(language_cd) == 0 {
            c.JSON(http.StatusBadRequest, gin.H{constants.WEBAPI_RESPONSE_KEYWORD_MESSAGE: fmt.Sprintf(messages.MESSAGE_EN_REQUEST_ITEM_MANDATORY, "language_cd")})
            return
        }
        results, err := la.query.QueryCurrencyForPayment(&ctx, language_cd)
        
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
            "c_currencys": results,
        })
        la.transaction.Commit(&ctx)
    }
}
