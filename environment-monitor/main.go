package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// User 结构体
type User struct {
	Username      string
	Password      string
	LastLoginTime time.Time
}

// 全局变量
var (
	onlineUsers = map[string]time.Time{}
	onlineUser  = "" // 用于记录当前登录的用户
	users       = map[string]User{
		"admin": {Username: "admin", Password: Hash("admin")},
		"user":  {Username: "user", Password: Hash("userpasswd")},
	}
)

// Hash 对密码进行哈希
func Hash(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// UserLogin 处理用户登录
func UserLogin(c *gin.Context) {
	var form struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "请求参数错误"})
		return
	}

	user, exists := users[form.Username]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "msg": "用户名不存在"})
		return
	}

	if user.Password != Hash(form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "msg": "密码错误"})
		return
	}

	user.LastLoginTime = time.Now()
	users[form.Username] = user
	onlineUsers[form.Username] = time.Now() // 记录在线用户
	onlineUser = form.Username              // 记录当前登录用户

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "登录成功"})
}

// UserRegister 处理用户注册
func UserRegister(c *gin.Context) {
	var form struct {
		Username       string `json:"username" binding:"required"`
		Password       string `json:"password" binding:"required"`
		PasswordRepeat string `json:"password_repeat" binding:"required"`
	}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "请求参数错误"})
		return
	}

	if form.Password != form.PasswordRepeat {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "密码和确认密码不一致"})
		return
	}

	if _, exists := users[form.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "用户名已存在"})
		return
	}

	users[form.Username] = User{
		Username: form.Username,
		Password: Hash(form.Password),
	}

	// 在注册时直接将用户标记为在线
	onlineUsers[form.Username] = time.Now()
	onlineUser = form.Username // 记录当前登录用户

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "注册成功"})
}

// GetOnlineUsers 查询在线用户
func GetOnlineUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"online_users": onlineUsers})
}

// GetCurrentUser 查询当前登录用户
func GetCurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"current_user": onlineUser})
}

func main() {
	router := gin.Default()

	// CORS 配置
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},         // 允许的源，配置你的前端地址
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},        // 允许的 HTTP 方法
		AllowHeaders:     []string{"Authorization", "Content-Type"}, // 允许的头部
		AllowCredentials: true,                                      // 是否允许客户端发送 Cookie
	}))

	router.POST("/login", UserLogin)
	router.POST("/register", UserRegister)
	router.GET("/online-users", GetOnlineUsers) // 查询在线用户的 API
	router.GET("/current-user", GetCurrentUser) // 查询当前用户的 API

	router.Run(":3070")
}
