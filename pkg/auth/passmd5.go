package auth

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// MD5 返回 MD5 加密字符串
func MD5(params []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(params)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// Base64Md5 先 base64 再 MD5，用于密码加密
func Base64Md5(params string) string {
	return MD5([]byte(base64.StdEncoding.EncodeToString([]byte(params))))
}
