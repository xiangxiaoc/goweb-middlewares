package goft

import (
	"github.com/gin-gonic/gin"
	"log"
)

type RequestLog struct {
}

func NewRequestLog() *RequestLog {
	return &RequestLog{}
}

func (r *RequestLog) OnRequest(ctx *gin.Context) error {
	log.Printf("Request: {uri: %v, params: %v}", ctx.Request.RequestURI, ctx.Params)
	return nil
}

func (r *RequestLog) OnResponse(result any) (any, error) {
	if h, ok := result.(gin.H); ok {
		log.Printf("Response: {code: %v msg: %v}", h["code"], h["msg"])
	}
	return result, nil
}
