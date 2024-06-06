package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go/websrv3/model"
)

var users model.UsersList

func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	offset := r.PathValue("offset")
	uoffset, err := strconv.ParseUint(offset, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	allUsers, err := users.ListAll(uint(uoffset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

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

	w.WriteHeader(http.StatusFound)
	err = json.NewEncoder(w).Encode(allUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userTemp model.User
	err := json.NewDecoder(r.Body).Decode(&userTemp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := users.Create(userTemp.Name, userTemp.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
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
	id := r.PathValue("id")

	msg, err := users.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotModified)
		return
	}

	response := Response{
		Msg: msg,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
