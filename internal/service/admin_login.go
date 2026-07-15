package service

import (
	"Diggpher/global"
	"Diggpher/internal/dao"
	"Diggpher/internal/service/errMsg"
	"Diggpher/pkg/middleware/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminLoginResp struct {
	Code   uint   `json:"code"`
	ErrMsg string `json:"errMsg"`
	Token  string `json:"token"`
}

func (a *AdminService) Login(username, password, loginIp string) *AdminLoginResp {
	resp := new(AdminLoginResp)

	var admin dao.Admin
	if errors.Is(global.DataBase.Where("username = ?", username).First(&admin).Error, gorm.ErrRecordNotFound) {
		resp.Code = errMsg.ErrorAdminUserNotFound
		return resp
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		// 兼容旧明文密码
		if admin.Password != password {
			resp.Code = errMsg.ErrorAdminPswError
			return resp
		}
	}

	token, err := auth.GenerateAdminToken(admin.ID)
	if err != nil {
		resp.Code = errMsg.ErrorAdminJWT
		resp.ErrMsg = err.Error()
		return resp
	}
	resp.Token = token

	if admin.LastIp != loginIp {
		admin.LastIp = loginIp
		admin.Ips += loginIp + ","
		global.DataBase.Save(&admin)
	}
	return resp
}
