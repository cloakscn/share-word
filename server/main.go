package main

import (
	"github.com/namsral/flag"

	"github.com/cloakscn/share-word/server/cmd"
	"github.com/cloakscn/share-word/server/config/https"
	"github.com/cloakscn/share-word/server/config/redis"
)

var (
	port      = flag.String("port", "8080", "对外开放的服务端口")
	redisAddr = flag.String("redisAddr", "docker.cloaks.cn:6379", "redis 连接地址")
)

func main() {
	flag.Parse()
	// parse config
	cmd.Start(&cmd.Config{
		Redis: &redis.Config{
			// 默认不需要密码，连接 0 号数据库
			Addr: *redisAddr,
			DB:   0,
		},
		Http: &https.Config{
			Port: *port,
		},
	})
}
