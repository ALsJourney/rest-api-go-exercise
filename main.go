package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "db:3306",
		DBName:               "projectmanager",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	api := NewApiServer(":3000", store)

	api.Serve()

}
