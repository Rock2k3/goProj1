package handlers

import (
	"app2/repository"
	"encoding/json"
	"log"
	"net/http"
)

func AddHandlers()  {
	http.HandleFunc("/users", getAllUsers)
}

func getAllUsers(w http.ResponseWriter, r *http.Request)   {
	log.Println("getAllUsers")

	users, err := repository.GetUsers()
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}