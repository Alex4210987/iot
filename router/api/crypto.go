package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 结构体
type User struct {
	gorm.Model
	Username      string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password      string
	LastLoginTime time.Time
}

// 全局变量
var (
	dbConn      *gorm.DB
	onlineUsers []string
	onlineUser  string
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

	var user User
	if err := dbConn.Where("username = ?", form.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "msg": "用户名不存在"})
		return
	}

	if user.Password != Hash(form.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "msg": "密码错误"})
		return
	}

	// Update LastLoginTime
	user.LastLoginTime = time.Now()
	dbConn.Save(&user)

	// Add the username to the online users array
	onlineUsers = append(onlineUsers, form.Username)

	// Set the onlineUser variable to the logged-in user's username
	onlineUser = form.Username

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

	var user User
	if err := dbConn.Where("username = ?", form.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "msg": "用户名已存在"})
		return
	}

	hashedPassword := Hash(form.Password)
	newUser := User{
		Username:      form.Username,
		Password:      hashedPassword,
		LastLoginTime: time.Now(),
	}
	dbConn.Create(&newUser)

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

// InitDB 初始化数据库连接
func InitDB() {
	// read from .env
	password := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("root:%s@tcp(114.116.221.179)/iot?charset=utf8mb4&parseTime=True&loc=Local", password)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	dbConn = db

	// 自动迁移数据库，创建User表
	err = dbConn.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}
}