package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"simple-distributed-system-swoz/registry"
)

func Start(ctx context.Context, reg registry.Registration,
	host, port string,
	registerHandlersFunc func()) (context.Context, error) {

	registerHandlersFunc()

	ctx = startService(ctx, reg.ServiceName, host, port)

	// 启动注册服务
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, nil
	}

	return ctx, nil
}

func startService(ctx context.Context,
	serviceName registry.ServiceName,
	host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v start. Press any key to stop .\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
