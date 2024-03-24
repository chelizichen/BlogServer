package service

import (
	"Simp/servers/BlogServer/configuration"
	handlers "Simp/src/http"
)

func InitService(ctx *handlers.SimpHttpServerCtx, pre string) {
	configuration.InitStorage(*ctx)
	CreateUploadPath()
}
