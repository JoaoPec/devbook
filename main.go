package main

import (
	"devbook/database"
	"devbook/server"
	"fmt"
	mux "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	database.StartDB()

	router.HandleFunc("/users", server.CreateUser).Methods("POST")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to DevBook")
	})

	router.HandleFunc("/users/{id}", server.GetUserById).Methods("GET")

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
