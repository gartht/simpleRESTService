package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Welcome")
}

func TodoIndex(response http.ResponseWriter, request *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(response).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(response, "Todos show:", todoId)
}
