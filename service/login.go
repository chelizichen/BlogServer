package service

import (
	"Simp/servers/BlogServer/obj/dto"
	"Simp/servers/BlogServer/obj/vo"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var userCache = make(map[string]string)

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

		userVo := &vo.UserVo{
			Name:       r.Name,
			Password:   "",
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
			userCache[string(rune(r.ID))] = token
			userVo.Token = userCache[string(rune(r.ID))]
			c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", userVo))
			return
		}
		c.AbortWithStatusJSON(200, handlers.Resp(-1, "error", err.Error()))
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
