package cmd

import (
	"sync"

	"github.com/cloakscn/share-word/server/config/https"
	"github.com/cloakscn/share-word/server/config/redis"
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
	go https.Server(cfg.Http, &wg)
	wg.Wait()
}
