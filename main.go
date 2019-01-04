package main

import (
	"fmt"
	"github.com/Deansquirrel/goMssqlDemo/common"
	"github.com/Deansquirrel/goMssqlDemo/dbOpr"
	"github.com/Deansquirrel/goMssqlDemo/global"
)

func main() {
	fmt.Println("程序启动")
	defer fmt.Println("程序退出")

	err := refreshConfig()
	if err != nil {
		common.PrintAndLog(err.Error())
		return
	}

	err = dbOpr.Select()
	if err != nil {
		common.PrintAndLog(err.Error())
	}

	err = dbOpr.MultipleCommand()
	if err != nil {
		common.PrintAndLog(err.Error())
	}

	err = dbOpr.TxCommand()
	if err != nil {
		common.PrintAndLog(err.Error())
	}
}

func refreshConfig() (err error) {
	//读取配置文件
	global.SysConfig, err = common.GetConfig("config.toml")
	if err != nil {
		common.PrintAndLog(err.Error())
		return
	}
	global.Conn, err = dbOpr.GetDbConn(global.SysConfig.MsSqlConfig.Server, global.SysConfig.MsSqlConfig.Port,
		global.SysConfig.MsSqlConfig.Database, global.SysConfig.MsSqlConfig.User, global.SysConfig.MsSqlConfig.Password)
	if err != nil {
		common.PrintAndLog(err.Error())
		return
	}

	return
}
