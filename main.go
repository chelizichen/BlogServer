package main

import (
	"Simp/servers/BlogServer/service"
	h "Simp/src/http"
)

func main() {
	ctx := h.NewSimpHttpCtx("simp.yaml")
	ctx.Use(service.LoginService)
	ctx.Use(service.InitService)
	ctx.Use(service.ArticleService)
	ctx.Use(service.ColumnsService)
	ctx.Use(service.ImgService)
	ctx.Use(service.UploadService)
	ctx.UseSPA("/web", "dist")
	ctx.Static("/imgs", "imgs")
	ctx.Static("/uploadPackages", "uploadFile")
	h.NewSimpHttpServer(ctx)
}
