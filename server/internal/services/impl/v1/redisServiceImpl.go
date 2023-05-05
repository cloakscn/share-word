package v1

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	x "github.com/cloakscn/share-word/server/config/redis"
	"github.com/cloakscn/share-word/server/internal/vo/req"
	"github.com/cloakscn/share-word/server/utils/constant"
	"github.com/cloakscn/share-word/server/utils/enums"
	"github.com/cloakscn/share-word/server/utils/errors"
)

type RedisServiceImpl struct{}

func NewRedisServiceImpl() *RedisServiceImpl {
	return &RedisServiceImpl{}
}

func (s *RedisServiceImpl) GetValue(data req.GetByID) (result string, err error) {
	var dataType = strings.Split(data.Key, "::")[0]

	switch dataType {
	case constant.ONCE: // 阅后即焚
		result, err = x.Server.Get(data.Key).Result()
		if err != nil {
			return result, err
		}
		_, err = x.Server.Del(data.Key).Result()
		if err != nil {
			return result, err
		}
	case constant.DISCUSS: // 讨论
	default:
		return "", errors.ErrorBadParam
	}

	return result, err
}

func (s *RedisServiceImpl) SetValue(data req.CreateWorld) (map[string]interface{}, error) {
	var (
		key        string
		value      map[string]interface{}
		expiration = data.Expiration * time.Minute
		password   string
		err        error
	)

	value = map[string]interface{}{
		"content":     data.Content,
		"contentType": data.ContentType,
	}

	// 判断文本类型
	switch data.Type {
	case enums.ONCE:
		key, err = x.GetRedisKey(constant.ONCE, value)
	case enums.DISCUSS:
		key, err = x.GetRedisKey(constant.DISCUSS, value)
	default:
		return nil, errors.ErrorBadParam
	}

	// 判断是否要加密
	if data.Encrypt {
		key = fmt.Sprintf("%s_1", key)
		password = "1234"
	} else {
		key = fmt.Sprintf("%s_0", key)
	}

	if err != nil {
		return nil, err
	}

	marshal, _ := json.Marshal(value)
	err = x.Server.Set(key, string(marshal), expiration).Err()
	return map[string]interface{}{
		"key": key,
		"pwd": password,
	}, err
}
