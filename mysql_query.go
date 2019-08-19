package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db,dbErr:=sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/laravel?charset=utf8")
	if dbErr!=nil{
		panic(db)
	}
	dbQuery,dbpErr:=db.Query("SELECT * FROM users")
	if dbpErr!=nil{
		panic(dbErr)
	}
	for dbQuery.Next(){
		var id int
		var no int
		var googleSecret string
		var createdAt int
		var updatedAt int
		dbqErr:=dbQuery.Scan(&id,&no,&googleSecret,&createdAt,&updatedAt)
		if dbqErr!=nil{
			panic(dbqErr)
		}
		fmt.Println(id,no,googleSecret,createdAt,updatedAt)
	}
}