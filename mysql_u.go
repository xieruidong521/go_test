package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db,dbErr:=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/laravel?charset=utf8")
	if dbErr!=nil {
		panic(dbErr)
	}
	dbPrepare,dbpError:=db.Prepare("UPDATE users SET no=?,googlesecret=?,created_at=?,updated_at=? WHERE id=?")
	if dbpError!=nil {
		panic(dbpError)
	}
	dbExec,dbeError:=dbPrepare.Exec("666666","asdKKad","1666666666","1666666666",11)
	if dbeError!=nil{
		panic(dbeError)
	}
	dbId,_:=dbExec.RowsAffected()
	fmt.Println(dbId)
}
