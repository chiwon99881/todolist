package main

import (
	"github.com/chiwon99881/todolist/db"
	"github.com/chiwon99881/todolist/env"
	"github.com/chiwon99881/todolist/web"
	_ "github.com/lib/pq"
)

func main() {
	defer db.Close()
	env.Start()

	web.Start()
}
