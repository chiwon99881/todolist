package main

import (
	"github.com/chiwon99881/todolist/db"
	"github.com/chiwon99881/todolist/todolist"
	_ "github.com/lib/pq"
)

func main() {
	todolist.LoadAllToDo()
	defer db.Close()
}
