package ajax

import (
	"net/http"

	"github.com/cloakscn/share-word/server/utils/errors"
	"github.com/gin-gonic/gin"
)

var (
	StatusOK                  = ResultType{http.StatusOK, "操作成功"}
	StatusBadRequest          = ResultType{http.StatusBadRequest, "错误的请求"}
	StatusNotFound            = ResultType{http.StatusNotFound, "没有找到相关数据"}
	StatusInternalServerError = ResultType{http.StatusInternalServerError, "服务器错误"}
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
		ResultType: StatusOK,
		Data:       data,
	})
}

func Error(context *gin.Context, resultType ResultType, err error) {
	errors.ErrorHandler(err)
	context.JSONP(http.StatusOK, Message{
		ResultType: resultType,
		Data:       err,
	})
}
