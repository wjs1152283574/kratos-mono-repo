/*
 * @Author: Casso
 * @Date: 2021-11-18 19:52:20
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-19 14:57:49
 * @Description: file content
 * @FilePath: /kratos-mono-repo/pkg/util/snowflake/snowflake_test.go
 */
package snowflake

import (
	"fmt"
	"testing"
)

func TestNewSnow(t *testing.T) {
	fmt.Println(RandomUID()) // go test -v snowflake_test.go snowflake.go
}
