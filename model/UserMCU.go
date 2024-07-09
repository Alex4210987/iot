package model

import "gorm.io/gorm"

// UserMCU 用户单片机
type UserMCU struct {
	ID     int    `json:"id"`      // 主键
	UserID int    `json:"user_id"` // 用户 ID
	McuID  int    `json:"mcu_id"`  // 单片机 ID
	Name   string `json:"name"`    // 单片机的名字
}

// Save 保存用户单片机
func (userMcu *UserMCU) Save(db *gorm.DB) (err error) {
	err = db.Create(userMcu).Error
	return
}

// GetUserMCUListByUserID 根据用户 ID 获取用户单片机
func (userMcu *UserMCU) GetUserMCUListByUserID(db *gorm.DB) (userMcus []UserMCU, err error) {
	err = db.Where("user_id = ?", userMcu.UserID).Find(&userMcus).Error
	return
}

// GetUserMCUByID 根据 ID 获取用户单片机
func (userMcu *UserMCU) GetUserMCUByID(db *gorm.DB) (err error) {
	err = db.Where("id = ?", userMcu.ID).First(userMcu).Error
	return
}

// GetUserMCUByUserIDAndMCUID 根据用户 ID 和单片机 ID 获取用户单片机
func (userMcu *UserMCU) GetUserMCUByUserIDAndMCUID(db *gorm.DB) (err error) {
	err = db.Where("user_id = ? AND mcu_id = ?", userMcu.UserID, userMcu.McuID).First(userMcu).Error
	return
}
