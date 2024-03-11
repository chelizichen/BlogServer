package service

import (
	"Simp/servers/BlogServer/configuration"
	"Simp/servers/BlogServer/storage"
	"Simp/servers/BlogServer/utils"
	handlers "Simp/src/http"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BlogService(ctx *handlers.SimpHttpServerCtx, pre string) {
	configuration.InitStorage(*ctx)

	E := ctx.Engine
	G := E.Group(pre)
	G.POST("/blogImg", func(c *gin.Context) {
		file, err := c.FormFile("file")
		u := uuid.New()
		file.Filename = "uid_" + u.String() + "_" + file.Filename
		path := utils.JoinPathUtil("./imgs/" + file.Filename)
		if err != nil {
			resp := make(map[string]interface{})
			resp["errno"] = -1
			resp["message"] = err.Error()
			c.AbortWithStatusJSON(200, resp)
			return
		}
		c.SaveUploadedFile(file, path)
		data := make(map[string]interface{})
		resp := make(map[string]interface{})
		resp["errno"] = 0
		resp["data"] = data
		SIMP_PRODUCTION := os.Getenv("SIMP_PRODUCTION")
		if SIMP_PRODUCTION == "Yes" {
			data["url"] = "/blogserver/imgs/" + file.Filename
		} else {
			data["url"] = "/blogserver/imgs/" + file.Filename
		}
		c.AbortWithStatusJSON(200, resp)
	})

	G.GET("/getPics", func(c *gin.Context) {
		path := utils.JoinPathUtil("./imgs/")
		de, err := os.ReadDir(path)
		if err != nil {
			c.AbortWithStatusJSON(0, handlers.Resp(-1, err.Error(), err))
		}
		pics := make([]map[string]interface{}, 0)
		for _, v := range de {
			resp := make(map[string]interface{})
			resp["name"] = v.Name()
			resp["path"] = "/blogserver/imgs/" + v.Name()
			pics = append(pics, resp)
		}
		c.AbortWithStatusJSON(0, handlers.Resp(0, "ok", pics))
	})
	G.GET("/getArticleList", func(c *gin.Context) {
		pagination := utils.NewPagination(c)
		resp := storage.GetArticleList(pagination.Offset, pagination.Size, pagination.Keyword)
		// fmt.Println("list", list.Scan())
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", resp))
	})

	G.POST("/saveArticle", func(c *gin.Context) {
		var article *configuration.Article
		c.BindJSON(&article)
		fmt.Println("article", article.ID)
		r := storage.SaveArticle(*article)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", r))
	})

	G.GET("/getArticleById", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		r := storage.GetArticle(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", r))
	})

	G.GET("/delArticleById", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		storage.DelArticleById(id)
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", nil))
	})

	E.Use(G.Handlers...)
}
