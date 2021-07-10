package todolist

import (
	"github.com/chiwon99881/todolist/db"
	"github.com/chiwon99881/todolist/types"
)

// LoadAllToDo is get all todo
func LoadAllToDo() []*types.ToDo {
	toDos := db.SelectAllToDo()
	return toDos
}

// AddToDo is create new to do
func AddToDo(caption string) {
	db.InsertToDo(caption, false)
}

// DoneToDo is my a to do is done
func DoneToDo(ID int) {
	toDo := db.SelectToDo(ID)
	db.UpdateToDo(toDo.ID, toDo.Excute)
}

// RemoveToDo is delete to do in todolist
func RemoveToDo(ID int) {
	db.DeleteToDo(ID)
}
