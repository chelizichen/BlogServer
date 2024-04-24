package service

import (
	"Sgrid/server/SubServer/BlogServer/obj/dao"
	"Sgrid/server/SubServer/BlogServer/storage"
	"Sgrid/server/SubServer/BlogServer/utils"
	handlers "Sgrid/src/http"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ArticleService(ctx *handlers.SgridServerCtx) {
	G := ctx.Engine.Group(strings.ToLower(ctx.Name))
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
}
