package dbOpr2

import (
	"github.com/Deansquirrel/goMssqlDemo/common"
	"github.com/Deansquirrel/goMssqlDemo/global"
)

func MultipleCommand() error {
	common.PrintAndLog("dbOpr2 MultipleCommand")
	defer common.PrintAndLog("dbOpr2 MultipleCommand Done")

	sqlSelectStr := "" +
		"Select * from IISlog"

	//sqlCreateStr := "" +
	//	"CREATE TABLE #IISlog" +
	//	"(" +
	//	"	[sj] [datetime] not NULL," +
	//	"	[ms] [int] not NULL" +
	//	") ON [PRIMARY]"
	//
	//sqlInsertStr := "" +
	//	"INSERT INTO #IISlog([sj],[ms]) select ?,?"
	//
	//sqlDropStr := "" +
	//	"Drop Table #IISlog"

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

	for rows.Next(){

	}


	//_, err = db.Exec(sqlCreateStr)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = db.Exec(sqlInsertStr)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = db.Exec(sqlDropStr)
	//if err != nil {
	//	return err
	//}

	return nil
}