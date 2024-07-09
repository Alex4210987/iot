package main

import (
	"backend/client/database"
	_const "backend/const"
	"backend/model"
	"backend/router"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var port string

func SettingUpEnvironment() {
	// 读取配置文件
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
	// 配置端口
	port = os.Getenv("FM_PORT")
	// 配置数据库
	database.InitDB()
	// 配置常量
	_const.InitConst()
	// 初始化map
	model.InitMap()

}

func main() {
	// 初始化环境
	SettingUpEnvironment()
	// 初始化路由
	r := gin.Default()
	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                        // 允许所有域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}                   // 允许请求方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "token"} // 允许的头部

	r.Use(cors.New(config))
	router.UseMyRouter(r)
	des := ":" + port
	_ = r.Run(des)
}
