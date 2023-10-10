package main

import (
	"context"
	"fmt"
	stlog "log"
	"simple-distributed-system-swoz/log"
	"simple-distributed-system-swoz/registry"
	"simple-distributed-system-swoz/service"
)

func main() {
	log.Run("./distributed.log")

	// 通常由配置文件或环境变量读取
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
	}

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()
}
