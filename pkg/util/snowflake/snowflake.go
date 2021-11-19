/*
 * @Author: Casso
 * @Date: 2021-11-18 19:47:56
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-19 10:32:59
 * @Description: 雪花算法
 * @FilePath: /kratos-mono-repo/pkg/util/snowflake/snowflake.go
 */
package snowflake

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node = nil

func init() {
	nodes, err := snowflake.NewNode(1)
	if err != nil {
		panic("snowflake init faild")
	}
	node = nodes
}

// RandomUID 简易发号器
func RandomUID() int64 {
	return node.Generate().Int64()
}
