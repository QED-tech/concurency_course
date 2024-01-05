package logger

import (
	"fmt"
	"go.uber.org/zap"
)

//go:generate go run go.uber.org/mock/mockgen -package logger -destination mock.go -source log.go Logger
type Logger interface {
	Infof(template string, args ...any)
	Errorf(template string, args ...any)
	Warnf(template string, args ...any)
}

type Log struct {
	base *zap.SugaredLogger
}

func (l Log) Errorf(template string, args ...any) {
	l.base.Errorf(template, args)
}

func (l Log) Warnf(template string, args ...any) {
	l.base.Warnf(template, args)
}

func (l Log) Infof(template string, args ...any) {
	l.base.Infof(template, args)
}

func NewLogger() (Logger, error) {
	zapConfig := zap.NewProductionConfig()

	l, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("[logger] failed to build logger: %w", err)
	}

	return &Log{base: l.Sugar()}, nil
}
