package main

import (
	"fmt"
	"resnet/database"
)

func main() {	
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer database.Disconnect(db)

	_, e := db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INT,
		email VARCHAR(255)
	)`)
	if err != nil {
		panic(e)
	}

	rows, e := db.Query("SELECT * FROM users")
	if e != nil {
		panic(e)
	}

	for rows.Next() {
		var id int
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d\nEmail: %s\n", id, email)
	}

	fmt.Print("Hello World!\n")
}