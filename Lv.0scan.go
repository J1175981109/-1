package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/test")
	defer db.Close()

	rowsQuery, err := db.Query("SELECT * FROM test WHERE score=7")
	if err != nil {
		fmt.Println("select db failed,err:", err)
		return
	} else {
		for rowsQuery.Next() {
			var uid int
			var name string
			var score string
			err = rowsQuery.Scan(&uid, &name, &score)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(uid,name,score)
		}
		defer rowsQuery.Close()
	}
}