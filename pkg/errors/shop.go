package errors

import "github.com/go-kratos/kratos/v2/errors"

// 如果无需客户端做多语言兼容,可以在此定义固定的Reason 跟 Message
var (
	InvalidParams  = errors.New(400, "InvalidParams", "Missing Params")
	UserNotExit    = errors.New(400, "UserNotExit", "User Not Exit")
	RecordNotFound = errors.New(404, "RecordNotFound", "Record Not Found")
	UnknownError   = errors.New(500, "UnknownError", "Unknown Errors")
	MakeTokenFaild = errors.New(500, "MakeTokenFaild", "Make Token Faild")
)
