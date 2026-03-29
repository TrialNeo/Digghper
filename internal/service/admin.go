package service

import (
	"Diggpher/global"
	"Diggpher/internal/dao"
	"Diggpher/internal/service/errMsg"
	"Diggpher/pkg/crypto"
	"Diggpher/pkg/middleware/auth"
	"Diggpher/pkg/utils"
	"errors"
	"go.uber.org/zap"
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

	global.Log.Info("Admin login attempt", zap.String("username", username), zap.String("ip", loginIp))

	if errors.Is(global.DataBase.Where("username = ?", username).First(&admin).Error, gorm.ErrRecordNotFound) {
		global.Log.Warn("Admin user not found", zap.String("username", username), zap.String("ip", loginIp))
		resp.Code = errMsg.ErrorAdminUserNotFound
		resp.ErrMsg = errMsg.GetErrMsg(resp.Code)
		return resp
	}

	//是不是来撞库的
	if !utils.SliceContains(strings.Split(admin.Ips, ","), loginIp) {
		global.Log.Warn("Admin login from unauthorized IP", zap.String("username", username), zap.String("ip", loginIp))
	}

	//密码错误
	if crypto.PswEnc(password) != admin.Password {
		global.Log.Warn("Admin password error", zap.String("username", username), zap.String("ip", loginIp))
		resp.Code = errMsg.ErrorAdminPswError
		resp.ErrMsg = errMsg.GetErrMsg(resp.Code)
		return resp
	}

	//jwt签发
	token, err := auth.GenerateToken(admin.ID)
	if err != nil {
		global.Log.Error("Admin JWT generation error", zap.String("username", username), zap.Error(err))
		resp.Code = errMsg.ErrorAdminJWT
		resp.ErrMsg = err.Error()
	} else {
		global.Log.Info("Admin login success", zap.String("username", username), zap.String("ip", loginIp))
	}
	resp.Token = token

	//刷新登录ip 防止出问题,gorm自己有乐观锁也不用怕。
	if admin.LastIp == loginIp {
		return resp
	}

	//保存一下新登录IP
	admin.LastIp = loginIp
	admin.Ips += loginIp + ","
	if err := global.DataBase.Save(&admin).Error; err != nil {
		global.Log.Error("Failed to update admin IP", zap.String("username", username), zap.Error(err))
	} else {
		global.Log.Info("Admin IP updated", zap.String("username", username), zap.String("ip", loginIp))
	}
	return resp
}
