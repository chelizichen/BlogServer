package service

import (
	"Simp/servers/BlogServer/obj/dao"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ColumnsService(ctx *handlers.SimpHttpServerCtx, pre string) {
	E := ctx.Engine
	G := E.Group(pre)

	G.GET("/getColumnDetail", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		res := storage.GetColumnDetail(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.GET("/saveArticleInColumn", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		cid, _ := strconv.Atoi(c.Query("cid"))
		res := storage.SaveArticleInColumn(id, cid)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.GET("/getColumns", func(c *gin.Context) {
		pagination := utils.NewPagination(c)
		res := storage.GetColumnList(*pagination)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.POST("/saveColumn", func(c *gin.Context) {
		var column *dao.Column
		c.BindJSON(&column)
		res := storage.SaveColumn(*column)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.GET("/delColumnById", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		i := storage.DeleteColumnById(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", i))
	})
	E.Use(G.Handlers...)
}
