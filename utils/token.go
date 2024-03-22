package utils

import (
	"Simp/servers/BlogServer/obj/dao"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("2021121")

// 生成token的函数
func GenerateToken(username string, expiresAt time.Time) (string, error) {
	fmt.Println("username", username)
	// 设置token的过期时间
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expiresAt.Unix(),
	}

	// 使用HS256算法创建一个新的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名token，并获取字符串格式的结果
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// 验证用户登录函数
func AuthenticateUser(username, password string, user dao.User) bool {
	if user.Name == username {
		fmt.Println("user.password", user.Password)
		fmt.Println("password", password)
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
		if err == nil {
			return true
		}
	}
	return false
}
