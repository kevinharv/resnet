package database

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "resnet"
)

func Connect() (dbptr *sql.DB, err error) {
	conn_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn_string)
    if err != nil {
		panic(err)
	}
 
    err = db.Ping()
    if err != nil {
		return nil, fmt.Errorf("failed to PING DB")
	}
 
    fmt.Println("Connected!")
	return db, nil
}

func Disconnect(db *sql.DB) {
	fmt.Println("Disconnected!")
	db.Close()
}