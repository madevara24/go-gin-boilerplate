package logger

import (
	"context"
)

var (
	instance Logger
)

func SetGlobalLogger(in Logger) {
	instance = in
}

func GetGlobalLogger() Logger {
	if instance == nil {
		return NewZapLogger()
	}

	return instance
}

func Debug(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Debug(ctx, message, fields...)
}

func Info(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Info(ctx, message, fields...)
}

func Warn(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Warn(ctx, message, fields...)
}

func Error(ctx context.Context, message string, err error, fields ...Field) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	fields = append(fields, Field{Key: "error", Val: msg})
	GetGlobalLogger().Error(ctx, message, fields...)
}

func Fatal(ctx context.Context, message string, err error, fields ...Field) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	fields = append(fields, Field{Key: "error", Val: msg})
	GetGlobalLogger().Fatal(ctx, message, fields...)
}

func Panic(ctx context.Context, message string, fields ...Field) {
	GetGlobalLogger().Panic(ctx, message, fields...)
}

func TDR(tdr LogTDRModel) {
	GetGlobalLogger().TDR(tdr)
}

func TDRV2(tdr LogTDRModel) {
	GetGlobalLogger().TDRV2(tdr)
}
