package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"shop-payment/shop-payment/sutil"
	"scm.tutorabc.com/jimzhong/grpc-idl-templete/base/include/base-idl"
	"scm.tutorabc.com/jimzhong/grpc-idl-templete/constants"
	"scm.tutorabc.com/tgo-framework/go-trace"
)

type Gin struct {
	C *gin.Context
}

// ResponseTemplete 少用少用...
type ResponseTemplete struct {
	ApiMeta base_idl.ResponseApiMeta `json:"api_meta"`
	Data    interface{}              `json:"data"`
}

// GenResponseApiMeta 從IDL gen出 response header
func GenResponseApiMeta(ctx context.Context, errCode constants.ErrNo) *base_idl.ResponseApiMeta {
	code := int32(errCode)
	return &base_idl.ResponseApiMeta{
		Code:    sutil.Int32Ptr(code),
		Success: sutil.BoolPtr(code >= 0),
		Msg:     sutil.StringPtr(errCode.String()),
		Trace:   trace.FromContext(ctx),
	}
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.C.JSON(httpCode, data)
	return
}

func (g *Gin) ResponseErr(httpCode int, errCode constants.ErrNo, err error) {
	msg := errCode.String()
	if err != nil {
		msg = err.Error()
	}
	data := base_idl.DefaultResponse{ApiMeta: GenResponseApiMeta(g.C.Request.Context(), errCode)}
	data.GetApiMeta().Msg = sutil.StringPtr(msg)
	g.C.JSON(httpCode, data)
	return
}

// 隨意的Response, 盡量走IDL規範過的
// Response setting gin.JSON
func (g *Gin) ResponseCasual(httpCode int, errCode int, data interface{}) {
	g.C.JSON(httpCode, ResponseTemplete{
		Data: data,
		ApiMeta: base_idl.ResponseApiMeta{
			Code:    sutil.Int32Ptr(int32(errCode)),
			Success: sutil.BoolPtr(errCode >= 0),
			Trace:   trace.FromContext(g.C.Request.Context()),
		},
	})
	return
}

// 隨意的Response, 盡量走IDL規範過的
// Response setting gin.JSON
func (g *Gin) ResponseErrCasual(httpCode, errCode int, err error) {
	msg := "unknown"
	if err != nil {
		msg = err.Error()
	}
	g.C.JSON(httpCode, ResponseTemplete{
		ApiMeta: base_idl.ResponseApiMeta{
			Code:    sutil.Int32Ptr(int32(errCode)),
			Msg:     sutil.StringPtr(msg),
			Success: sutil.BoolPtr(false),
			Trace:   trace.FromContext(g.C.Request.Context()),
		},
	})
	return
}

const (
	SUCCESS  = 0
	ERROR    = -1
	ParamBad = -2
)

var MsgFlags = map[int]string{
	SUCCESS:  "ok",
	ERROR:    "server error",
	ParamBad: "param invalid",
}

// GetMsg get error information based on Code
func getMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
