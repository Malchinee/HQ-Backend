package main

import (
	"github.com/Malchinee/HQ-Backend/api"
	"github.com/zeromicro/go-zero/core/service"
)

func main() {
	group := service.NewServiceGroup()
	defer group.Stop()

	apiServer := api.InitTrainApiServer()

	group.Add(apiServer)

	group.Start()
}
