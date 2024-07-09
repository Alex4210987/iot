package _const

import "os"

var Secret_key string

// InitConst 初始化常量
func InitConst() {
	Secret_key = os.Getenv("FM_SECRET_KEY")
}
