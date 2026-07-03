package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Encode 返回 32 位 MD5 加密字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
