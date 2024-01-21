package logger

import (
	"context"
)

type Logger interface {
	Debug(ctx context.Context, message string, fields ...Field)
	Info(ctx context.Context, message string, fields ...Field)
	Warn(ctx context.Context, message string, fields ...Field)
	Error(ctx context.Context, message string, fields ...Field)
	Fatal(ctx context.Context, message string, fields ...Field)
	Panic(ctx context.Context, message string, fields ...Field)
	TDR(tdr LogTDRModel)
	TDRV2(tdr LogTDRModel)
	Close()
}

type Field struct {
	Key string
	Val interface{}
}

type ctxKeyLogger struct{}

var ctxKey = ctxKeyLogger{}

// Context is when you need bring something via context to another app, you will get the values (like distributed values)
type Context struct {
	JourneyID string `json:"_correlation_id"`
	// TODO: add more later
}

func InjectCtx(parent context.Context, ctx Context) context.Context {
	if parent == nil {
		return InjectCtx(context.Background(), ctx)
	}

	return context.WithValue(parent, ctxKey, ctx)
}

func ExtractCtx(ctx context.Context) Context {
	if ctx == nil {
		return Context{}
	}

	val, ok := ctx.Value(ctxKey).(Context)
	if !ok {
		return Context{}
	}

	return val
}
