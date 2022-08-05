package errors

import "github.com/go-kratos/kratos/v2/errors"

// 如果无需客户端做多语言兼容,可以在此定义固定的Reason 跟 Message
var (
	ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")
)
