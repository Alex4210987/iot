package router

import (
	"backend/middleware"
	"backend/router/api"
	"github.com/gin-gonic/gin"
)

func UseMyRouter(r *gin.Engine) {
	mcu := r.Group("/mcu")
	{
		mcu.POST("/auth",
			api.MCUAuth(),
		)
		mcu.POST("/data",
			middleware.MCUAuth(),
			api.MCUData(),
		)
	}

	// 用户接口
	user := r.Group("/user")
	{
		user.POST("/login",
			api.UserLogin(),
		)
		user.POST("/register",
			api.UserRegister(),
		)
		user.POST("/bind_mcu",
			middleware.UserAuth(),
			api.UserBindMCU(),
		)
		user.GET("/data/realtime",
			api.UserRealTimeData(),
		)

		user.GET("/data/history",
			middleware.UserAuth(),
			api.UserHistoryData(),
		)
		user.GET("/data/analysis") //UserAnalysisData,
		// 获取所有绑定的 MCU
		user.GET("/mcuList",
			middleware.UserAuth(),
			api.UserMCUList(),
		)
		// 获取所有的用户
		user.GET("/userList",
			middleware.UserAuth(),
			api.UserList(),
		)
		// 获取用户名字
		user.GET("/username",
			middleware.UserAuth(),
			api.Username(),
		)
		// 获取gpt回复
		user.GET("/gpt",
			api.GPT(),
		)

	}
}
