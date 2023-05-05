package errors

import (
	"golang.org/x/xerrors"
	"log"
)

var (
	ErrorBadParam = xerrors.New("错误的请求内容")
)

func ErrorHandler(err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	panic(err)
}
