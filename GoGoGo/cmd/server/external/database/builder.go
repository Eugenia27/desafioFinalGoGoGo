package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	dataSource := "user:password@tcp(localhost:3309)/finalGogogo"
	storageDB, err := sql.Open("mysql", dataSource)
	fmt.Println("Dentro del init()")
	if err != nil {
		panic(err)
	}
	return storageDB, nil
}
