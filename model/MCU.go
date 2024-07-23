package model

// // MCU 单片机
// type MCU struct {
// 	ID       int       `json:"id"`       // 单片机的 ID
// 	Mac      string    `json:"mac"`      // 单片机的 MAC 地址
// 	AuthCode string    `json:"authCode"` // 单片机的认证码
// 	Time     time.Time `json:"time"`     // 单片机的时间
// }

// // GetMCUByAuthCode 根据认证码获取单片机
// func (mcu *MCU) GetMCUByAuthCode(db *gorm.DB) (err error) {
// 	err = db.Where("auth_code = ?", mcu.AuthCode).First(mcu).Error
// 	return
// }

// // GetMCUByMac 根据 MAC 地址获取单片机
// func (mcu *MCU) GetMCUByMac(db *gorm.DB) (err error) {
// 	err = db.Where("mac = ?", mcu.Mac).First(mcu).Error
// 	return
// }

// // GetMCUByID 根据 ID 获取单片机
// func (mcu *MCU) GetMCUByID(db *gorm.DB) (err error) {
// 	err = db.Where("id = ?", mcu.ID).First(mcu).Error
// 	return
// }

// // Save 保存单片机
// func (mcu *MCU) Save(db *gorm.DB) (err error) {
// 	err = db.Create(mcu).Error
// 	return
// }

type Event struct {
	Resource    string     `json:"resource"`
	Event       string     `json:"event"`
	EventTime   string     `json:"event_time"`
	EventTimeMs string     `json:"event_time_ms"`
	RequestID   string     `json:"request_id"`
	NotifyData  NotifyData `json:"notify_data"`
	Body        Body       `json:"body"`
}

type NotifyData struct {
	Header Header `json:"header"`
}

type Header struct {
	DeviceID  string `json:"device_id"`
	ProductID string `json:"product_id"`
	AppID     string `json:"app_id"`
	GatewayID string `json:"gateway_id"`
	NodeID    string `json:"node_id"`
	Tags      []Tag  `json:"tags"`
}

type Tag struct {
	TagValue string `json:"tag_value"`
	TagKey   string `json:"tag_key"`
}

type Body struct {
	Services []Service `json:"services"`
}

type Service struct {
	ServiceID  string     `json:"service_id"`
	Properties Properties `json:"properties"`
	EventTime  string     `json:"event_time"`
}

type Properties struct {
	Temperature     *int  `json:"temperature,omitempty"`
	Humidity        *int  `json:"humidity,omitempty"`
	AirQuality      *int  `json:"air_quility,omitempty"`
	Rainfall        *bool `json:"rainfall,omitempty"`
	ElectricCurrent *int  `json:"electric_current,omitempty"`
	WaterDischarge  *int  `json:"water_discharge,omitempty"`
	Sunlight        *int  `json:"sunlight,omitempty"`
	ExternalLight   *bool `json:"external_light,omitempty"`
	IndoorLight     *bool `json:"indoor_light,omitempty"`
	HumanExistence  *bool `json:"human_existence,omitempty"`
	FireOccurrence  *bool `json:"fire_occurence,omitempty"`
}

type Switch struct {
	WindowSwitch          *bool `json:"window_switch,omitempty"`
	PumpSwitch            *bool `json:"pump_switch,omitempty"`
	AirConditionerSwitch  *bool `json:"air_conditioner_switch,omitempty"`
	IndoorLightSwitch     *bool `json:"indoor_light_switch,omitempty"`
	ExternalLightSwitch   *bool `json:"external_light_switch,omitempty"`
	AccessControlSwitch   *bool `json:"access_control_switch,omitempty"`
	BuzzerSwitch          *bool `json:"buzzer_switch,omitempty"`
	HumidifierSwitch      *bool `json:"humidifier_switch,omitempty"`
}
