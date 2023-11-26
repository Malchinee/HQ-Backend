package main

import (
	"HQ-Backend/common"
	"flag"
	"fmt"

	"HQ-Backend/api/internal/config"
	"HQ-Backend/api/internal/handler"
	"HQ-Backend/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(common.ConfigFile, &c)

	server := rest.MustNewServer(c.HTTP)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.HTTP.Host, c.HTTP.Port)
	server.Start()
}
