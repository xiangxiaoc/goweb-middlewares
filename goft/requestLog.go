package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/xiangxiaoc/goweb-middlewares/internal"
	"log/slog"
)

type RequestLog struct {
}

func NewRequestLog() *RequestLog {
	return &RequestLog{}
}

func (r *RequestLog) OnRequest(ctx *gin.Context) error {
	slog.Info(fmt.Sprintf(
		"Request: {uri: %v, params: %v}", ctx.Request.RequestURI, ctx.Params),
	)
	return nil
}

func (r *RequestLog) OnResponse(result any) (any, error) {
	if h, ok := result.(gin.H); ok {
		slog.Info(fmt.Sprintf(
			"Response: {code: %v msg: %v}", h["code"], h["msg"]),
		)
	}
	return result, nil
}
