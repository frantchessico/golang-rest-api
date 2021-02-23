package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	
	
)

type tasks struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []tasks

var task = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3200", router))
}

