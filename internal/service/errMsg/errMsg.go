package errMsg

const (
	SUCCESS = uint(0)
	ERROR   = uint(500)
)

// Admin模块
const (
	ErrorAdminPswError     = uint(10001)
	ErrorAdminUserNotFound = uint(10002)
	ErrorAdminJWT          = uint(10003)
	ErrorTokenInvalid      = uint(9999)
)

var codeMsg = map[uint]string{
	SUCCESS:                "GetErrMsg",
	ErrorAdminPswError:     "密码错误",
	ErrorAdminUserNotFound: "用户不存在",
	ErrorTokenInvalid:      "Token已失效",
}

func GetErrMsg(code uint) string {
	return codeMsg[code]
}
