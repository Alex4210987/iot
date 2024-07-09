package api

import (
	"backend/client/database"
	"backend/http_param"
	"backend/model"
	"backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// MCUAuth 单片机认证
func MCUAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求表单
		var form http_param.MCUAuthInfo
		err := c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "请求参数错误",
			})
			return
		}
		// 通过mac地址获取单片机
		var mcu model.MCU
		mcu.Mac = form.Mac
		err = mcu.GetMCUByMac(database.DbEngine)
		if err != nil {
			// 没有找到单片机，生成一个认证码
			mcu.AuthCode = util.GenerateAuthCode()
			mcu.Time = time.Now()
			// 保存单片机
			err = mcu.Save(database.DbEngine)
			// 返回认证码
			c.JSON(200, gin.H{
				"status":   "success",
				"authCode": mcu.AuthCode,
			})
			return
		}
		// 找到单片机，查看认证码是否正确
		if mcu.AuthCode == form.AuthCode {
			// 认证码正确, 生成一个 token
			token, err := util.CreateToken(mcu.AuthCode)
			if err != nil {
				c.JSON(200, gin.H{
					"status": "error",
					"msg":    "生成 token 失败",
				})
				return
			} else {
				// 返回 token
				c.JSON(200, gin.H{
					"status": "success",
					"token":  token,
				})
				return
			}
		} else {
			// 认证码错误
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "认证码错误",
			})
			return
		}
	}
}

// MCUData 单片机上传数据
func MCUData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求表单
		var form http_param.MCUData
		err := c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "请求参数错误",
			})
			fmt.Println(err)
			return
		}
		value, _ := c.Get("mcu")
		mcu := value.(model.MCU)
		// 保存数据
		data := model.Data{
			McuID:         mcu.ID,
			Humidity:      form.Humidity,
			Temperature:   form.Temperature,
			Raindrops:     form.Raindrops,
			EarthHumidity: form.EarthHumidity,
			BootTime:      form.TimeStamp,
			Timestamp:     time.Now(),
		}
		err = data.Save(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "保存数据失败",
			})
			return
		}

		// 如果McuDataMap map[int]chan Data中没有该单片机的数据通道，则创建一个
		if _, ok := model.McuDataMapChan[mcu.ID]; !ok {
			model.McuDataMapChan[mcu.ID] = make(chan model.Data, 1)
		}
		// 将数据覆盖到通道中
		// 锁
		model.Mu.Lock()
		model.McuDataMapChan[mcu.ID] <- data
		// 解锁
		model.Mu.Unlock()

		// 返回成功
		c.JSON(200, gin.H{
			"status": "success",
		})
		return
	}
}
