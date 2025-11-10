package main

import (
	"fmt"
	"go-micro-todoList/app/task/repository/db/dao"
	_ "go-micro-todoList/app/task/service"
	"go-micro-todoList/config"
	_ "go-micro-todoList/idl/pb"

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
	// _ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserServ())
	_ = microService.Run()
}
