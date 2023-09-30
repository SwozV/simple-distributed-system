package main

import (
	"context"
	stlog "log"
	"simple-distributed-system-swoz/log"
	"simple-distributed-system-swoz/service"
)

func main() {
	log.Run("./distributed.log")

	// 通常由配置文件或环境变量读取
	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()
}
