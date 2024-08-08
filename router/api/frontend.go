package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5/model"
	// gorm
)

type Device struct {
	Industry   string    `json:"industry"`
	AppID      string    `json:"app_id"`
	CreateTime time.Time `json:"create_time"`
}

type DeviceResponse struct {
	Products []Device `json:"products"`
}

func GetAllDevices(c *gin.Context) {
	if HWClient == nil {
		SettingUpEnvironment()
		InitHuaweiCloudClient()
	}
	client := HWClient
	request := &model.ListProductsRequest{}
	response, err := client.ListProducts(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		c.JSON(400, gin.H{"HWerror": err.Error()})
	}

	var res []Device
	var t time.Time
	for _, product := range *response.Products {
		// 20240608T040435Z->2024-06-08 04:04:35
		t, _ = time.Parse("20060102T150405Z", *product.CreateTime)
		fmt.Println(t)
		res = append(res, Device{
			Industry:   *product.Industry,
			AppID:      *product.AppId,
			CreateTime: t,
		})
	}

	c.JSON(200, res)
}

// {
//     "products": [
//         {
//             "app_id": "8a260d5a0983410ca60342bde21b54a0",
//             "app_name": "DefaultApp_665cvmzh",
//             "product_id": "66641fa77dbfd46fabbf5db0",
//             "name": "Face_Recognition_Access_Control",
//             "device_type": "门口机",
//             "protocol_type": "MQTT",
//             "data_format": "binary",
//             "industry": "智能生活-家居安防",
//             "create_time": "20240608T090855Z"
//         },
//         {
//             "app_id": "8a260d5a0983410ca60342bde21b54a0",
//             "app_name": "DefaultApp_665cvmzh",
//             "product_id": "6663d8537dbfd46fabbf54b9",
//             "name": "Smart_Park",
//             "device_type": "智能水质监测器",
//             "protocol_type": "MQTT",
//             "data_format": "json",
//             "industry": "智慧城市-环境感知",
//             "create_time": "20240608T040435Z"
//         }
//     ],
//     "page": {
//         "count": 2,
//         "marker": "6663d8537dbfd46fabbf54bb"
//     }
// }

type Data struct {
	Timestamp       time.Time `gorm:"column:TIME;type:datetime;default:null" json:"time_stamp"`
	Temperature     float64   `gorm:"column:TEMPERATURE" json:"temperature"`
	Humidity        float64   `gorm:"column:HUMIDITY" json:"humidity"`
	SoilMoisture    float64   `gorm:"column:SOIL_MOISTURE" json:"soil_moisture"`
	WaterDischarge  float64   `gorm:"column:WATER_DISCHARGE" json:"water_discharge"`
	ElectricCurrent float64   `gorm:"column:ELECTRIC_CURRENT" json:"electric_current"`
	Sunlight        float64   `gorm:"column:SUN_LIGHT" json:"sunlight"`
}

// 0.key(编号，递增编号)
// 1.time（DATETIME格式为'YYYY-MM-DD HH:MM:SS'VALUES ('2024-07-08 14:30:00', 25.3, 60.5, 42, 0.45, 12.3, 700.0);）

// - 1.temperature（浮点数）
// - 2.humindity(float)
// - 3.air_quality(int)
// - 4.electric_current(float)
// - 5.water_discharge(float)
// - 6.sun_light(float)
// - 7.soil_moisure(float)

// db connection: mysql, root@114.116.221.179, db:iot,
// password read from .env
// using gorm select the most recent 100 records from the table

func GetHistoryData(c *gin.Context) {
	if dbConn == nil {
		InitDB()
	}
	var data []Data
	dbConn.Table("iot").Order("time desc").Limit(100).Find(&data)
	c.JSON(200, data)
}

func GetNewestData(c *gin.Context) {
	if dbConn == nil {
		InitDB()
	}
	var data Data
	dbConn.Table("iot").Order("time desc").Limit(1).Find(&data)
	c.JSON(200, data)
}
