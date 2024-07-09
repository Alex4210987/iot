package api

import (
	"backend/client/database"
	"backend/http_param"
	"backend/model"
	"backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

// UserLogin 用户登录
func UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求表单
		var form http_param.UserLoginInfo
		err := c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "请求参数错误",
			})
			return
		}
		// 通过用户名获取用户
		var user model.User
		user.Username = form.Username
		err = user.GetUserByUsername(database.DbEngine)
		if err != nil {
			// 没有找到用户
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "用户名不存在",
			})
			return
		}
		// 对表单中的密码进行hash
		password := util.Hash(form.Password)
		// 找到用户，验证密码
		if user.Password == password {
			// 密码正确, 生成一个 token
			token, err := util.CreateToken(user.Username)
			if err != nil {
				c.JSON(200, gin.H{
					"status": "error",
					"msg":    "生成 token 失败",
				})
				return
			} else {
				// 刷新最后登录时间
				user.LastLoginTime = time.Now()
				err = user.Save(database.DbEngine)
				if err != nil {
					fmt.Println(err)
				}
				// 返回 token
				c.JSON(200, gin.H{
					"status": "success",
					"token":  token,
				})
				return
			}
		} else {
			// 密码错误
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "密码错误",
			})
			return
		}
	}
}

// UserRegister 用户注册
func UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求表单
		var form http_param.UserRegisterInfo
		err := c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "请求参数错误",
			})
			return
		}
		// 检查密码和确认密码是否一致
		if form.Password != form.PasswordRepeat {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "密码和确认密码不一致",
			})
			return
		}
		// 通过用户名获取用户
		var user model.User
		user.Username = form.Username
		err = user.GetUserByUsername(database.DbEngine)
		if err == nil {
			// 用户已存在
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "用户名已存在",
			})
			return
		}
		// 用户不存在，创建用户
		user.Password = util.Hash(form.Password)
		// 创建用户
		err = user.Save(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "创建用户失败",
			})
			return
		}
		// 创建用户成功, 生成一个 token
		token, err := util.CreateToken(user.Username)
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
	}
}

// UserBindMCU 绑定 MCU
func UserBindMCU() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求表单
		var form http_param.UserBindMCUInfo
		err := c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "请求参数错误",
			})
			return
		}
		// 寻找是否有对应MCU
		mcu := model.MCU{AuthCode: form.AuthCode}
		err = mcu.GetMCUByAuthCode(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "没有找到对应的MCU",
			})
			return
		}
		// 将MCU绑定到用户
		value, _ := c.Get("user")
		user := value.(model.User)
		userMcu := model.UserMCU{Name: form.Name, McuID: mcu.ID, UserID: user.ID}
		err = userMcu.Save(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "绑定失败",
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "success",
			"msg":    "绑定成功",
		})
	}
}

// UserHistoryData 获取单片机历史数据
func UserHistoryData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取userMCU
		userMcuId, err := strconv.Atoi(c.Query("user_mcu_id"))
		userMcu := model.UserMCU{ID: userMcuId}
		err = userMcu.GetUserMCUByID(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "没有找到对应的单片机",
			})
			return
		}
		// 获取用户
		value, _ := c.Get("user")
		user := value.(model.User)
		// 查看用户是否有该单片机
		if user.ID != userMcu.UserID {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "没有找到对应的单片机",
			})
			return
		}

		// 获取单片机所有的历史数据
		data := model.Data{McuID: userMcu.McuID}
		historyDatas, err := data.GetDatasByMCUID(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "获取建议失败",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
			"name":   userMcu.Name,
			"data":   historyDatas,
		})
	}
}

type UserMCUListResutl struct {
	ID     int       `json:"id"`      // 主键
	UserID int       `json:"user_id"` // 用户 ID
	McuID  int       `json:"mcu_id"`  // 单片机 ID
	Name   string    `json:"name"`    // 单片机的名字
	Mac    string    `json:"mac"`     // 单片机的 mac 地址
	Status bool      `json:"status"`  // 单片机的状态
	Time   time.Time `json:"time"`    // 单片机的时间
}

// UserMCUList 获取用户绑定的单片机列表
func UserMCUList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户
		value, _ := c.Get("user")
		user := value.(model.User)
		// 获取用户绑定的单片机列表
		userMcu := model.UserMCU{UserID: user.ID}
		userMcuList, err := userMcu.GetUserMCUListByUserID(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "没有找到对应的单片机",
			})
			return
		}
		var userMcuListResutl []UserMCUListResutl
		fmt.Println(userMcuList)
		// 获取单片机的mac地址
		for i, userMcu := range userMcuList {
			mcu := model.MCU{ID: userMcu.McuID}
			err := mcu.GetMCUByID(database.DbEngine)
			if err != nil {
				c.JSON(200, gin.H{
					"status": "error",
					"msg":    "没有找到对应的单片机",
				})
				return
			}
			userMcuListResutl = append(userMcuListResutl, UserMCUListResutl{
				ID:     userMcu.ID,
				UserID: userMcu.UserID,
				McuID:  userMcu.McuID,
				Name:   userMcu.Name,
				Mac:    mcu.Mac,
				Time:   mcu.Time,
			})
			// 获取单片机最近的一条数据
			data := model.Data{McuID: userMcu.McuID}
			err = data.GetLastDataByMCUID(database.DbEngine)
			// 如果时间超过30分钟，就认为单片机离线
			if err != nil || time.Now().Sub(data.Timestamp).Minutes() > 30 {
				userMcuListResutl[i].Status = false
			} else {
				userMcuListResutl[i].Status = true
			}
			if err != nil {
				c.JSON(200, gin.H{
					"status": "error",
					"msg":    "没有找到对应的单片机",
				})
				return
			}
		}
		c.JSON(200, gin.H{
			"status": "success",
			"data":   userMcuListResutl,
		})
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UserRealTimeDataResult struct {
	ID            int       `json:"id"`
	McuID         int       `json:"mcu_id"`
	Humidity      float32   `json:"humidity"`
	Temperature   float32   `json:"temperature"`
	Raindrops     int32     `json:"precipitation"`
	EarthHumidity float32   `json:"soilHumidity"`
	BootTime      uint32    `json:"boot_time"`
	Timestamp     time.Time `json:"timestamp"`
	GptContent    string    `json:"gptContent"`
}

