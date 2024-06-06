package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go/websrv2/model"
)

var users model.UsersList

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)

	allUsers := users.ListAll()
	if len(allUsers) == 0 {
		response := Response{
			Msg: "users not found",
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	err := json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	var userTemp model.User
	err := json.NewDecoder(r.Body).Decode(&userTemp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := users.Create(userTemp.Name, userTemp.Age)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	user, err := users.Read(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	var userTemp model.User

	id := r.PathValue("id")
	err := json.NewDecoder(r.Body).Decode(&userTemp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := users.Update(id, userTemp.Name, userTemp.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotModified)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	err := users.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotModified)
		return
	}

	msg := fmt.Sprintf("user %v deleted", id)
	response := Response{
		Msg: msg,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
