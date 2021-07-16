package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/chiwon99881/todolist/db"
	"github.com/chiwon99881/todolist/todolist"
	"github.com/chiwon99881/todolist/utils"
	"github.com/gorilla/mux"
)

type errResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type responseData struct {
	Caption string `json:"caption"`
}

type url string

func (u url) MarshalText() ([]byte, error) {
	urlString := fmt.Sprintf("http://localhost%s%s", os.Getenv("RESTPORT"), u)
	return []byte(urlString), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

func home(rw http.ResponseWriter, r *http.Request) {
	docs := []urlDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See documentation",
		},
		{
			URL:         "/todo/add",
			Method:      "POST",
			Description: "Add to do in todolist",
		},
	}
	marshal, err := json.Marshal(docs)
	if err != nil {
		panic(err.Error())
	}
	rw.Header().Add("Content-Type", "application/json")
	_, err = fmt.Fprintf(rw, "%s", marshal)
	if err != nil {
		fmt.Fprintf(rw, "%s", errResponse{err.Error()})
	}
}

func addToDo(rw http.ResponseWriter, r *http.Request) {
	responseData := &responseData{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(rw, "%s", errResponse{ErrorMessage: "error..."})
		panic(err.Error())
	}
	err = json.Unmarshal(body, responseData)
	if err != nil {
		fmt.Fprintf(rw, "%s", errResponse{ErrorMessage: "error..."})
		panic(err.Error())
	}
	todolist.AddToDo(responseData.Caption)
	rw.WriteHeader(http.StatusCreated)
}

func updateToDo(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convInt, err := strconv.Atoi(id)
	utils.HandleError(err)
	todolist.DoneToDo(convInt)
	rw.WriteHeader(http.StatusOK)
}

func deleteToDo(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convInt, err := strconv.Atoi(id)
	utils.HandleError(err)
	todolist.RemoveToDo(convInt)
	rw.WriteHeader(http.StatusOK)
}

func getToDo(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convInt, err := strconv.Atoi(id)
	utils.HandleError(err)

	toDo := db.SelectToDo(convInt)
	marshal, err := json.Marshal(toDo)
	utils.HandleError(err)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s", marshal)
}

func allToDo(rw http.ResponseWriter, r *http.Request) {
	toDos := todolist.LoadAllToDo()
	marshal, err := json.Marshal(toDos)
	utils.HandleError(err)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "%s", marshal)
}

// Start is trigger for rest api gateway
func Start() {
	fmt.Printf("Server Listening on http://localhost%s\n", os.Getenv("RESTPORT"))
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/todo/add", addToDo).Methods("GET", "POST")
	router.HandleFunc("/todo/{id:[0-9]+}", getToDo).Methods("GET")
	router.HandleFunc("/todos", allToDo).Methods("GET")
	router.HandleFunc("/update/todo/{id:[0-9]+}", updateToDo).Methods("GET")
	router.HandleFunc("/delete/todo/{id:[0-9]+}", deleteToDo).Methods("GET")
	err := http.ListenAndServe(os.Getenv("RESTPORT"), router)
	if err != nil {
		panic(err.Error())
	}
}
