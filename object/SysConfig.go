package object

import "fmt"

type SysConfig struct {
	MsSqlConfig mssqlConfig `toml:"mssql"`
}

type mssqlConfig struct {
	Server   string `toml:"server"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

func (config SysConfig) PrintFormat() {
	fmt.Println("============================================")
	fmt.Println("Config Format")
	fmt.Println(`[mssql]
server=""
port=
user=""
password=""
database=""`)
	fmt.Println("============================================")
}
