package errors

import (
	"log"
)

func ErrorHandler(err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	panic(err)
}
