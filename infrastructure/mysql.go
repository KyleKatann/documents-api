package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var RDB *sql.DB

func init() {
	dbHost := "localhost"
	dbUser := "user"
	dbPassword := "password"
	dbName := "documents-api"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	var err error
	RDB, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Println("error")
	}
}
