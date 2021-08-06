package errreason

// user-service errors text
// error_code was defined in user_error.proto

var (
	INVALID_PARAMS   = "参数不可用"
	USER_NOT_EXIT    = "用户不存在"
	INVALID_PASS     = "用户名或密码不正确"
	MAKE_TOKEN_ERROR = "token生成失败"
)
