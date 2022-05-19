package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	InvalidParams  = errors.New(400, "InvalidParams", "Missing Params")
	UserNotExit    = errors.New(400, "UserNotExit", "User Not Exit")
	RecordNotFound = errors.New(404, "RecordNotFound", "Record Not Found")
	UnknownError   = errors.New(500, "UnknownError", "Unknown Errors")
	MakeTokenFaild = errors.New(500, "MakeTokenFaild", "Make Token Faild")
)
