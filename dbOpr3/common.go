package dbOpr3

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-adodb"
	"strconv"
	"strings"
)

func GetDbConn(server string, port int, dbName string, user string, pwd string) (*sql.DB, error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+server + "," + strconv.Itoa(port))
	//if m.windows {
	//	// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
	//	// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
	//	conf = append(conf, "integrated security=SSPI")
	//}
	conf = append(conf, "Initial Catalog="+dbName)
	conf = append(conf, "user id="+user)
	conf = append(conf, "password="+pwd)

	fmt.Println(strings.Join(conf, ";"))

	db,err := sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return nil,err
	}

	err = db.Ping()
	if err != nil {
		return nil,err
	} else {
		return db,nil
	}
}