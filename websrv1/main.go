package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Response struct definition
type Response struct {
	Msg string `json:"msg"`
}

// User struct definition
type User struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

// UserList struct definition and methods for CRUD operations
type UsersList struct {
	Users []User
}

func (u *UsersList) listAll() []User {
	return u.Users
}

func (u *UsersList) create(name string, age int) User {
	user := User{
		Id:   uuid.New(),
		Name: name,
		Age:  age,
	}

	u.Users = append(u.Users, user)
	return user
}

func (u *UsersList) read(id string) (User, error) {
	for _, user := range u.Users {
		if user.Id.String() == id {
			return user, nil
		}
	}

	err := errors.New("user not found to read")
	return User{}, err
}

func (u *UsersList) update(id string, name string, age int) (User, error) {
	for i, user := range u.Users {
		if user.Id.String() == id {
			u.Users[i].Name = name
			u.Users[i].Age = age
			return u.Users[i], nil
		}
	}

	err := errors.New("user not found to update")
	return User{}, err
}

func (u *UsersList) delete(id string) error {
	for i, user := range u.Users {
		if user.Id.String() == id {
			u.Users = append(u.Users[:i], u.Users[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found to update")
}

func main() {
	// mux for http server
	mux := http.NewServeMux()

	// routes and handle functions
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: [GET] - Route: [/] - Remote Address: [%v]", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("server running ..."))
	})

	// CRUD routes for UserList struct ----------------------------------------
	var users UsersList

	// curl -X GET http://localhost:8080/users -H "Accept: application/json" | jq
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: [GET] - Route: [/users] - Remote Address: [%v]", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)

		allUsers := users.listAll()
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
	})

	// curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name": "Clark Kent", "age": 30}' | jq
	mux.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: [POST] - Route: [/user] - Remote Address: [%v]", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")

		id := uuid.New()
		var userTemp User
		err := json.NewDecoder(r.Body).Decode(&userTemp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userTemp.Id = id

		response := users.create(userTemp.Name, userTemp.Age)
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// curl -X GET http://localhost:8080/user/<id> -H "Accept: application/json" | jq
	mux.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: [%v] - Route: [%v] - Remote Address: [%v]", r.Method, r.RequestURI, r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")
		user, err := users.read(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// curl -X PUT http://localhost:8080/user/<id> -H "Content-Type: application/json" -d '{"name": "John Doe", "age": 26}' | jq
	mux.HandleFunc("PUT /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: [PUT] - Route: [/user/{id}] - Remote Address: [%v]", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")

		var userTemp User

		id := r.PathValue("id")
		err := json.NewDecoder(r.Body).Decode(&userTemp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := users.update(id, userTemp.Name, userTemp.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotModified)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// curl -X DELETE http://localhost:8080/user/<id> -H "Accept: application/json" | jq
	mux.HandleFunc("DELETE /user/{id}", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Method: [DELETE] - Route: [/user/{id}] - Remote Address: [%v]", r.RemoteAddr)
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")

		err := users.delete(id)
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
	})

	// http server instance
	log.Println("websrv1 running on 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("server error %v", err)
	}
}
