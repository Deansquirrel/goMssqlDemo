package dbOpr

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goMssqlDemo/common"
	"github.com/Deansquirrel/goMssqlDemo/global"
	"github.com/Deansquirrel/goZlDianzqOfferTicket/repository"
	"time"
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

func Select() (err error) {
	if !repository.CheckV(conn) {
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

	common.PrintAndLog("signal Test")

	common.PrintAndLog("Done")
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
	if !repository.CheckV(conn) {
		err := refreshConn()
		if err != nil {
			common.PrintAndLog("RefreshCon error:" + err.Error())
			return err
		}
	}

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	tx,err := conn.Begin()
	if err != nil {
		return err
	}
	defer func(){
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	for i:=0;i<100;i++{
		_,err = tx.Exec(getInsertStrSql(),go_tool.GetDateTimeStr(time.Now()),1)
		if err != nil {
			return err
		}
	}

	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())
	return nil
}

func getInsertStrSql() string{
	sqlStr := "" +
		"INSERT INTO IISlog([sj],[ms]) select ?,?"
	return sqlStr
}