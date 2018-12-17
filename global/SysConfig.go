package global

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
