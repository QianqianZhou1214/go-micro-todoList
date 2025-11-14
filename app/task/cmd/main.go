package main

import (
	"context"
	"fmt"
	"go-micro-todoList/app/task/repository/db/dao"
	"go-micro-todoList/app/task/script"
	"go-micro-todoList/app/task/service"
	"go-micro-todoList/idl/pb"

	"go-micro-todoList/config"

	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()

	etcdReg := registry.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// new an instance of micro
	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address(config.TaskServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskServ())
	_ = microService.Run()
}

func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
