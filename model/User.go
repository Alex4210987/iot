package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	ID            int       `json:"id"`              // 用户的 ID
	Username      string    `json:"username"`        // 用户名
	Password      string    `json:"password"`        // 用户密码
	LastLoginTime time.Time `json:"last_login_time"` // 最后登录时间
}

// GetUserByUsername 根据用户名获取用户
func (user *User) GetUserByUsername(db *gorm.DB) (err error) {
	err = db.Where("username = ?", user.Username).First(user).Error
	return
}

// Save 保存用户
func (user *User) Save(db *gorm.DB) (err error) {
	// 如果用户已经存在，就更新用户
	if err = user.GetUserByUsername(db); err == nil {
		err = db.Save(user).Error
		return
	}
	err = db.Create(user).Error
	return
}

// GetUserList 获取用户列表
func (user *User) GetUserList(db *gorm.DB) (users []User, err error) {
	err = db.Find(&users).Error
	return
}
