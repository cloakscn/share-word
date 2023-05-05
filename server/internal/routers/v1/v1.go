package v1

import (
	"github.com/cloakscn/share-word/internal/services"
	v1 "github.com/cloakscn/share-word/internal/services/impl/v1"
	"github.com/cloakscn/share-word/internal/vo/req"
	"github.com/cloakscn/share-word/utils/ajax"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	redisService services.IRedisService = v1.NewRedisServiceImpl()
)

func RouterV1(group *gin.RouterGroup) {
	group.POST("/", GetByID)
	group.POST("/create", Create)
}

func GetByID(context *gin.Context) {
	var (
		data   req.GetByID
		result string
		err    error
	)

	if err := context.ShouldBindJSON(&data); err != nil {
		ajax.Error(context, ajax.StatusBadRequest, err)
		return
	}

	if data.Key == "" {
		ajax.Error(context, ajax.StatusBadRequest, nil)
		return
	}

	// 判断是否加密
	encrypt := strings.Split(data.Key, "_")[2]
	switch encrypt {
	case "1": // 加密
		fallthrough
	case "0": // 未加密

		result, err = redisService.GetValue(data)
	default:
		ajax.Success(context, ajax.StatusBadRequest)
	}

	if err != nil {
		ajax.Error(context, ajax.StatusInternalServerError, err)
		return
	}

	ajax.Success(context, result)
}

func Create(context *gin.Context) {
	var data req.CreateWorld
	if err := context.ShouldBindJSON(&data); err != nil {
		ajax.Error(context, ajax.StatusBadRequest, err)
		return
	}

	result, err := redisService.SetValue(data)
	if err != nil {
		ajax.Error(context, ajax.StatusInternalServerError, err)
		return
	}

	ajax.Success(context, result)
}

/**

// 哈希表操作
err = client.HSet("hash", "field", "value").Err()
if err != nil {
	panic(err)
}

val2, err := client.HGet("hash", "field").Result()
if err != nil {
	panic(err)
}
fmt.Println("hash", val2)

// 列表操作
err = client.RPush("list", "item1", "item2").Err()
if err != nil {
	panic(err)
}

val3, err := client.LPop("list").Result()
if err != nil {
	panic(err)
}
fmt.Println("list", val3)

// 集合操作
err = client.SAdd("set", "member1", "member2").Err()
if err != nil {
	panic(err)
}

val4, err := client.SMembers("set").Result()
if err != nil {
	panic(err)
}
fmt.Println("set", val4)
*/
