package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/chiwon99881/todolist/todolist"
	"github.com/chiwon99881/todolist/types"
	"github.com/gorilla/mux"
)

var templates *template.Template
var templateDir = "web/templates/"

func home(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		toDos := todolist.LoadAllToDo()

		err := templates.Execute(rw, types.LoadAllToDoData{ToDos: toDos})
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
		break
	case "POST":
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		id := fmt.Sprintf("%s", res)
		getID := strings.Split(id, "=")
		fmt.Printf("\nID:%s\n", getID[1])
	default:
		break
	}
}

// Start is function of web server
func Start() {
	fmt.Printf("Web server running on http://localhost:%s", os.Getenv("WEBPORT"))
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET", "POST")

	err := http.ListenAndServe(os.Getenv("WEBPORT"), router)
	if err != nil {
		panic(err.Error())
	}
}
