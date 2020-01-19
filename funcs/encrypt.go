package funcs

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash 获取md5加密串
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
