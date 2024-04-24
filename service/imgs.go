package service

import (
	handlers "Sgrid/src/http"
	"Sgrid/src/public"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ImgService(ctx *handlers.SgridServerCtx) {
	G := ctx.Engine.Group(strings.ToLower(ctx.Name))

	G.POST("/blogImg", func(c *gin.Context) {
		file, err := c.FormFile("file")
		u := uuid.New()
		file.Filename = "uid_" + u.String() + "_" + file.Filename
		path := public.Join("./imgs/" + file.Filename)
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
		path := public.Join("./imgs/")
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

	G.GET("/delPic", func(c *gin.Context) {
		s := c.Query("imgPath")
		fmt.Println("s", s)
		path := public.Join("./imgs/", s)

		err := os.Remove(path)
		if err != nil {
			c.AbortWithStatusJSON(200, handlers.Resp(-1, err.Error(), nil))
			return
		}
		c.AbortWithStatusJSON(200, handlers.Resp(0, "ok", nil))
	})
}
