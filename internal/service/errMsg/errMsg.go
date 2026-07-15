package errMsg

const (
	SUCCESS            = 0
	ERROR              = 500
	ErrorDataBaseErr   = 501
	ErrorInvalidParams = 502
)

// Admin模块
const (
	ErrorAdminPswError     = 10001
	ErrorAdminUserNotFound = 10002
	ErrorAdminJWT          = 10003
	ErrorTokenInvalid      = 9999
)

var codeMsg = map[uint]string{
	SUCCESS:                "成功",
	ERROR:                  "服务器内部错误",
	ErrorInvalidParams:     "参数有误",
	ErrorDataBaseErr:       "数据库异常",
	ErrorAdminPswError:     "密码错误",
	ErrorAdminUserNotFound: "用户不存在",
	ErrorAdminJWT:          "JWT签发失败",
	ErrorTokenInvalid:      "Token已失效",
}

func GetErrMsg(code uint) string {
	return codeMsg[code]
}
