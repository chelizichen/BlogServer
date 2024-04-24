package service

import (
	"Sgrid/server/SubServer/BlogServer/configuration"
	handlers "Sgrid/src/http"
	"Sgrid/src/public"
	"fmt"
)

func InitService(ctx *handlers.SgridServerCtx) {
	sc, err := public.NewConfig()
	if err != nil {
		fmt.Println("Error To NewConfig", err)
	}
	configuration.InitStorage(sc)
	CreateUploadPath()
	ctx.RegistryHttpRouter(LoginService)
	ctx.RegistryHttpRouter(ArticleService)
	ctx.RegistryHttpRouter(ColumnsService)
	ctx.RegistryHttpRouter(ImgService)
	ctx.RegistryHttpRouter(UploadService)
	ctx.RegistryHttpRouter(UserService)
}
