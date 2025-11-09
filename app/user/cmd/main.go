package main

import (
	"fmt"
	"go-micro-todoList/app/user/repository/db/dao"
	"go-micro-todoList/app/user/service"
	"go-micro-todoList/config"
	"go-micro-todoList/idl/pb"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()

	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	/*
		fmt.Println("Connecting to etcd at:", fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort))

		if err := etcdReg.Init(); err != nil {
			fmt.Println("Etcd registry init error:", err)
		} else {
			fmt.Println("Etcd registry initialized successfully")
		}

	*/

	// new an instance of micro
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserServ())
	_ = microService.Run()
}
