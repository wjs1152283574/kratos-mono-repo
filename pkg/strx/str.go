package strx

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// IsGBK 检测数据是否为 GBK 编码
func IsGBK(data []byte) bool {
	length := len(data)
	i := 0
	for i < length {
		if data[i] <= 0x7f {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

// Int64ToStr int64 转 string
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

// IntToStr int 转 string
func IntToStr(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

// Float32ToStr float32 转 string
func Float32ToStr(num float32) string {
	return fmt.Sprintf("%f", num)
}

// Float64ToStr float64 转 string
func Float64ToStr(num float64) string {
	return fmt.Sprintf("%f", num)
}

// StrToFloat64 string 转 uint64
func StrToFloat64(str string) (uint64, error) {
	bint, ok := new(big.Int).SetString(str, 0)
	if !ok {
		return 0, errors.New("SetString err")
	}
	return bint.Uint64(), nil
}
