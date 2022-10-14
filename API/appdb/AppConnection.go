package appdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guancang10/BookStore/API/helper"
	"os"
	"time"
)

func GetConnection() *sql.DB {
	//Get username and pass from env
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dataSource := username + ":" + password + "@tcp(localhost:3306)/BookStore?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	helper.CheckError(err)

	//Setting max idle connection
	db.SetMaxIdleConns(10)
	//setting max con max (so if the request too many, system will open new connection)
	db.SetMaxOpenConns(20)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	fmt.Println(dataSource)
	return db
}
