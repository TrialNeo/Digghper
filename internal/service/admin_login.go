package service

import (
	"context"
	"errors"

	"Diggpher/internal/query"
	"Diggpher/internal/service/errMsg"
	"Diggpher/pkg/middleware/auth"

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
	ctx := context.Background()

	adminDO := query.Admin
	admin, err := adminDO.WithContext(ctx).Where(adminDO.Username.Eq(username)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Code = errMsg.ErrorAdminUserNotFound
		return resp
	}
	if err != nil {
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
		_ = adminDO.WithContext(ctx).Save(admin)
	}
	return resp
}
