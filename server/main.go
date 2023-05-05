package main

import (
	"github.com/cloakscn/share-word/cmd"
	"github.com/cloakscn/share-word/config/https"
	"github.com/cloakscn/share-word/config/redis"
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
