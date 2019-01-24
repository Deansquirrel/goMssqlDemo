package dbOpr2

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDbConn(server string, port int, dbName string, user string, pwd string) (*sql.DB, error) {
	dbConnStrFormat := "server=%s;port=%d;user id=%s;password=%s;database=%s"
	connStr := fmt.Sprintf(dbConnStrFormat,server,port,user,pwd,dbName)

	db,err := sql.Open("mssql",connStr)
	if err != nil {
		return nil,err
	}
	return db,nil
}


//
//func GetDbConn(server string, port int, dbName string, user string, pwd string) (*sql.DB, error) {
////	query := url.Values{}
////	//query.Add("app name","MyAppName")
////	query.Add("database",dbName)
////
////	var userInfo *url.Userinfo
////	if pwd == "" {
////		userInfo = url.User(user)
////	} else {
////		userInfo = url.UserPassword(user,pwd)
////	}
////
////	u := &url.URL{
////		Scheme:"sqlserver",
////		User:userInfo,
////		Host:fmt.Sprintf("%s:%d",server,port),
////		RawQuery:query.Encode(),
////}
////
////	fmt.Println(u.String())
////
////	db,err := sql.Open("sqlserver",u.String())
////	if err != nil {
////		return nil,err
////	}
//
//	dbConnStrFormat := "server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable"
//	connStr := fmt.Sprintf(dbConnStrFormat,server,dbName,user,pwd,port)
//
//	db,err := sql.Open("mssql",connStr)
//
//	err = db.Ping()
//	if err != nil {
//		return nil,err
//	}
//
//	db.SetMaxIdleConns(30)
//	db.SetMaxOpenConns(30)
//	db.SetConnMaxLifetime(time.Second * 60 * 10)
//	return db, nil
//}