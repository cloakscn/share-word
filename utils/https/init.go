package https

import (
	"log"
	"net/http"
)

type Config struct {
	Port string
}

func (a *App) Server(cfg *Config) {
	log.Println("初始化 http 服务。")
	log.Println("http://localhost:8080")
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalln("初始化 http 服务失败。")
	}
}
