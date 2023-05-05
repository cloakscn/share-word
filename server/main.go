package main

import (
	"github.com/cloakscn/share-word/server/cmd"
	"github.com/cloakscn/share-word/server/config/https"
	"github.com/cloakscn/share-word/server/config/redis"
)

func main() {
	// parse config
	cmd.Start(&cmd.Config{
		Redis: &redis.Config{
			Addr: "docker.cloaks.cn:6379",
			DB:   0,
		},
		Http: &https.Config{
			Port: "8080",
		},
	})
}
