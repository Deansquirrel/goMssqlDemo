package dbOpr

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Deansquirrel/goMssqlDemo/common"
	"github.com/Deansquirrel/goMssqlDemo/global"
	"github.com/Deansquirrel/goZlDianzqOfferTicketV3/repository"
)

var conn *sql.DB

var ctx = context.Background()

func refreshConn() (err error) {
	common.PrintAndLog("Refresh Conn")
	c, err := GetDbConn(global.SysConfig.MsSqlConfig.Server, global.SysConfig.MsSqlConfig.Port,
		global.SysConfig.MsSqlConfig.Database, global.SysConfig.MsSqlConfig.User, global.SysConfig.MsSqlConfig.Password)
	if err != nil {
		common.PrintAndLog("refreshConn error" + err.Error())
		return err
	}
	conn = c
	return
}

func GetConn()error{
	return refreshConn()
}

func getTestSql()string{
	sSql := "SELECT * FROM master..SysDatabases"
	return sSql
}

func Test()error{

	rows,err := conn.Query(getTestSql())
	if err != nil {
		common.PrintAndLog("Query failed:" + err.Error())
		return err
	}
	defer func() {
		//_ = rows.Close()
		errLs := rows.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	cls, err := rows.Columns()
	if err != nil {
		common.PrintAndLog("rows close error:" + err.Error())
		return err
	}
	fmt.Println(cls)

	return nil
}

func Select() (err error) {
	common.PrintAndLog("begin Select")
	defer common.PrintAndLog("Select Done")
	if !IsConnValid(conn){
		err := refreshConn()
		if err != nil {
			common.PrintAndLog("RefreshCon error:" + err.Error())
			return err
		}
	}

	//产生查询语句的Statement
	stmt, err := conn.Prepare("" +
		"SELECT * FROM master..SysDatabases")
	if err != nil {
		common.PrintAndLog("Prepare failed:" + err.Error())
	}
	defer func() {
		errLs := stmt.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	//通过Statement执行查询
	rows, err := stmt.Query()
	if err != nil {
		common.PrintAndLog("Query failed:" + err.Error())
	}
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	cls, err := rows.Columns()
	if err != nil {
		common.PrintAndLog("rows close error:" + err.Error())
	}
	fmt.Println(cls)
	return
}

func MultipleCommand() error {
	common.PrintAndLog("begin MultipleCommand")
	defer common.PrintAndLog("MultipleCommand Done")
	if !repository.CheckV(conn) {
		err := refreshConn()
		if err != nil {
			common.PrintAndLog("RefreshCon error:" + err.Error())
			return err
		}
	}

	_,err := conn.Exec(getCreateTempTableSqlStr() + " " + getDropTempTableSqlStr())
	if err != nil {
		return err
	}

	return nil
}


func getCreateTempTableSqlStr() string {
	sqlStr := "" +
		"CREATE TABLE #IISlog" +
		"(" +
		"	[sj] [datetime] not NULL," +
		"	[ms] [int] not NULL" +
		") ON [PRIMARY]"
	return sqlStr
}

func getInsertSqlStrModel() string {
	sqlStr := "" +
		"INSERT INTO #IISlog([sj],[ms]) select ?,?"
	return sqlStr
}

func getDropTempTableSqlStr() string {
	sqlStr := "" +
		"Drop Table #IISlog"
	return sqlStr
}

func TxCommand() error{
	common.PrintAndLog("begin TxCommand")
	defer common.PrintAndLog("TxCommand Done")

	tabName := "#IISlog20190107"

	db,err := GetDbConn(global.SysConfig.MsSqlConfig.Server,global.SysConfig.MsSqlConfig.Port,global.SysConfig.MsSqlConfig.Database,global.SysConfig.MsSqlConfig.User,global.SysConfig.MsSqlConfig.Password)
	if err != nil {
		return err
	}
	defer func(){
		errLs := db.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	tx,err := db.Begin()
	defer func(){
		_ = tx.Rollback()
	}()

	sqlStr := "" +
		"CREATE TABLE " + tabName +
		"(" +
		"	[sj] [datetime] not NULL," +
		"	[ms] [int] not NULL" +
		") ON [PRIMARY]"

	common.PrintAndLog("Create Table")

	_,err = tx.Exec(sqlStr)
	if err != nil {
		return err
	}
	common.PrintAndLog("Create Table Done")

	common.PrintAndLog("Search Table")
	rows,err := tx.Query("" +
		"SELECT * FROM " + tabName)
	if err != nil {
		return err
	}

	for rows.Next() {

	}
	common.PrintAndLog("Search Table Done")


	common.PrintAndLog("Drop Table")
	_,errLs := tx.Exec("" +
		"Drop table " + tabName)
	if errLs != nil {
		common.PrintAndLog(errLs.Error())
	}
	common.PrintAndLog("Drop Table Done")

	//
	//stmt,err := tx.PrepareContext(ctx,getInsertSqlStrModel())
	//defer func(){
	//	_ = stmt.Close()
	//}()
	//
	//_,err = stmt.ExecContext(ctx,go_tool.GetDateTimeStr(time.Now()),10)
	//if err != nil {
	//	return err
	//}

	return nil
}

func getInsertStrSql() string{
	sqlStr := "" +
		"INSERT INTO IISlog([sj],[ms]) select ?,?"
	return sqlStr
}