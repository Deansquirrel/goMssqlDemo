package dbOpr3

import (
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goMssqlDemo/common"
	"github.com/Deansquirrel/goMssqlDemo/global"
	"time"
)

func Select() error {
	common.PrintAndLog("dbOpr3 Select")
	defer common.PrintAndLog("dbOpr3 Select Done")

	sqlSelectStr := "" +
		"Select top 10 * from IISlog"

	db, err := GetDbConn(global.SysConfig.MsSqlConfig.Server, global.SysConfig.MsSqlConfig.Port, global.SysConfig.MsSqlConfig.Database, global.SysConfig.MsSqlConfig.User, global.SysConfig.MsSqlConfig.Password)
	if err != nil {
		return err
	}
	defer func() {
		errLs := db.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	rows,err := db.Query(sqlSelectStr)
	if err != nil {
		return err
	}
	defer func(){
		errLs := rows.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	var t time.Time
	var c int
	for rows.Next(){
		err = rows.Scan(&t,&c)
		if err != nil {
			return err
		} else {
			fmt.Println(go_tool.GetDateTimeStr(t),c)
		}
	}

	return nil
}

func MultipleCommand() error {
	common.PrintAndLog("dbOpr3 MultipleCommand")
	defer common.PrintAndLog("dbOpr3 MultipleCommand Done")

	sqlCreateStr := "" +
		"CREATE TABLE #IISlog" +
		"(" +
		"	[sj] [datetime] not NULL," +
		"	[ms] [int] not NULL" +
		") ON [PRIMARY]"

	//sqlInsertStr := "" +
	//	"INSERT INTO #IISlog([sj],[ms]) select ?,?"

	sqlDropStr := "" +
		"Drop Table #IISlog"

	db, err := GetDbConn(global.SysConfig.MsSqlConfig.Server, global.SysConfig.MsSqlConfig.Port, global.SysConfig.MsSqlConfig.Database, global.SysConfig.MsSqlConfig.User, global.SysConfig.MsSqlConfig.Password)
	if err != nil {
		return err
	}
	defer func() {
		errLs := db.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	tx,err := db.Begin()
	if err != nil {
		return err
	}
	defer func(){
		errLs := tx.Rollback()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	_,err = tx.Exec(sqlCreateStr)
	if err != nil {
		return err
	}

	//_,err = tx.Exec(sqlInsertStr)
	//if err != nil {
	//	return err
	//}

	_,err = tx.Exec(sqlDropStr)
	if err != nil {
		return err
	}

	return nil
}