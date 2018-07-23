package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"os"
)

func main(){
	db, err := xorm.NewEngine("mysql", "root@tcp(127.0.0.1:3306)/app_test")
	if err != nil {
		fmt.Printf("init database error: %v", err)
		os.Exit(1)
	}
	db.ShowSQL(true)

	session := db.NewSession()
	var results []interface{}
	err = session.Table("interva").Where("id = ?; drop table interva;", 1).Find(&results)
	if err != nil {
		fmt.Printf("sql exec fail: %v", err)
		os.Exit(1)
	}

	fmt.Printf("results: %v", results)
}

