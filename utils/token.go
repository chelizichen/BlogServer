package utils

import (
	c "Simp/servers/BlogServer/configuration"
	"Simp/servers/BlogServer/obj/dao"
	"Simp/servers/BlogServer/obj/vo"
	handlers "Simp/src/http"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	TOKENS    = "tokens"
	USER_INFO = "userInfo"
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

func ValidateToken(values []string) (val string, err error) {
	for _, v := range values {
		s := GetTokenKey(v)
		if s != "" {
			return s, nil
		}
		continue
	}
	return "", nil
}

// hget tokens eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTExODA2MjAsInVzZXJuYW1lIjoiY2hlbGl6aWNoZW4ifQ.p1n7bUApEYbMRX2Hy_-OUN-z4Xap5iZ44qEf5rFwI1o
func SetTokenKey(Key string, expireTime time.Duration, val vo.UserVo) {
	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println("Error to Marshal", err.Error(), val)
		return
	}
	c.GRDB.SetEX(c.RDBContext, Key, b, expireTime)
}

func GetTokenKey(Key string) string {
	s, err := c.GRDB.Get(c.RDBContext, Key).Result()
	if err != nil {
		fmt.Println("get  key Error", err.Error())
		return ""
	}
	return s
}

func ValidateTokenMiddleware(c *gin.Context) {
	s := c.Request.Header["Token"]
	if len(s) == 0 {
		c.AbortWithStatusJSON(http.StatusOK, handlers.Resp(-1101, "token validate error", nil))
		return
	}
	value, err := ValidateToken(s)
	if value == "" || err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, handlers.Resp(-1102, "token validate error", err.Error()))
		} else {
			c.AbortWithStatusJSON(http.StatusOK, handlers.Resp(-1101, "token validate error", nil))
		}
	} else {
		c.Set(USER_INFO, value)
		c.Next()
	}
}

func Add12Hours() time.Duration {
	return time.Hour * 12
}
