/*
 * @Author: Casso
 * @Date: 2021-11-18 19:47:56
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-18 19:54:28
 * @Description: 雪花算法
 * @FilePath: /kratos-mono-repo/pkg/util/snow/snow.go
 */
package snow

import "github.com/bwmarrin/snowflake"

// var node *snowflake.Node = nil

func init() {
	_, err := snowflake.NewNode(1)
	if err != nil {
		panic("snowflake init faild")
	}
	// node = nodes
}

func NewSnowFlake() int64 {
	node, _ := snowflake.NewNode(1)
	return node.Generate().Int64()
}
