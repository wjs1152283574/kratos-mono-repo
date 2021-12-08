/*
 * @Author: Casso
 * @Date: 2021-11-23 17:49:03
 * @LastEditors: Casso
 * @LastEditTime: 2021-12-08 14:25:06
 * @Description: 字符串处理
 * @FilePath: /kratos-mono-repo/pkg/util/str/str.go
 */
package str

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Encode 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
