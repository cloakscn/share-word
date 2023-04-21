package main

import (
	"github.com/cloakscn/share-word/cmd"
	"github.com/cloakscn/share-word/utils/https"
	"github.com/cloakscn/share-word/utils/redis"
)

func main() {
	// parse config
	cmd.Start(&cmd.Config{
		Redis: &redis.Config{
			Addr: "docker.cloaks.cn:6379",
		},
		Http: &https.Config{
			Port: "8080",
		},
	})
	//
	//// SET 和 GET 操作
	//err := client.Set("key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := client.Get("key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//// 哈希表操作
	//err = client.HSet("hash", "field", "value").Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val2, err := client.HGet("hash", "field").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("hash", val2)
	//
	//// 列表操作
	//err = client.RPush("list", "item1", "item2").Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val3, err := client.LPop("list").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("list", val3)
	//
	//// 集合操作
	//err = client.SAdd("set", "member1", "member2").Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val4, err := client.SMembers("set").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("set", val4)
}
