package main

import (
	"database/sql"
	"log"

	"github.com/bibarsss/simple-rest-api-go/cmd/api"
	"github.com/bibarsss/simple-rest-api-go/config"
	"github.com/bibarsss/simple-rest-api-go/db"
	"github.com/go-sql-driver/mysql"
)

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
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":25555", db)
	if err := server.Run(); err != nil {
		log.Fatal(err, err.Error())
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
