package main

import (
	"fmt"
	"log"
	"net/http"
    "devbook/database"
    "devbook/server"
	mux "github.com/gorilla/mux"
)


func main(){

    router := mux.NewRouter()

    database.StartDB()

    router.HandleFunc("/users", server.CreateUser).Methods("POST")


    fmt.Println("Server is running on port 3000")
    log.Fatal(http.ListenAndServe(":3000", router))
}
