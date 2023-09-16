package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	dataSource := "user:password@tcp(localhost:3309)/finalGogogo"
	// Open inicia un pool de conexiones. Solo abrir una vez
	storageDB, err := sql.Open("mysql", dataSource)
	fmt.Println("Dentro del init()")
	if err != nil {
		panic(err)
	}
	return storageDB, nil
}

//type SqlStore struct {
//	*sql.DB
//}

/* func NewPostgresSQLDatabase(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection to ensure it's valid
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
} */
