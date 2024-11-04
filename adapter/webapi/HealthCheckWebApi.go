package webapi

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


type HealthCheckWebApi interface {
    CreateAccessPoint(r *gin.Engine)*gin.Engine
}

type healthCheckWebApi struct {
}

func NewHealthCheckWebApi() HealthCheckWebApi {
    return &healthCheckWebApi{
    }
}


func (la healthCheckWebApi) CreateAccessPoint(r *gin.Engine)*gin.Engine{
    r.GET("/", la.HealthCheck())

    return r
}

func (la healthCheckWebApi)HealthCheck() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.JSON(http.StatusOK, gin.H{
        })
    }
}
