package api

import (
	"backend/model"
	"backend/util"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	iotda "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iotda/v5"
	"github.com/joho/godotenv"
)

var (
	HWClient            *iotda.IoTDAClient
	DeviceId            string
	ElectricCurrentFlag bool
	ExsitCount          int = 0
)

func SettingUpEnvironment() {
	// 读取配置文件
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("fatal error loading .env file: %s", err))
	}
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

func IotMessages() gin.HandlerFunc {
	return func(c *gin.Context) {
		if HWClient == nil {
			InitHuaweiCloudClient()
			SettingUpEnvironment()
		}

		rawData, err := c.GetRawData()
		if err != nil {
			fmt.Println("Failed to get raw data:", err)
		} else {
			fmt.Println("Request JSON:", string(rawData))
		}

		if len(rawData) == 0 {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "empty request",
			})
			return
		}

		var form model.Event
		err = json.Unmarshal(rawData, &form)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "failed to parse JSON",
			})
			fmt.Println("Failed to parse JSON:", err)
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
		})

		HandleMessage(form)
	}
}

func HandleMessage(event model.Event) {
	// pass services to every handler
	for _, service := range event.NotifyData.Body.Services {
		fmt.Println(service.ServiceID)
	}
	HandleParkFireProtection(event.NotifyData.Body.Services)
	HandleFireProtection(event.NotifyData.Body.Services)
	HandleAirConditioner(event.NotifyData.Body.Services)
	HandleParkLighting(event.NotifyData.Body.Services)
	HandleWindowControl(event.NotifyData.Body.Services)
	HandlePumpControl(event.NotifyData.Body.Services)
	HandleElectricCurrent(event.NotifyData.Body.Services)
}

// 2、室内灯光开关逻辑  白天关灯、晚上开灯，有人也开关,光照低于200开室内灯
func HandleParkLighting(services []model.Service) {
	fmt.Println("Handling park_lighting service")

	var flag bool
	// if time.Now().Hour() > 6 && time.Now().Hour() < 18 {
	// 	flag = false
	// } else {
	// 	flag = true
	// }

	for _, service := range services {
		if service.ServiceID == "light_switch_commands" {
			fmt.Println("External light is: ", *service.Properties.ExternalLight)
			if *service.Properties.Sunlight < 300 {
				flag = true
			} else {
				flag = false
			}
		}
		if service.ServiceID == "personnal_access" {
			fmt.Println("Human existence is: ", *service.Properties.HumanExistence)
			if *service.Properties.HumanExistence == "true" {
				ExsitCount = 0
			} else {
				ExsitCount++
			}
			if ExsitCount > 3 {
				flag = false
			}
		}
	}

	fmt.Println(flag)
	commandParams := map[string]interface{}{
		"indoor_light_switch": flag,
	}

	util.SendIoTCommand(HWClient, DeviceId, commandParams, "light_switch_commands", "park_lighting")
}

// 3、报警蜂鸣器逻辑 遇到火光就发送报警信号，有火光报警。
func HandleParkFireProtection(services []model.Service) {
	fmt.Println("Handling park_fire_protection service")

	var flag = false

	for _, service := range services {
		if service.ServiceID == "park_fire_protection" {
			fmt.Println("Fire is: ", *service.Properties.FireOccurrence)
			if *service.Properties.FireOccurrence == "true" {
				flag = true
			}
		}
	}

	fmt.Println(flag)
	commandParams := map[string]interface{}{
		"buzzer_switch": flag,
	}

	if flag {
		fmt.Println("Sending fire alarm")
	}

	util.SendIoTCommand(HWClient, DeviceId, commandParams, "park_fire_protection_commands", "park_fire_protection")
}

