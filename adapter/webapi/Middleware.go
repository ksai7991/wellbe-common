package webapi

import (
	"bytes"
	"io/ioutil"
	"time"
	constants "wellbe-common/share/commonsettings/constants"
	log "wellbe-common/share/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordUaAndTime(c *gin.Context){
	logger := log.GetLogger()
	defer logger.Sync()
   oldTime := time.Now()

   var bodyBytes []byte
   if c.Request.Body != nil {
     bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
   }
   c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
   
   c.Next()
   if c.Request.URL.Path != "/" {
    logger.Info("incoming request",
        zap.String("path", c.Request.URL.Path),
        zap.String("ua", c.GetHeader("User-Agent")),
        zap.String("key", c.Request.Header.Get(constants.API_KEY_REUQEST_HEADER_NAME)),
        zap.String("clientIP", c.ClientIP()),
        zap.Int("status", c.Writer.Status()),
        zap.Duration("elapsed", time.Since(oldTime)),
    )
   }
}