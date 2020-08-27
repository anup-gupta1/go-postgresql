package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "@Anup027gupta"
	DB_NAME     = "bookstore"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()
	checkError(err)
	var lastInsertedID int
	err = db.QueryRow("INSERT INTO EMPLOYEES(ID, NAME) VALUES($1,$2) returning ID", 4, "Ankit").Scan(&lastInsertedID)
	checkError(err)
	fmt.Print("last inserted id -------------------")
	fmt.Println(lastInsertedID)
	// stmt, err := db.Prepare("update employees set name=Ashish Aggaarwal where id=3")
	// checkError(err)
	// fmt.Println(stmt)
	rows, err := db.Query("SELECT * FROM EMPLOYEES")
	checkError(err)
	fmt.Println(rows)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		checkError(err)
		fmt.Println(id, name)
	}
	fmt.Println("hello server")
	fmt.Println(http.ListenAndServe(":5000", nil))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
