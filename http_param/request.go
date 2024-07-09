package http_param

// UserLoginInfo 用户登录信息
type UserLoginInfo struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// MCUAuthInfo 单片机认证信息
type MCUAuthInfo struct {
	Mac      string `json:"mac"`       // 单片机的 MAC 地址
	AuthCode string `json:"auth_code"` // 单片机的认证码
}

// MCUData 单片机数据
type MCUData struct {
	Humidity      float32 `json:"humidity"`       // 湿度
	Temperature   float32 `json:"temperature"`    // 温度
	Raindrops     int32   `json:"raindrops"`      // 雨滴
	EarthHumidity float32 `json:"earth_humidity"` // 土壤湿度
	TimeStamp     uint32  `json:"time_stamp"`     // 时间戳
}

// UserRegisterInfo 用户注册信息
type UserRegisterInfo struct {
	Username       string `json:"username"`        // 用户名
	Password       string `json:"password"`        // 密码
	PasswordRepeat string `json:"password_repeat"` // 重复密码
}

// UserBindMCUInfo 用户绑定单片机信息
type UserBindMCUInfo struct {
	Name     string `json:"name"`      // 单片机的名字
	AuthCode string `json:"auth_code"` // 单片机的认证码
}

// GPTInfo GPT信息
type GPTInfo struct {
	Humidity      float32 `form:"humidity"`      // 湿度
	Temperature   float32 `form:"temperature"`   // 温度
	Raindrops     int32   `form:"precipitation"` // 雨滴
	EarthHumidity float32 `form:"soilHumidity"`  // 土壤湿度
}
