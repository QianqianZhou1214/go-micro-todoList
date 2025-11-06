package cmd

import (
	"go-micro-todoList/app/user/repository/db/dao"
	"go-micro-todoList/config"
)

func main() {
	config.Init()
	dao.InitDB()
}
