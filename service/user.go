package service

import (
	"Simp/servers/BlogServer/obj/dto"
	"Simp/servers/BlogServer/obj/vo"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"

	"github.com/gin-gonic/gin"
)

func UserService(ctx *handlers.SimpHttpServerCtx, pre string) {
	E := ctx.Engine
	G := E.Group(pre)
	G.GET("/getUserList", func(c *gin.Context) {
		pagination := utils.NewPagination(c)
		resp, tt, err := storage.GetUserList(*pagination)
		if err != nil {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "getUserList error", nil))
			return
		}
		m := make(map[string]interface{})
		userVoList := make([]*vo.UserVo, len(resp))
		for _, v := range resp {
			userVoList = append(userVoList, &vo.UserVo{
				Name:       v.Name,
				Level:      v.Level,
				Password:   v.Password,
				CreateTime: v.CreateTime,
				ID:         v.ID,
			})
		}
		m["list"] = userVoList
		m["total"] = tt
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", m))
	})

	G.POST("/changeUserLevel", func(c *gin.Context) {
		var lvl *dto.ChangeLevelDto
		err := c.BindJSON(&lvl)
		if err != nil {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, "bind json error", nil))
			return
		}
		storage.ChangeUserLevel(*lvl)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", nil))
	})
	E.Use(G.Handlers...)

}
