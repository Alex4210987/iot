package util

import (
	_const "backend/const"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken 创建 JWT
func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"nbf":      time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(_const.Secret_key))

	return tokenString, err
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (string, time.Time, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(_const.Secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		nbf := int64(claims["nbf"].(float64))

		return username, time.Unix(nbf, 0), nil
	} else {
		return "", time.Time{}, err
	}
}
