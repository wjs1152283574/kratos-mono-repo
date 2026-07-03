package errors

import "github.com/go-kratos/kratos/v2/errors"

// 通用错误码
var (
	InvalidParams  = errors.New(400, "InvalidParams", "Missing Params")
	UserNotExit    = errors.New(400, "UserNotExit", "User Not Exit")
	RecordNotFound = errors.New(404, "RecordNotFound", "Record Not Found")
	UnknownError   = errors.New(500, "UnknownError", "Unknown Errors")
	MakeTokenFaild = errors.New(500, "MakeTokenFaild", "Make Token Faild")
)

// Shop BFF 服务错误码
var (
	ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")
)
