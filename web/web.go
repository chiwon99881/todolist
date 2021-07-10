package web

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var templates *template.Template
var templateDir = "templates/"

func home(rw http.ResponseWriter, r *http.Request) {
	err := templates.Execute(rw, "HI! MY HOME")
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
