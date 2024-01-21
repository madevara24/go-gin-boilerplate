package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"runtime"
	"strings"
	"time"

	"go-gin-boilerplate/internal/pkg"
	"go-gin-boilerplate/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

const (
	beginTime = "begin"
)

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}

func TDRLog(additionalData ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, isDisallowLog := disallowLogOnPathMap[c.Request.URL.Path]; isDisallowLog {
			return
		}

		reqID := c.Request.Header.Get(pkg.HeaderXRequestID)
		if len(reqID) == 0 {
			reqID = uuid.New().String()
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, pkg.ThreadIDKey, reqID)
		ctx = context.WithValue(ctx, beginTime, time.Now())
		c.Request = c.Request.WithContext(ctx)

		c.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Next()

		bw, ok := c.Writer.(*bodyWriter)
		if !ok {
			logger.Fatal(ctx, "the writer was override, can not read bodyCache", nil)
			return
		}

		errMsg, _ := c.Get(pkg.ErrorMessageKey)
		threadID := ctx.Value(pkg.ThreadIDKey)

		var (
			reqRaw, _ = c.GetRawData()
			reqBody   interface{}
			respBody  interface{}
		)

		if len(reqRaw) > 0 && strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			reqBody = json.RawMessage(reqRaw)
		}

		if len(bw.bodyCache.Bytes()) > 0 && strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") {
			respBody = json.RawMessage(bw.bodyCache.Bytes())
		}

		tdr := logger.LogTDRModel{
			XTime:         time.Now().Format(time.RFC3339),
			AppName:       os.Getenv("APP_NAME"),
			Runtime:       runtime.Version(),
			Env:           os.Getenv("ENV"),
			AppVersion:    os.Getenv("APP_VERSION"),
			CorrelationID: cast.ToString(threadID),
			Path:          c.Request.URL.String(),
			Method:        c.Request.Method,
			IP:            "localhost", // TODO ?
			Port:          os.Getenv("PORT"),
			SrcIP:         c.ClientIP(),
			ResponseCode:  c.Writer.Status(),
			Header:        getRequestHeaders(c),
			Request:       reqBody,
			Response:      respBody,
			Error:         cast.ToString(errMsg),
		}

		if len(additionalData) > 0 {
			tdr.AdditionalData = additionalData[0]
		}

		beginTime, ok := ctx.Value(beginTime).(time.Time)
		if ok {
			tdr.RespTime = time.Since(beginTime).Milliseconds()
		}

		logger.TDR(tdr)
	}
}

func getRequestHeaders(c *gin.Context) map[string]interface{} {
	headers := map[string]interface{}{}

	for key, val := range c.Request.Header {
		if key == "Authorization" && val[0] != "" {
			headers[key] = "#####"
			continue
		}

		headers[key] = val[0]
	}

	return headers
}

var disallowLogOnPathMap = map[string]struct{}{
	"/metrics": {},
	"/":        {},
	"/health":  {},
}
