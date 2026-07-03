package crypto

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic("snowflake init failed: " + err.Error())
	}
}

// GenID 生成雪花算法 ID
func GenID() int64 {
	return node.Generate().Int64()
}
