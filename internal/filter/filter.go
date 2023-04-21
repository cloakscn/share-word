package filter

import (
	"github.com/cloakscn/share-word/utils/https"
	"log"
	"net/http"
)

func Filter(a *https.App) {
	a.RegisterInterceptor(logHttpReq)
}

func logHttpReq(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	return w, r
}
