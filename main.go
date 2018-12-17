package main

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goMssqlDemo/global"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

func main() {
	fmt.Println("程序启动")
	defer fmt.Println("程序退出")

	//读取配置文件
	var config global.SysConfig
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		msg := "配置文件读取失败\n" + err.Error()
		fmt.Println(msg)
		go_tool.Log(msg)
	}

	var isDebug = true
	var server = config.MsSqlConfig.Server
	var port = config.MsSqlConfig.Port
	var user = config.MsSqlConfig.User
	var password = config.MsSqlConfig.Password
	var database = config.MsSqlConfig.Database

	//连接字符串
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", server, database, user, password, port)
	if isDebug {
		fmt.Println(connString)
	}

	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()

	//产生查询语句的Statement
	stmt, err := conn.Prepare(`SELECT * FROM Authorization_info`)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	//通过Statement执行查询
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}

	//建立一个列数组
	cols, err := rows.Columns()
	var colsdata = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()

	//遍历每一行
	for rows.Next() {
		rows.Scan(colsdata...) //将查到的数据写入到这行中
		PrintRow(colsdata)     //打印此行
	}
	defer rows.Close()
}

//打印一行记录，传入一个行的所有列信息
func PrintRow(colsdata []interface{}) {
	for _, val := range colsdata {
		switch v := (*(val.(*interface{}))).(type) {
		case nil:
			fmt.Print("NULL")
		case bool:
			if v {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}
		case []byte:
			fmt.Print(string(v))
		case time.Time:
			//fmt.Print(v.Format("2016-01-02 15:05:05.999"))
			fmt.Print(go_tool.GetDateTimeStr(v))
		default:
			fmt.Print(v)
		}
		fmt.Print("\t")
	}
	fmt.Println()
}
