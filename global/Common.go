package global

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/kataras/iris/core/errors"
)

func GetConfig(fileName string) (config SysConfig, err error) {
	currPath, err := go_tool.GetCurrPath()
	if err != nil {
		currPath = ""
	} else {
		currPath = currPath + "\\" + fileName
	}
	b, err := go_tool.PathExists(currPath)
	if err != nil {
		err = errors.New("配置文件读取失败 - " + err.Error())
		return
	}
	if !b {
		err = errors.New("未发现配置文件")
		config.PrintFormat()
		return
	}
	_, err = toml.DecodeFile(currPath, &config)
	if err != nil {
		err = errors.New("配置文件读取失败 - " + err.Error())
		return
	}
	return
}

func PrintAndLog(msg string) {
	fmt.Println(msg)
	err := go_tool.Log(msg)
	if err != nil {
		fmt.Println(err)
	}
}
