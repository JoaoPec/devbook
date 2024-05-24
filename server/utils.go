package server

import (
	"devbook/database"
	"devbook/models"
	"encoding/json"
	"io"
	"net/http"
    mux "github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	req, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request"))
		return
	}

	var user models.User

	err = json.Unmarshal(req, &user)
	if err != nil {
		w.Write([]byte("Error parsing request"))
		return
	}

	database.InsertUser(user)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)

    id := vars["id"]

    err := database.SearchUserById(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User not found"))
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User found successfully"))

}
