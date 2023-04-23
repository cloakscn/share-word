package ajax

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	StateOK          = ResultType{http.StatusOK, "操作成功"}
	StatusBadRequest = ResultType{http.StatusBadRequest, "错误的请求"}
	ErrorNotFound    = ResultType{http.StatusNotFound, "没有找到相关数据"}
)

type ResultType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	ResultType
	Data interface{} `json:"data"`
}

func Success(context *gin.Context, data interface{}) {
	context.JSONP(http.StatusOK, Message{
		ResultType: StateOK,
		Data:       data,
	})
}

func Error(context *gin.Context, resultType ResultType, data ...interface{}) {
	context.JSONP(http.StatusOK, Message{
		ResultType: resultType,
		Data:       data,
	})
}
