package cmd

import (
	"github.com/cloakscn/share-word/internal/filter"
	"github.com/cloakscn/share-word/internal/routers"
	"github.com/cloakscn/share-word/utils/https"
	"github.com/cloakscn/share-word/utils/redis"
	"sync"
)

var (
	wg sync.WaitGroup
)

type Config struct {
	Redis *redis.Config
	Http  *https.Config
}

func init() {
	wg = sync.WaitGroup{}
}

func Start(cfg *Config) {
	// 初始化 Redis 服务
	redis.Init(cfg.Redis)

	wg.Add(1)
	go startHttpServer(cfg, &wg)
	wg.Wait()
}

func startHttpServer(cfg *Config, wg *sync.WaitGroup) {
	defer wg.Done()

	app := https.NewApp()
	// 注册拦截器
	filter.Filter(app)
	// 注册路由
	routers.Router(app)
	// 启动服务
	app.Server(cfg.Http)
}
