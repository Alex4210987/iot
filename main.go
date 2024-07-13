package main

import (
	_const "backend/const"
	"backend/router"
	"backend/util"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	region "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	// "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
)

var (
	port   string
	client *iotda.IoTDAClient
	HWClient *iotda.IoTDAClient
	DeviceId string
)

func SettingUpEnvironment() {
	// 读取配置文件
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
	// 配置端口
	port = os.Getenv("FM_PORT")
	// 配置设备ID
	DeviceId = os.Getenv("DEVICE_ID")
	// 配置数据库
	// database.InitDB()
	// 配置常量
	_const.InitConst()
	// 初始化map
	// model.InitMap()
	// 初始化华为云客户端
	InitHuaweiCloudClient()
}

func InitHuaweiCloudClient() {
	// 从环境变量中获取 AK 和 SK
	ak := os.Getenv("CLOUD_SDK_AK")
	sk := os.Getenv("CLOUD_SDK_SK")
	// 定义 endpoint
	endpoint := os.Getenv("CLOUD_SDK_ENDPOINT")
	projectID := os.Getenv("CLOUD_SDK_PROJECT_ID")

	if ak == "" || sk == "" || endpoint == "" {
		panic("AK, SK or endpoint environment variables are not set")
	}

	auth := basic.NewCredentialsBuilder().
	WithAk(ak).
	WithSk(sk).
	WithProjectId(projectID).
	// 企业版/标准版需要使用衍生算法，基础版请删除该配置"WithDerivedPredicate"
	WithDerivedPredicate(auth.GetDefaultDerivedPredicate()).
	Build()

	client := iotda.NewIoTDAClient(
		iotda.IoTDAClientBuilder().
			// 标准版/企业版需要自行创建region，基础版使用IoTDARegion中的region对象
			WithRegion(region.NewRegion("cn-north-4", endpoint)).
			WithCredential(auth).
			Build())

	HWClient = client
}

func main() {
	// 初始化环境
	SettingUpEnvironment()
	commandParams := map[string]interface{}{
		"buzzer_switch": true,
  		"window_switch": true,
	}
	i := 0
	for i < 1000 {
		util.SendIoTCommand(HWClient, DeviceId, commandParams)
		i++
	}
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
