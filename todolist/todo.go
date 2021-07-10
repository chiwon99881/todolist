package todolist

import (
	"fmt"

	"github.com/chiwon99881/todolist/db"
)

// ToDo Type
type ToDo struct {
	ID      int
	Caption string
	Excute  bool
}

// LoadAllToDo is get all todo
func LoadAllToDo() {
	toDos := db.SelectAllToDo()
	for _, toDo := range toDos {
		fmt.Println(toDo)
	}
}
