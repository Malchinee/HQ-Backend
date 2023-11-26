package api

import (
	"fmt"
	"github.com/Malchinee/HQ-Backend/common"

	"github.com/Malchinee/HQ-Backend/api/internal/config"
	"github.com/Malchinee/HQ-Backend/api/internal/handler"
	"github.com/Malchinee/HQ-Backend/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func InitTrainApiServer() *rest.Server {
	var c config.Config
	conf.MustLoad(common.ConfigFile, &c)

	server := rest.MustNewServer(c.HTTP)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.HTTP.Host, c.HTTP.Port)

	return server
}
