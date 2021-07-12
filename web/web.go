package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/chiwon99881/todolist/todolist"
	"github.com/chiwon99881/todolist/types"
	"github.com/gorilla/mux"
)

var templates *template.Template
var templateDir = "web/templates/"

func home(rw http.ResponseWriter, r *http.Request) {
	toDos := todolist.LoadAllToDo()

	err := templates.Execute(rw, types.LoadAllToDoData{ToDos: toDos})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
}

// Start is function of web server
func Start() {
	fmt.Printf("Web server running on http://localhost:%s", os.Getenv("WEBPORT"))
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")

	err := http.ListenAndServe(os.Getenv("WEBPORT"), router)
	if err != nil {
		panic(err.Error())
	}
}
