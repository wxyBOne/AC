package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims 自定义声明结构体
type CustomClaims struct {
	UserID   uint   `json:"user_id"` // 添加用户 ID
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtSecret = []byte("0xF2038d6fD44dc9A980502FacD92c761020402af5") // 替换为您的密钥

// GenerateToken 生成 JWT
func GenerateToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour) // 过期时间为 2 小时
	claims := &CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken 验证 JWT
func ValidateToken(tokenString string) (uint, string, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return 0, "", err
	}

	return claims.UserID, claims.Username, nil // 返回用户
}

// GenerateAdminToken 生成管理员JWT token
func GenerateAdminToken(adminID uint, adminname string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id":  adminID,
		"adminname": adminname,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 24小时过期
	})

	return token.SignedString([]byte("your_admin_secret_key"))
}
