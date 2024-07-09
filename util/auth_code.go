package util

import (
	"math/rand"
	"time"
)

// GenerateAuthCode 生成单片机的认证码
func GenerateAuthCode() string {
	// 生成 8 位随机字符串
	rand.Seed(time.Now().UnixNano())
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
