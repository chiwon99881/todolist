package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chiwon99881/todolist/types"
)

// DB is connection func for database
func DB() *sql.DB {
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

// SelectAllToDo is excute select SQL
func SelectAllToDo() []*types.ToDo {
	var toDos []*types.ToDo
	stmt := `select * from "todo" order by id asc`
	rows, err := DB().Query(stmt)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		toDo := &types.ToDo{}
		rows.Scan(&toDo.ID, &toDo.Caption, &toDo.Excute, &toDo.Index)
		toDos = append(toDos, toDo)
	}
	return toDos
}

// SelectToDo is find To Do
func SelectToDo(ID int) *types.ToDo {
	toDo := &types.ToDo{}
	stmt := `select * from "todo" where "id"=$1`
	row := DB().QueryRow(stmt, ID)

	row.Scan(&toDo.ID, &toDo.Caption, &toDo.Excute, &toDo.Index)
	return toDo
}

// InsertToDo is excute insert SQL
func InsertToDo(caption string, check bool) {
	toDos := SelectAllToDo()
	index := len(toDos)
	stmt := `insert into "todo"("caption", "excute", "index") values($1, $2, $3)`
	_, err := DB().Exec(stmt, caption, check, index+1)
	if err != nil {
		panic(err.Error())
	}
}

// UpdateToDo is true -> false or false -> true your ToDo
func UpdateToDo(ID int, check bool) {
	toggleExcute := !check
	stmt := `update "todo" set "excute"=$2 where "id"=$1`
	_, err := DB().Exec(stmt, ID, toggleExcute)
	if err != nil {
		panic(err.Error())
	}
}

// DeleteToDo is delete to do in database
func DeleteToDo(ID int) {
	stmt := `delete from "todo" where id=$1`
	_, err := DB().Exec(stmt, ID)
	if err != nil {
		panic(err.Error())
	}
	toDos := SelectAllToDo()
	for i := 0; i < len(toDos); i++ {
		stmt := `update "todo" set "index"=$1 where "id"=$2`
		_, err := DB().Exec(stmt, i+1, toDos[i].ID)
		if err != nil {
			panic(err.Error())
		}
	}
}
