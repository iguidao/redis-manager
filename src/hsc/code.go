package hsc

// 提供状态码
const (
	SUCCESS        = 0
	ERROR          = 500
	INVALID_PARAMS = 400
	NO_LOGIN       = 401
	NOT_PROMISE    = 403
	SERVER_ERROR   = 544

	NOT_FOUND                      = 10001
	Method_FAILS                   = 10002
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	MESSAGE_RE                     = 400
	ERROR_NO_CONNEC                = 50001
	ERROR_BACKGROUND               = 50002
	ERROR_CLOUD_CONNECT            = 50003
	ERROR_CLOUD_GET                = 50004
	WARN_CLICK_REPEATEDLY          = 60000
	WARN_NO_USE                    = 60002
	WARN_BACKGROUND                = 60004
)
