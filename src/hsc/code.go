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
	ERROR_WRITE_MYSQL              = 50005
	WARN_CLICK_REPEATEDLY          = 60000
	WARN_NO_USE                    = 60002
	WARN_BACKGROUND                = 60004
	WARN_NOT_FOUND_CLOUD           = 60005
	WARN_NOT_PROMISE_RULE          = 60006
	WARN_USER_NAME_EXIST           = 60007
	WARN_USER_MAIL_EXIST           = 60008
	WARN_USER_PASSWORD_CHECK       = 60009
	WARN_CODIS_NOT_CONNECT         = 60010
	WARN_CODIS_IS_REBALANCE        = 60011
	WARN_CODIS_NOT_OPTION          = 60012
	WARN_CODIS_PROXY_MIN_NUMBER    = 60013
	WARN_CODIS_GROUP_MIN_NUMBER    = 60014
	WARN_CODIS_GROUP_MIN_CAPACITY  = 60015
	WARN_CHECK_IPPORT_FAIL         = 60016
)
