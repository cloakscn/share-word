package redis

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cloakscn/share-word/server/utils/md5"
	"github.com/go-redis/redis"
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

func GetRedisKey(namespace string, value map[string]interface{}) (string, error) {
	marshal, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s::%s_%s", namespace, md5.GetMD5(marshal), strconv.FormatInt(time.Now().Unix(), 10)), nil
}
