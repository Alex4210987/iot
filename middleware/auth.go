package middleware

import (
	"backend/client/database"
	"backend/model"
	"backend/util"
	"github.com/gin-gonic/gin"
	"time"
)

// UserAuth 用户认证
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取 token
		token := c.Request.Header.Get("token")
		// 将 token 解析为用户名和时间戳
		username, oldTime, err := util.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "token 解析失败",
			})
			c.Abort()
			return
		}
		// 查看时间是否超过 7 天
		if time.Now().Sub(oldTime).Hours() > 7*24 {
			c.JSON(401, gin.H{
				"message": "登录信息已过期，请重新登录",
			})
			c.Abort()
			return
		}
		// 检测用户名是否存在
		user := model.User{Username: username}
		err = user.GetUserByUsername(database.DbEngine)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "寻找用户失败",
			})
			c.Abort()
			return
		}
		// 将用户信息写入上下文
		c.Set("user", user)
		// 继续执行
		c.Next()
		return
	}
}

// MCUAuth MCU 认证
func MCUAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取 token
		token := c.Request.Header.Get("token")
		// 将 token 解析为MCU的认证码和时间戳
		authCode, oldTime, err := util.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "token 解析失败",
			})
			c.Abort()
			return
		}
		// 查看时间是否超过 7 天
		if time.Now().Sub(oldTime).Hours() > 7*24 {
			c.JSON(401, gin.H{
				"message": "登录信息已过期，请重新登录",
			})
			c.Abort()
			return
		}
		// 检测认证码是否存在
		mcu := model.MCU{AuthCode: authCode}
		err = mcu.GetMCUByAuthCode(database.DbEngine)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "寻找MCU失败",
			})
			c.Abort()
			return
		}
		// 将MCU信息写入上下文
		c.Set("mcu", mcu)
		// 继续执行
		c.Next()
	}
}