// 窗户开关逻辑，室内外温度判断，在24到28°的时候窗户处于打开状态，或者当空气环境质量较差时关闭窗户。如果有火灾发生，打开窗户
func HandleWindowControl(services []model.Service) {
	fmt.Println("Handling window_control service")

	var flag = false

	for _, service := range services {
		if service.ServiceID == "atmospheric_environment" {
			fmt.Println("Temperature is: ", *service.Properties.Temperature)
			if *service.Properties.Temperature > 24 && *service.Properties.Temperature < 28 {
				flag = true
			}
			if service.Properties.AirQuality != nil && *service.Properties.AirQuality < 60 {
				flag = false
			}
			if service.Properties.Rainfall != nil && *service.Properties.Rainfall == "true" {
				flag = false
			}
		}
		if service.ServiceID == "park_fire_protection" {
			fmt.Println("Fire is: ", *service.Properties.FireOccurrence)
			if *service.Properties.FireOccurrence == "true" {
				flag = true
			}
		}
		if service.ServiceID == "personnal_access" {
			if *service.Properties.HumanExistence == "true" {
				ExsitCount = 0
			} else {
				ExsitCount++
			}
			if ExsitCount > 3 {
				flag = false
			}
		}
		if service.ServiceID == "park_fire_protection" {
			fmt.Println("Fire is: ", *service.Properties.FireOccurrence)
			if *service.Properties.FireOccurrence == "true" {
				fmt.Println("It is a fire")
				flag = true
			}
		}
	}

	fmt.Println(flag)
	commandParams := map[string]interface{}{
		"window_switch": flag,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "atmospheric_environment_commands", "atmospheric_environment")

}

func HandleFireProtection(services []model.Service) {

	var flag = false

	for _, service := range services {
		if service.ServiceID == "park_fire_protection" {
			fmt.Println("Fire is: ", *service.Properties.FireOccurrence)
			if *service.Properties.FireOccurrence == "true" {
				fmt.Println("It is a fire")
				flag = true
			}
		}
	}

	if flag {
		commandParams := map[string]interface{}{
			"access_control_switch": flag,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "access_control_commands", "personnal_access")
		commandParams = map[string]interface{}{
			"humidifier_switch": flag,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "park_fire_protection_commands", "park_fire_protection")
	}

}

// 6、卧式水泵触发逻辑 当土壤湿度较低时，触发卧式水泵进行灌溉，当土壤湿度恢复到一定值时关闭卧式水泵
func HandlePumpControl(services []model.Service) {
	fmt.Println("Handling pump_control service")

	var flag bool

	for _, service := range services {
		if service.ServiceID == "atmospheric_environment" {
			fmt.Println("Humidity is: ", *service.Properties.Humidity)
			if *service.Properties.Humidity < 60 {
				flag = true
			}
			if *service.Properties.Humidity > 80 {
				flag = false
			}
		}
	}

	fmt.Println(flag)
	commandParams := map[string]interface{}{
		"pump_switch": flag,
	}

	util.SendIoTCommand(HWClient, DeviceId, commandParams, "atmospheric_environment_commands", "atmospheric_environment")

}

// 7、电流提醒逻辑，设备测上报电流，当电流过大时，在前端发送提醒警告电流过大
func HandleElectricCurrent(services []model.Service) {
	fmt.Println("Handling electric_current service")

	var flag = false

	for _, service := range services {
		if service.ServiceID == "park_energy" {
			fmt.Println("Electric current is: ", *service.Properties.ElectricCurrent)
			if *service.Properties.ElectricCurrent > 80 {
				flag = true
			}
		}
	}

	fmt.Println(flag)
	// 返回给前端。设成一个global flag

	ElectricCurrentFlag = flag
}

