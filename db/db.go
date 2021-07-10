package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chiwon99881/todolist/env"
)

// DB is connection func for database
func DB() *sql.DB {
	env.Start()
	pqConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", pqConn)

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}

	return db
}

// Close is func that terminate database
func Close() {
	DB().Close()
}

// InsertToDoDB is excute insert SQL
func InsertToDoDB(caption string, check bool) {
	stmt := `insert into "todo"("caption", "excute") values($1, $2)`
	_, err := DB().Exec(stmt, caption, check)
	if err != nil {
		panic(err.Error())
	}
}
