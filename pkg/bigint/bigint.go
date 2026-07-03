package bigint

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// ParserPrice 转大数专用，传入价格和精度
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

// Commission 获取扣除手续费后的金额，传入总价大数和手续费率（例如 0.2）
func Commission(total string, commiss float64) (string, error) {
	bbint, ok := new(big.Int).SetString(total, 10)
	if !ok {
		return "", errors.New("bigint SetString fail")
	}
	f1 := big.NewFloat(float64(bbint.Int64()))
	f3 := big.NewFloat(float64(bbint.Int64()))
	yk := f1.Mul(f1, big.NewFloat(commiss))
	result := f1.Sub(f3, yk)
	return result.String(), nil
}

// StringToBigInt string to bigint
func StringToBigInt(amount string) (*big.Int, bool) {
	return new(big.Int).SetString(amount, 10)
}
