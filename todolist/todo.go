package todolist

import (
	"fmt"

	"github.com/chiwon99881/todolist/db"
)

// LoadAllToDo is get all todo
func LoadAllToDo() {
	toDos := db.SelectAllToDo()
	for _, toDo := range toDos {
		fmt.Println(*toDo)
	}
}

// DoneToDo is my a to do is done
func DoneToDo(ID int) {
	toDo := db.SelectToDo(ID)
	db.UpdateToDo(toDo.ID, toDo.Excute)
}
