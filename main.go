package main

import (
	"github.com/chiwon99881/todolist/db"
	_ "github.com/lib/pq"
)

// ToDo Type
type ToDo struct {
	ID      int
	Caption string
	Excute  bool
}

func main() {
	db.InsertToDoDB("Make To Do List Project", false)
	defer db.Close()
}
