package main

import (
	"github.com/chiwon99881/todolist/db"
	_ "github.com/lib/pq"
)

func main() {

	defer db.Close()
}
