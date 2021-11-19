/*
 * @Author: Casso
 * @Date: 2021-11-18 19:47:56
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-19 16:41:04
 * @Description: 雪花发号器
 * @FilePath: /kratos-mono-repo/pkg/util/snowflake/snowflake.go
 */
package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node = nil

func init() {
	if nodes, err := snowflake.NewNode(1); err != nil {
		panic("snowflake init faild")
	} else {
		node = nodes
	}
}

// RandomUID 简易发号器
func RandomUID() int64 {
	return node.Generate().Int64()
}
