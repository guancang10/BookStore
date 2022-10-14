package main

import (
	"github.com/guancang10/BookStore/API/appdb"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/joho/godotenv"
)

func main() {
	//Load env from .env
	err := godotenv.Load(".env")
	helper.CheckError(err)
	appdb.GetConnection()
}
