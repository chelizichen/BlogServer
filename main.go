package main

import (
	"Sgrid/server/SubServer/BlogServer/service"
	h "Sgrid/src/http"
	"Sgrid/src/public"
	"fmt"
)

func main() {
	ctx := h.NewSgridServerCtx(
		h.WithSgridServerType(public.PROTOCOL_HTTP),
		h.WithSgridGinStatic([2]string{"/web", "dist"}),
		h.WithCors(),
	)
	ctx.RegistryHttpRouter(service.InitService)

	ctx.Static("/imgs", "imgs")
	ctx.Static("/uploadPackages", "uploadFile")
	h.NewSgridServer(ctx, func(port string) {
		ctx.Engine.Run(port)
		fmt.Println("Server started on " + port)
	})
}
