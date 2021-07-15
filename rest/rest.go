package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port string = ":3000"
)

type errResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type url string

func (u url) MarshalText() ([]byte, error) {
	urlString := fmt.Sprintf("http://localhost%s%s", port, u)
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

// Start is trigger for rest api gateway
func Start() {
	fmt.Printf("Server Listening on http://localhost%s\n", port)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err.Error())
	}
}
