package dao

import (
	"time"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` // json:"-" 表示返回给前端时隐藏密码
	LastIp    string    `gorm:"type:varchar(255);" json:"lastIp"`
	Ips       string    `gorm:"type:text" json:"ips"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