// UserRealTimeData 获取单片机实时数据
func UserRealTimeData() gin.HandlerFunc {
	return func(c *gin.Context) {
		func(w http.ResponseWriter, r *http.Request) {
			// 升级为 websocket
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Println(err)
				return
			}
			// 获取UserMCU和token
			userMcuID, _ := strconv.Atoi(c.Query("user_mcu_id"))
			token := c.Query("token")
			// 解析token
			username, Time, err := util.ParseToken(token)
			// 判断token是否有效
			if err != nil {
				log.Println(err)
				conn.Close()
				return
			}
			// 判断token是否超过7天
			if time.Now().Unix()-Time.Unix() > 60*60*24*7 {
				log.Println("token过期")
				conn.Close()
				return
			}
			// 获取用户
			user := model.User{Username: username}
			err = user.GetUserByUsername(database.DbEngine)
			if err != nil {
				log.Println(err)
				conn.Close()
				return
			}
			// 获取用户单片机
			userMcu := model.UserMCU{ID: userMcuID}
			err = userMcu.GetUserMCUByID(database.DbEngine)
			if err != nil {
				log.Println(err)
				conn.Close()
				return
			}

			// 判断用户是否有该单片机
			if user.ID != userMcu.UserID {
				log.Println("用户没有该单片机")
				conn.Close()
				return
			}

			// 获取单片机
			mcu := model.MCU{ID: userMcu.McuID}
			err = mcu.GetMCUByID(database.DbEngine)
			if err != nil {
				log.Println(err)
				conn.Close()
				return
			}

			// 不断发送最新的数据
			go func() {
				defer conn.Close()
				for {
					// 当有新数据时，发送给客户端
					select {
					case data := <-model.McuDataMapChan[mcu.ID]:
						var result UserRealTimeDataResult
						result.ID = data.ID
						result.McuID = data.McuID
						result.Humidity = data.Humidity
						result.Temperature = data.Temperature
						result.Raindrops = data.Raindrops
						result.EarthHumidity = data.EarthHumidity
						result.BootTime = data.BootTime
						result.Timestamp = data.Timestamp
						result.GptContent = ""
						// 将data转化为JSON并发送给客户端
						err := conn.WriteJSON(result)
						if err != nil {
							log.Println(err)
							return
						}
					}
				}
			}()
		}(c.Writer, c.Request)
	}
}

type UserListResutl struct {
	Username      string    `json:"username"`
	Status        bool      `json:"status"`
	LastLoginTime time.Time `json:"last_login_time"`
}

// UserList 获取用户列表
func UserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户
		value, _ := c.Get("user")
		user := value.(model.User)
		// 获取用户列表
		userList, err := user.GetUserList(database.DbEngine)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "没有找到对应的用户",
			})
			return
		}
		var userListResutl []UserListResutl
		for _, user := range userList {
			// 如果用户上次登录时间超过1h，就认为用户离线
			var status bool
			if time.Now().Sub(user.LastLoginTime).Minutes() > 60 {
				status = false
			} else {
				status = true
			}
			userListResutl = append(userListResutl, UserListResutl{
				Username:      user.Username,
				LastLoginTime: user.LastLoginTime,
				Status:        status,
			})
		}

		c.JSON(200, gin.H{
			"status": "success",
			"data":   userListResutl,
		})
	}
}

// Username 获取用户名
func Username() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, _ := c.Get("user")
		user := value.(model.User)
		c.JSON(200, gin.H{
			"status": "success",
			"data":   user.Username,
		})
	}
}

// GPT 获取GPT
func GPT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数
		var form http_param.GPTInfo
		err := c.ShouldBind(&form)
		if err != nil {
			c.JSON(200, gin.H{
				"status": "error",
				"msg":    "参数错误",
			})
			return
		}
		gptStr := fmt.Sprintf("温度：%v 湿度：%v 土壤湿度：%v 降雨量：%v", form.Temperature, form.Humidity, form.EarthHumidity, form.Raindrops)
		GptContent, err := util.GetAdvice(gptStr)
		if err != nil {
			GptContent = err.Error()
		}
		c.JSON(200, gin.H{
			"status": "success",
			"data":   GptContent,
		})
	}
}
