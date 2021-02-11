package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



type Book struct {
	Id        int    `json:"id,omitempty"`
	Author    string `json:"author"`
	Publisher  string `json:"publisher"`

}


func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/health", health)
	myRouter.HandleFunc("/detail", detail)

	log.Fatal(http.ListenAndServe(":9080", myRouter))
}

func main() {
	fmt.Println(" Rest api start")

	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func health(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "health")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Details is healthy")
}

func detail(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "health")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)




	book := Book{
		Author: "William S",
		Publisher:  "Publisher A",

	}

	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr1 := string(data)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr1)
	fmt.Println(data)

	json.NewEncoder(w).Encode(book)
}

