package main

import (
	"backend/router/api"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	"github.com/joho/godotenv"
)

var (
	port     string
	HWClient *iotda.IoTDAClient
	DeviceId string
)

func SettingUpEnvironment() {
	// 读取配置文件
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("fatal error loading .env file: %s", err))
	}
	// 配置端口
	port = os.Getenv("FM_PORT")
	// 配置设备ID
	DeviceId = os.Getenv("DEVICE_ID")
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

	auth, _ := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		WithProjectId(projectID).
		// 企业版/标准版需要使用衍生算法，基础版请删除该配置"WithDerivedPredicate"
		WithDerivedPredicate(auth.GetDefaultDerivedPredicate()).
		SafeBuild()

	builder, _ := iotda.IoTDAClientBuilder().
		// 标准版/企业版需要自行创建region，基础版使用IoTDARegion中的region对象
		WithRegion(region.NewRegion("cn-north-4", endpoint)).
		WithCredential(auth).
		SafeBuild()

	client := iotda.NewIoTDAClient(builder)

	HWClient = client
}

func main() {
	// 初始化环境
	SettingUpEnvironment()
	api.InitDB()
	// api.InitializeSwitches()

	// 初始化路由
	r := gin.Default()
	// 配置CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                                        // 允许所有域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}                   // 允许请求方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "token"} // 允许的头部

	r.Use(cors.New(config))
	// router.UseMyRouter(r)

	r.POST("/iot/messages", api.IotMessages())
	r.POST("/iot/completion", api.GetCompletions)

	// 添加人脸识别接口路由
	r.POST("/face/add", api.AddFaceHandler)
	r.POST("/face/search", api.SearchFaceHandler)

	r.GET("/mcuList", api.GetAllDevices)
	r.GET("/data/history", api.GetHistoryData)
	r.GET("/user/data/history", api.GetHistoryData)
	r.GET("/user/data/realtime", api.GetNewestData)
	r.POST("/frontend/buttons", api.HandleSwitch)

	r.POST("/login", api.UserLogin)
	r.POST("/register", api.UserRegister)
	r.GET("/online-users", api.GetOnlineUsers) // 查询在线用户的 API
	r.GET("/current-user", api.GetCurrentUser) // 查询当前用户的 API

	des := ":" + port
	_ = r.Run(des)
}
