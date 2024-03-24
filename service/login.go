package service

import (
	"Simp/servers/BlogServer/obj/dto"
	"Simp/servers/BlogServer/obj/vo"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func LoginService(ctx *handlers.SimpHttpServerCtx, pre string) {
	E := ctx.Engine
	G := E.Group(pre)
	G.POST("/loginByCache", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		value, exists := c.Get(utils.USER_INFO)
		if !exists {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "LoginUser Error", nil))
		} else {
			var vo *vo.UserVo
			json.Unmarshal([]byte(value.(string)), &vo)
			c.AbortWithStatusJSON(200, handlers.Resp(0, "Ok", vo))
		}
	})
	G.GET("/logout", func(c *gin.Context) {
		s := c.Query("token")
		i := utils.DelTokenKey(s)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "OK", i))
	})
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
		if err == nil {
			H12 := utils.Add12Hours()
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //加密处理
			if err != nil {
				fmt.Println("GenerateFromPasswordError", err.Error())
				c.AbortWithStatusJSON(200, handlers.Resp(-1, "GenerateFromPasswordError", nil))
				return
			}
			userVo := &vo.UserVo{
				Name:       r.Name,
				Password:   string(hash),
				CreateTime: r.CreateTime,
				Level:      r.Level,
				ID:         r.ID,
			}
			token, err := utils.GenerateToken(r.Name, time.Now().Add(H12))
			utils.SetTokenKey(token, H12, *userVo)
			if err != nil {
				c.AbortWithStatusJSON(200, handlers.Resp(-2, "error", err.Error()))
				return
			}
			userVo.Token = token
			c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", userVo))
			return
		}
		c.AbortWithStatusJSON(200, handlers.Resp(-1, "error", err.Error()))
	})
	E.Use(G.Handlers...)
}
