package main

import (
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/injector"
	"github.com/joho/godotenv"
)

func main() {
	//Load env from .env
	err := godotenv.Load(".env")
	helper.CheckError(err)

	server := injector.InitServer()
	err = server.ListenAndServe()
	helper.CheckError(err)
}
