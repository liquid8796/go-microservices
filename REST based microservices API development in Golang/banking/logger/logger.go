package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	var err error

	log, err = zap.NewProduction(zap.AddCallerSkip(4))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	fmt.Println(fields)
	log.Info(message, fields...)
}
