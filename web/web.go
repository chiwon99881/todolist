package web

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

		err := templates.ExecuteTemplate(rw, "home", types.LoadAllToDoData{ToDos: toDos})
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
		toDoID, err := strconv.Atoi(getID[1])
		if err != nil {
			panic(err.Error())
		}
		todolist.DoneToDo(toDoID)
		http.Redirect(rw, r, "/", http.StatusFound)
		break
	default:
		break
	}
}

func addToDo(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := templates.ExecuteTemplate(rw, "add-todo", "")
		if err != nil {
			panic(err.Error())
		}
		break
	case "POST":
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		convValue := bytes.NewBuffer(val).String()
		splitValue := strings.Split(convValue, "=")
		replaceValue := strings.ReplaceAll(splitValue[1], "+", " ")
		todolist.AddToDo(replaceValue)
		http.Redirect(rw, r, "/", http.StatusFound)
		break
	}
}

func deleteToDo(rw http.ResponseWriter, r *http.Request) {
	val, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	convValue := bytes.NewBuffer(val).String()
	splitValue := strings.Split(convValue, "=")
	aToIntValue, err := strconv.Atoi(splitValue[1])
	if err != nil {
		panic(err.Error())
	}
	todolist.RemoveToDo(aToIntValue)
	http.Redirect(rw, r, "/", http.StatusFound)
}

// Start is function of web server
func Start() {
	fmt.Printf("Web server running on http://localhost:%s", os.Getenv("WEBPORT"))
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET", "POST")
	router.HandleFunc("/todo/add", addToDo).Methods("GET", "POST")
	router.HandleFunc("/todo/delete", deleteToDo).Methods("POST")
	err := http.ListenAndServe(os.Getenv("WEBPORT"), router)
	if err != nil {
		panic(err.Error())
	}
}
