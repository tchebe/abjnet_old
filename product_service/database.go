package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func createSqlServerDBConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return gorm.Open(
		"mssql",
		fmt.Sprintf("sqlserver://%s:%s@%s?database=%s?", user, password, host, dbname),
	)

}
