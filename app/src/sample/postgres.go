package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	// "time"
)

const (
	DB_HOST     = "192.168.20.100"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_SCHEMA   = "go_sample"
)

func main() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	fmt.Println("# Querying")

	query := fmt.Sprintf("SELECT * FROM %s.category", DB_SCHEMA)
	rows, err := db.Query(query)
	checkErr(err)

	for rows.Next() {
		var uid int
		var name string
		err = rows.Scan(&uid, &name)
		checkErr(err)
		fmt.Println("uid | name ")
		fmt.Printf("%3v | %8v \n", uid, name)
	}

	// fmt.Println("# Deleting")
	// stmt, err = db.Prepare("delete from userinfo where uid=$1")
	// checkErr(err)

	// res, err = stmt.Exec(lastInsertId)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
