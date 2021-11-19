/*
 * @Author: Casso
 * @Date: 2021-11-18 19:52:20
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-18 19:54:40
 * @Description: file content
 * @FilePath: /kratos-mono-repo/pkg/util/snow/snow_test.go
 */
package snow

import (
	"fmt"
	"testing"
)

func TestNewSnow(t *testing.T) {
	fmt.Println(NewSnowFlake())
}
