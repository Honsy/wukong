package enum

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	// USER
	LOGIN_FAIL = 10001
	// AUTH
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	// GB28181
	ERROR_GET_DEVICE_LIST = 30001
)
