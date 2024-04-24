package service

import (
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"Sgrid/server/SubServer/BlogServer/storage"
	"Sgrid/server/SubServer/BlogServer/utils"
	handlers "Sgrid/src/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ColumnsService(ctx *handlers.SgridServerCtx) {
	G := ctx.Engine.Group(strings.ToLower(ctx.Name))
	G.GET("/getColumnDetail", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		res := storage.GetColumnDetail(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.GET("/saveArticleInColumn", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		cid, _ := strconv.Atoi(c.Query("cid"))
		res := storage.SaveArticleInColumn(id, cid)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.GET("/getColumns", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		pagination := utils.NewPagination(c)
		res := storage.GetColumnList(*pagination)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.POST("/saveColumn", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		var column *dao.Column
		c.BindJSON(&column)
		res := storage.SaveColumn(*column)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", res))
	})

	G.GET("/delColumnById", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		i := storage.DeleteColumnById(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", i))
	})
}
