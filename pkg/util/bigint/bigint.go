package bigint

import (
	"fmt"
	"math/big"
)

var sock = "1000000000000000000"

// ParserPrice 浮点数 * 基数 返回运算后的字符串
func ParserPrice(price float64) string {
	var a, _ = new(big.Rat).SetString(fmt.Sprintf("%f", price))
	var b, _ = new(big.Rat).SetString(sock)
	a.Mul(a, b)
	var c, _ = new(big.Rat).SetString(a.String())
	return c.RatString()
}

// PrecisionPrice 转6位大数
func PrecisionPrice6(price uint64) string {
	return new(big.Int).Mul(big.NewInt(1e+6), new(big.Int).SetUint64(price)).String()
}
