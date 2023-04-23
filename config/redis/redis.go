package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var (
	Server *redis.Client
)

type Config struct {
	Addr     string
	Password string // Redis 数据库密码
	DB       int    // Redis 数据库索引（0-15）
}

func Init(cfg *Config) {
	// 连接 Redis 数据库
	log.Println("初始化 Redis。")
	Server = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	log.Println("初始化 Redis 成功。")
}
