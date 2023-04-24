package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(str []byte) string {
	handler := md5.New()
	handler.Write(str)
	return hex.EncodeToString(handler.Sum(nil))
}
