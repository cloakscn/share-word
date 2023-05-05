package req

import "time"

type CreateWorld struct {
	Type        int           `json:"type"`        // 文本类型
	Content     string        `json:"content"`     // 内容
	ContentType int           `json:"contentType"` // 内容类型
	Expiration  time.Duration `json:"expiration"`  // 有效期
	Encrypt     bool          `json:"encrypt"`     // 是否加密
}

type GetByID struct {
	Key      string `json:"key"`
	Password string `json:"password"`
}
