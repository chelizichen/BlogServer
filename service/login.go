package service

import (
	"Simp/servers/BlogServer/obj/dto"
	"Simp/servers/BlogServer/obj/vo"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

var userCache = make(map[string]string)

func CheckToken(token string) bool {
	if token == "" {
		return false
	}
	s := strings.Split(token, "|")
	name := s[0]
	cacheToken := s[1]
	fmt.Println("name", name)
	fmt.Println("cacheToken", cacheToken)
	fmt.Println(" userCache[name]", userCache[name])
	return userCache[name] == cacheToken
}

func ValidateTokenMiddleware(c *gin.Context) {
	s := c.Request.Header["Token"]
	fmt.Println("s", s)
	if len(s) == 0 {
		c.AbortWithStatusJSON(http.StatusOK, handlers.Resp(-1101, "token validate error", nil))
		return
	}
	tkn := s[0]
	if CheckToken(tkn) {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusOK, handlers.Resp(-1101, "token validate error", nil))
	}
}

func LoginService(ctx *handlers.SimpHttpServerCtx, pre string) {
	E := ctx.Engine
	G := E.Group(pre)

	G.POST("/login", func(c *gin.Context) {
		var user *dto.UserDto
		err := c.BindJSON(&user)
		if err != nil {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "bind json error", nil))
			return
		}
		r, err := storage.LoginUser(*user)
		if err != nil {
			fmt.Println("LoginUserError", err.Error())
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "LoginUser Error", nil))
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //加密处理
		userVo := &vo.UserVo{
			Name:       r.Name,
			Password:   string(hash),
			CreateTime: r.CreateTime,
			Level:      r.Level,
			ID:         r.ID,
		}

		if err == nil {
			token, err := utils.GenerateToken(r.Name, time.Now().Add(time.Hour*12))
			if err != nil {
				c.AbortWithStatusJSON(200, handlers.Resp(-2, "error", err.Error()))
				return
			}
			userCache[r.Name] = token
			userVo.Token = token
			c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", userVo))
			return
		}
		c.AbortWithStatusJSON(200, handlers.Resp(-1, "error", err.Error()))
	})

	G.POST("/loginByCache", func(c *gin.Context) {

		var user *dto.UserDto
		err := c.BindJSON(&user)
		if err != nil {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "bind json error", nil))
			return
		}
		if userCache[user.Name] == "" {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "token validate error", nil))
			return
		}
		fmt.Println("user", user)
		r, err := storage.LoginByCache(*user)

		if err != nil {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "LoginByCache Error", nil))
			return
		}
		userVo := &vo.UserVo{
			Name:       r.Name,
			Password:   user.Password,
			CreateTime: r.CreateTime,
			Level:      r.Level,
			ID:         r.ID,
		}
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", userVo))
	})

	G.POST("/syncCache", func(c *gin.Context) {
		fmt.Println("c.Request.Host", c.Request.RemoteAddr)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", userCache))
	})

	G.POST("/getCache", func(c *gin.Context) {
		s := c.Query("user")
		if s != "" {
			c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", nil))
			return
		}
		fmt.Println("c.Request.Host", c.Request.RemoteAddr)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", userCache[s]))
	})
	E.Use(G.Handlers...)
}
