package normal

import "github.com/go-kratos/kratos/v2/errors"

// 客户端错误
var InvalidParams = errors.New(400, "InvalidParams", "Missing Params")

var RecordNotFound = errors.New(404, "RecordNotFound", "Record Not Found")

// 服务端错误
var MakeTokenFaild = errors.New(500, "MakeTokenFaild", "Make Token Faild")
