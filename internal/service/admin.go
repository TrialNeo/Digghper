package service

import (
	"Diggpher/global"
	"Diggpher/internal/dao"
	"Diggpher/internal/service/errMsg"
	"Diggpher/pkg/crypto"
	"Diggpher/pkg/middleware/auth"
	"Diggpher/pkg/utils"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type AdminService struct{}

type LoginResp struct {
	Code   uint   `json:"code"`
	ErrMsg string `json:"errMsg"`
	Token  string `json:"token"`
}

func (*AdminService) Login(username, password, loginIp string) *LoginResp {
	var (
		admin dao.Admin
		resp  = new(LoginResp)
	)

	if errors.Is(global.DataBase.Where("username = ?", username).First(&admin).Error, gorm.ErrRecordNotFound) {
		resp.Code = errMsg.ErrorAdminUserNotFound
		resp.ErrMsg = errMsg.GetErrMsg(resp.Code)
		return resp
	}

	//是不是来撞库的
	if !utils.SliceContains(strings.Split(admin.Ips, ","), loginIp) {
	}

	//密码错误
	if crypto.PswEnc(password) != admin.Password {
		resp.Code = errMsg.ErrorAdminPswError
		resp.ErrMsg = errMsg.GetErrMsg(resp.Code)
		return resp
	}

	//jwt签发
	token, err := auth.GenerateToken(admin.ID)
	if err != nil {
		resp.Code = errMsg.ErrorAdminJWT
		resp.ErrMsg = err.Error()
	}
	resp.Token = token

	//刷新登录ip 防止出问题,gorm自己有乐观锁也不用怕。
	if admin.LastIp == loginIp {
		return resp
	}

	//保存一下新登录IP
	global.DataBase.Save(&admin)
	admin.Ips += loginIp + ","
	return resp
}
