package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PieterJDSw/GoRestApi/cmd/api"
	"github.com/PieterJDSw/GoRestApi/config"
	"github.com/PieterJDSw/GoRestApi/db"
	"github.com/go-sql-driver/mysql"
)

const port = ":8080"

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {

		log.Fatal("Error With Db config")
	}
	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	errorRun := server.Run()

	if errorRun != nil {

		log.Fatal("Error")
	} else {

		log.Println("Running on port:  " + port)
	}
	// 31:25

}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {

		log.Fatal(err)

	}
	log.Println("DB connected")

}
