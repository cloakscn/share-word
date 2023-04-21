package routers

import (
	"fmt"
	"github.com/cloakscn/share-word/utils/errors"
	"github.com/cloakscn/share-word/utils/https"
	"io"
	"net/http"
)

func Router(a *https.App) {
	a.Handle("/create", Create)
}

func Create(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	for key, values := range queryParams {
		_, _ = fmt.Fprintf(w, "%s = %s\n", key, values)
	}

	_, err := io.WriteString(w, "hello world.")
	if err != nil {
		errors.ErrorHandler(err)
	}
}
