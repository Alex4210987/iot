package model

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

// Data 数据表
type Data struct {
	ID            int       `json:"id"`            // 主键
	McuID         int       `json:"mcu_id"`        // 单片机 ID
	Humidity      float32   `json:"humidity"`      // 湿度
	Temperature   float32   `json:"temperature"`   // 温度
	Raindrops     int32     `json:"precipitation"` // 雨滴
	EarthHumidity float32   `json:"soilHumidity"`  // 土壤湿度
	BootTime      uint32    `json:"boot_time"`     // 开机时间
	Timestamp     time.Time `json:"timestamp"`     // 数据上传的时间
}

// McuDataMap 单片机数据映射
type McuDataMap map[int]chan Data

var McuDataMapChan McuDataMap

// Mu 互斥锁
var Mu sync.Mutex

// InitMap 初始化 map
func InitMap() {
	McuDataMapChan = make(McuDataMap)
}

// Save 保存数据
func (data *Data) Save(db *gorm.DB) (err error) {
	err = db.Create(data).Error
	return
}

// GetDatasByMCUID 根据单片机 ID 获取数据
func (data *Data) GetDatasByMCUID(db *gorm.DB) (datas []Data, err error) {
	// 按照时间降序排序, 取前 2000 条数据
	err = db.Where("mcu_id = ?", data.McuID).Order("timestamp desc").Limit(2000).Find(&datas).Error
	return
}

// GetLastDataByMCUID 根据单片机 ID 获取最新的数据
func (data *Data) GetLastDataByMCUID(db *gorm.DB) (err error) {
	err = db.Where("mcu_id = ?", data.McuID).Order("timestamp desc").First(data).Error
	return
}
