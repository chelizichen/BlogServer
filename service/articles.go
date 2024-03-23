package service

import (
	"Simp/servers/BlogServer/obj/dao"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ArticleService(ctx *handlers.SimpHttpServerCtx, pre string) {
	E := ctx.Engine
	G := E.Group(pre)
	G.GET("/getArticleList", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		pagination := utils.NewPagination(c)
		resp := storage.GetArticleList(pagination.Offset, pagination.Size, pagination.Keyword)
		// fmt.Println("list", list.Scan())
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", resp))
	})

	G.POST("/saveArticle", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		var article *dao.Article
		c.BindJSON(&article)
		fmt.Println("article", article.ID)
		r := storage.SaveArticle(*article)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", r))
	})

	G.GET("/getArticleById", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		r := storage.GetArticle(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", r))
	})

	G.GET("/delArticleById", utils.ValidateTokenMiddleware, func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		storage.DelArticleById(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", nil))
	})

	E.Use(G.Handlers...)
}
