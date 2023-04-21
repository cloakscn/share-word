package https

import (
	"net/http"
)

type HandleFunc func(pattern string, handlerFunc http.HandlerFunc)
type Interceptor func(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request)

type App struct {
	interceptor []Interceptor
}

func NewApp() *App {
	return &App{}
}

func (a *App) RegisterInterceptor(fc Interceptor) {
	a.interceptor = append(a.interceptor, fc)
}

func (a *App) Handle(pattern string, handlerFunc http.HandlerFunc) {
	http.Handle(pattern, a.safeHandler(handlerFunc))
}

func (a *App) safeHandler(fc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 执行拦截器
		for _, interceptor := range a.interceptor {
			w, r = interceptor(w, r)
		}
		// 执行函数
		fc(w, r)
	}
}
