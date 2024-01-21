package logger

import (
	"context"
	"fmt"
	"go-gin-boilerplate/internal/pkg"
	"os"
	"time"

	"go.elastic.co/apm/module/apmzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	zapLogger *zap.Logger
}

var _ Logger = (*logger)(nil)

func NewZapLogger() Logger {
	fmt.Println("Try newLogger ...")

	config := zap.NewDevelopmentConfig()
	if os.Getenv("ENV") == "prod" || os.Getenv("ENV") == "production" {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "xtime",
		MessageKey:     "msg",
		LevelKey:       "level",
		EncodeDuration: millisDurationEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	defaultLogger := &logger{
		zapLogger: zap.New(core),
	}

	return defaultLogger
}

func (d *logger) Close() {
	d.zapLogger.Sync()
}

func (d *logger) Debug(ctx context.Context, message string, fields ...Field) {
	zapLogs := []zap.Field{
		zap.String("level", "debug"),
		zap.String("type", "SYS"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, fields...)...)
	d.zapLogger.Debug(message, zapLogs...)
}

func (d *logger) Info(ctx context.Context, message string, fields ...Field) {
	zapLogs := []zap.Field{
		zap.String("level", "info"),
		zap.String("type", "SYS"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, fields...)...)
	d.zapLogger.Info(message, zapLogs...)
}

func (d *logger) Warn(ctx context.Context, message string, fields ...Field) {
	zapLogs := []zap.Field{
		zap.String("level", "warn"),
		zap.String("type", "SYS"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, fields...)...)
	d.zapLogger.Warn(message, zapLogs...)
}

func (d *logger) Error(ctx context.Context, message string, fields ...Field) {
	zapLogs := []zap.Field{
		zap.String("level", "error"),
		zap.String("type", "SYS"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, fields...)...)
	d.zapLogger.Error(message, zapLogs...)
}

func (d *logger) Fatal(ctx context.Context, message string, fields ...Field) {
	zapLogs := []zap.Field{
		zap.String("level", "fatal"),
		zap.String("type", "SYS"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, fields...)...)
	d.zapLogger.Fatal(message, zapLogs...)
}

func (d *logger) Panic(ctx context.Context, message string, fields ...Field) {
	zapLogs := []zap.Field{
		zap.String("level", "panic"),
		zap.String("type", "SYS"),
	}

	zapLogs = append(zapLogs, formatLogs(ctx, fields...)...)
	d.zapLogger.Panic(message, zapLogs...)
}

func (d *logger) newFieldTDR(tdr LogTDRModel) []zapcore.Field {
	fields := make([]zap.Field, 0)
	fields = append(fields, zap.String("level", "info"))
	fields = append(fields, zap.String("type", "TDR"))

	fields = append(fields, zap.String("app", tdr.AppName))
	fields = append(fields, zap.String("env", tdr.Env))
	fields = append(fields, zap.String("runtime", tdr.Runtime))
	fields = append(fields, zap.String("app_version", tdr.AppVersion))
	fields = append(fields, zap.String("correlation_id", tdr.CorrelationID))

	fields = append(fields, zap.Any("path", tdr.Path))
	fields = append(fields, zap.String("method", tdr.Method))
	fields = append(fields, zap.Any("ip", tdr.IP))
	fields = append(fields, zap.String("port", tdr.Port))
	fields = append(fields, zap.String("src_ip", tdr.SrcIP))
	fields = append(fields, zap.Int64("rt", tdr.RespTime))
	fields = append(fields, zap.Int("rc", tdr.ResponseCode))

	if tdr.Error != "" {
		fields = append(fields, zap.String("error", tdr.Error))
	}

	fields = append(fields, zap.Any("additional_data", tdr.AdditionalData))

	return fields
}

func (d *logger) TDR(tdr LogTDRModel) {
	fields := d.newFieldTDR(tdr)
	fields = append(fields, zap.Any("header", tdr.Header))
	fields = append(fields, zap.Any("req", tdr.Request))
	fields = append(fields, zap.Any("resp", tdr.Response))

	d.zapLogger.Info("TDR", fields...)
}

func (d *logger) TDRV2(tdr LogTDRModel) {
	fields := d.newFieldTDR(tdr)
	fields = append(fields, zap.Any("trace_id", tdr.TraceID))

	d.zapLogger.Info("TDR", fields...)
}

func formatLogs(ctx context.Context, fields ...Field) []zap.Field {
	zapLogs := make([]zap.Field, 0)

	corrID, _ := ctx.Value(pkg.ThreadIDKey).(string)

	ctxVal := ExtractCtx(ctx)
	zapLogs = append(zapLogs, zap.String("_journey_id", ctxVal.JourneyID))
	zapLogs = append(zapLogs, zap.String("correlationID", corrID))
	zapLogs = append(zapLogs, apmzap.TraceContext(ctx)...)

	for _, d := range fields {
		zapLogs = append(zapLogs, zap.Any(d.Key, d.Val))
	}

	return zapLogs
}

func millisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Nanoseconds() / 1000000)
}
