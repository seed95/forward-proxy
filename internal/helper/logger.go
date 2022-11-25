package helper

import (
	"fmt"
	"github.com/seed95/forward-proxy/pkg/log/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level int) (zap.Logger, error) {
	var cores []zapcore.Core

	stdCore, err := zap.NewStandardCore(true, zap.Level(level))
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create std log instance, error: %v\n", err)
	}
	cores = append(cores, stdCore)
	loggerCores := zap.NewZapLoggerWithCores(cores...)

	return loggerCores, nil
}
