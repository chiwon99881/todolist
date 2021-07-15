package main

import (
	"github.com/chiwon99881/todolist/db"
	"github.com/chiwon99881/todolist/env"
	"github.com/chiwon99881/todolist/rest"
	_ "github.com/lib/pq"
)

func main() {
	defer db.Close()
	env.Start()
	rest.Start()
}