func HandleAirConditioner(services []model.Service) {
	fmt.Println("Handling atmospheric_environment service")

	// "service_id": "atmospheric_environment",
	// "properties": {
	//     "temperature": 80,
	//     "humidity": 80,
	//     "air_quility": 80,
	//     "rainfall": true
	// },

	//  temperature>29, air_conditioner_commands-air_conditioner_switch->true
	//  temperature<20, air_conditioner_commands-air_conditioner_switch->false

	//  humidity>80, air_conditioner_commands-air_conditione_switch->true
	//  humidity<60, air_conditioner_commands-air_conditione_switch->false

	var flag bool

	for _, service := range services {
		if service.ServiceID == "atmospheric_environment" {
			fmt.Println("Temperature is: ", *service.Properties.Temperature)
			if *service.Properties.Temperature > 29 {
				flag = true
			} else if *service.Properties.Temperature < 20 {
				flag = false
			}
			if *service.Properties.Humidity > 80 {
				flag = true
			} else if *service.Properties.Humidity < 60 {
				flag = false
			}
			if service.Properties.Rainfall != nil && *service.Properties.Rainfall == "true" {
				flag = true
			}
		}
		if service.ServiceID == "park_fire_protection" {
			fmt.Println("Fire is: ", *service.Properties.FireOccurrence)
			if *service.Properties.FireOccurrence == "true" {
				fmt.Println("It is a fire")
				flag = true
			}
		}
	}

	fmt.Println(flag)
	commandParams := map[string]interface{}{
		"air_conditioner_switch": flag,
	}

	if flag {
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "air_conditioner_commands", "park_energy")
	}
}

//	{
//	    "window_switch": true
//	}
func HandleSwitch(c *gin.Context) {
	var form model.Switch
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "wrong request",
		})
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
	fmt.Println(form)
	if HWClient == nil {
		SettingUpEnvironment()
		InitHuaweiCloudClient()
	}
	if form.WindowSwitch != nil {
		commandParams := map[string]interface{}{
			"window_switch": *form.WindowSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "atmospheric_environment_commands", "atmospheric_environment")
	}
	if form.PumpSwitch != nil {
		commandParams := map[string]interface{}{
			"pump_switch": *form.PumpSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "atmospheric_environment_commands", "atmospheric_environment")
	}
	if form.ExternalLightSwitch != nil {
		commandParams := map[string]interface{}{
			"external_light_switch": *form.ExternalLightSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "light_switch_commands", "park_lighting")
	}
	if form.IndoorLightSwitch != nil {
		commandParams := map[string]interface{}{
			"indoor_light_switch": *form.IndoorLightSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "light_switch_commands", "park_lighting")
	}
	if form.AccessControlSwitch != nil {
		commandParams := map[string]interface{}{
			"access_control_switch": *form.AccessControlSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "access_control_commands", "personnal_access")
	}
	if form.AirConditionerSwitch != nil {
		commandParams := map[string]interface{}{
			"air_conditioner_switch": *form.AirConditionerSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "air_conditioner_commands", "park_energy")
	}
	if form.BuzzerSwitch != nil {
		commandParams := map[string]interface{}{
			"buzzer_switch": *form.BuzzerSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "park_fire_protection_commands", "park_fire_protection")
	}
	if form.HumidifierSwitch != nil {
		commandParams := map[string]interface{}{
			"humidifier_switch": *form.HumidifierSwitch,
		}
		util.SendIoTCommand(HWClient, DeviceId, commandParams, "park_fire_protection_commands", "park_fire_protection")
	}

}

func InitializeSwitches() {
	if HWClient == nil {
		SettingUpEnvironment()
		InitHuaweiCloudClient()
	}
	commandParams := map[string]interface{}{
		"window_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "atmospheric_environment_commands", "atmospheric_environment")
	commandParams = map[string]interface{}{
		"pump_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "atmospheric_environment_commands", "atmospheric_environment")
	commandParams = map[string]interface{}{
		"external_light": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "light_switch_commands", "park_lighting")
	commandParams = map[string]interface{}{
		"indoor_light_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "light_switch_commands", "park_lighting")
	commandParams = map[string]interface{}{
		"access_control_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "access_control_commands", "personnal_access")
	commandParams = map[string]interface{}{
		"air_conditioner_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "air_conditioner_commands", "park_energy")
	commandParams = map[string]interface{}{
		"buzzer_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "park_fire_protection_commands", "park_fire_protection")
	commandParams = map[string]interface{}{
		"humidifier_switch": false,
	}
	util.SendIoTCommand(HWClient, DeviceId, commandParams, "park_fire_protection_commands", "park_fire_protection")
}
