package model

import (
	"gorm.io/gorm"
	"time"
)

// MCU 单片机
type MCU struct {
	ID       int       `json:"id"`       // 单片机的 ID
	Mac      string    `json:"mac"`      // 单片机的 MAC 地址
	AuthCode string    `json:"authCode"` // 单片机的认证码
	Time     time.Time `json:"time"`     // 单片机的时间
}

// GetMCUByAuthCode 根据认证码获取单片机
func (mcu *MCU) GetMCUByAuthCode(db *gorm.DB) (err error) {
	err = db.Where("auth_code = ?", mcu.AuthCode).First(mcu).Error
	return
}

// GetMCUByMac 根据 MAC 地址获取单片机
func (mcu *MCU) GetMCUByMac(db *gorm.DB) (err error) {
	err = db.Where("mac = ?", mcu.Mac).First(mcu).Error
	return
}

// GetMCUByID 根据 ID 获取单片机
func (mcu *MCU) GetMCUByID(db *gorm.DB) (err error) {
	err = db.Where("id = ?", mcu.ID).First(mcu).Error
	return
}

// Save 保存单片机
func (mcu *MCU) Save(db *gorm.DB) (err error) {
	err = db.Create(mcu).Error
	return
}
