package server

import (
	"fmt"
    "net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating User")
}