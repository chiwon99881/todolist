package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chiwon99881/todolist/env"
	_ "github.com/lib/pq"
)

func main() {
	env.Start()
	dbSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", dbSource)
	defer db.Close()

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}
}
