package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

//配置参数
var (
	AppMode  string
	HttpPort string
	Db       string
	DbHost   string
	Dbport   string
	DbUser   string
	DbName   string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("Appmode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").String()
	Dbport = file.Section("database").Key("Dbport").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbName = file.Section("database").Key("DbName").String()
}
