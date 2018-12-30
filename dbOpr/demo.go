package dbOpr

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Deansquirrel/goMssqlDemo/common"
	"github.com/Deansquirrel/goMssqlDemo/global"
	"log"
)

var conn *sql.DB

func refreshConn() (err error) {
	common.PrintAndLog("Refresh Conn")
	c, err := GetDbConn(global.SysConfig.MsSqlConfig.Server, global.SysConfig.MsSqlConfig.Port,
		global.SysConfig.MsSqlConfig.Database, global.SysConfig.MsSqlConfig.User, global.SysConfig.MsSqlConfig.Password)
	if err != nil {
		return err
	}
	conn = c
	return
}

func Select() (err error) {
	if conn == nil {
		err := refreshConn()
		if err != nil {
			return err
		}
	}
	defer func() {
		errLs := conn.Close()
		if errLs != nil {
			log.Fatal("Conn close error:", errLs.Error())
		}
	}()

	//产生查询语句的Statement
	stmt, err := conn.Prepare("" +
		"SELECT * FROM master..SysDatabases")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
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
		common.PrintAndLog(err.Error())
		log.Fatal("Query failed:", err.Error())
	}
	defer func() {
		errLs := rows.Close()
		if errLs != nil {
			common.PrintAndLog(errLs.Error())
		}
	}()

	cls, err := rows.Columns()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(cls)

	return
}

func MultipleCommand() error {
	if conn == nil {
		err := refreshConn()
		if err != nil {
			return err
		}
	}
	defer func() {
		errLs := conn.Close()
		if errLs != nil {
			log.Fatal("Conn close error:", errLs.Error())
		}
	}()

	ctx := context.TODO()
	defer ctx.Done()
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			errLs := tx.Rollback()
			if errLs != nil {
				common.PrintAndLog(errLs.Error())
			}
		} else {
			errLs := tx.Commit()
			if errLs != nil {
				common.PrintAndLog(errLs.Error())
			}
		}
	}()

	_, err = tx.Exec("CREATE TABLE #TktInfo" +
		"(" +
		"    Appid	varchar(30)," +
		"    Accid	bigint," +
		"    Tktno	varchar(30)," +
		"    Cashmy	decimal(18,2)," +
		"    Addmy	decimal(18,2)," +
		"    Tktname	nvarchar(30)," +
		"    TktKind	varchar(30)," +
		"    Pcno	varchar(30)," +
		"    EffDate	smalldatetime," +
		"    Deadline	smalldatetime," +
		"    CrYwlsh	varchar(12)," +
		"    CrBr	varchar(30)" +
		")")
	if err != nil {
		return err
	}

	_, err = tx.Exec("" +
		"insert into #TktInfo(Appid,Accid,Tktno,Cashmy,Addmy,Tktname,TktKind,Pcno,EffDate,Deadline,CrYwlsh,CrBr)" +
		"select ?,?,?,?,?,?,?,?,?,?,?,?")
	if err != nil {
		return err
	}

	_, err = tx.Exec("" +
		"exec pr_CreateLittleTkt_Create")
	if err != nil {
		return err
	}

	_, err = tx.Exec("" +
		"drop table #TktInfo")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}
