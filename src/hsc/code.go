package hsc

// 提供状态码
const (
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400
	NO_LOGIN       = 401
	NOT_PROMISE    = 403

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	NO_CONNECT_CODIS               = 50001
)
