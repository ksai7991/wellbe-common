package log

import (
    "testing"
    "github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	
)

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	defer logger.Sync()
	assert.NotNil(t, logger)
	logger.Info("test", zap.String("url", "http://hogehoge"))
}