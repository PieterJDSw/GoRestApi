package main

import (
	"log"

	"github.com/PieterJDSw/GoRestApi/cmd/api"
	"github.com/PieterJDSw/GoRestApi/db"
	"github.com/go-sql-driver/mysql"
)

const port = ":8080"

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "asd",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(port, nil)
	err := server.Run()
	if err != nil {

		log.Fatal("Error")
	} else {

		log.Println("Running on port:  " + port)
	}
	// 14:47
}
