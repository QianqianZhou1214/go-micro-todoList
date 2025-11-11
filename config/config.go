package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Charset    string

	EtcdHost string
	EtcdPort string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitHost       string
	RabbitPort       string

	UserServiceAddress string
	TaskServiceAddress string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("config.ini not found", err)
	}
	LoadMySqlData(file)
	LoadEtcd(file)
	LoadRabbitMq(file)
	LoadServer(file)
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadRabbitMq(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassword = file.Section("rabbitmq").Key("RabbitMQPassword").String()
	RabbitHost = file.Section("rabbitmq").Key("RabbitHost").String()
	RabbitPort = file.Section("rabbitmq").Key("RabbitPort").String()
}

func LoadServer(file *ini.File) {
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}

func LoadMySqlData(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}
