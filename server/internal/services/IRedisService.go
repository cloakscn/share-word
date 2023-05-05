package services

import (
	"github.com/cloakscn/share-word/server/internal/vo/req"
)

type IRedisService interface {
	GetValue(data req.GetByID) (string, error)
	SetValue(data req.CreateWorld) (map[string]interface{}, error)
}
