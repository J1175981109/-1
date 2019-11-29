package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Println("链接数据库失败")
	}else{
		fmt.Println("连接数据库成功")
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO test VALUES ( 3, 'JJP',42)")

	if err != nil {

		panic(err.Error())

	}
	defer insert.Close()

	}
