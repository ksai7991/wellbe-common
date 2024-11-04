package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger() (*zap.Logger) {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	
	myConfig := zap.Config{
		Level: level,
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err :=  myConfig.Build()
	if err != nil {
        panic(err)
    }

	return logger
}