package bigint

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// PrecisionPrice 转大数专用 传入 价格，精度
func ParserPrice(price uint64, precision int64) (res string, err error) {
	if precision == 0 {
		res = fmt.Sprintf("%v", price)
		return
	}
	i, err := strconv.ParseFloat(fmt.Sprintf("1e+%v", precision), 64)
	if err != nil {
		return
	}

	res = new(big.Int).Mul(big.NewInt(int64(i)), new(big.Int).SetUint64(price)).String()
	return
}

// Commisssion 获取扣除手续费后的金额  传入 总价大数，手续费（例如0.2） 返回字符串
func Commisssion(total string, commiss float64) (string, error) {
	var err error
	bbint, ok := new(big.Int).SetString(total, 10)
	if !ok {
		return "", errors.New("bigint SetString fail")
	}
	f1 := big.NewFloat(float64(bbint.Int64()))
	f3 := big.NewFloat(float64(bbint.Int64()))
	yk := f1.Mul(f1, big.NewFloat(commiss))
	result := f1.Sub(f3, yk)
	return result.String(), err
}

//StringToBigInt string to bigint
func StringToBigInt(amount string) (*big.Int, bool) {
	bigIntAmount := new(big.Int)
	return bigIntAmount.SetString(amount, 10)
}
