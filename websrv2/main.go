package main

import (
	"log"
	"net/http"

	"go/websrv2/routes"
)

func main() {
	// mux for http server
	mux := http.NewServeMux()

	// root/home route
	// curl -X GET http://localhost:8080/ -H "Accept: application/json" | jq 
	mux.HandleFunc("GET /", routes.Home)

	// CRUD routes for UserList struct
	// curl -X GET http://localhost:8080/users -H "Accept: application/json" | jq
	mux.HandleFunc("GET /users", routes.ListAllUsers)

	// curl -X POST http://localhost:8080/user -H "Content-Type: application/json" -d '{"name": "John Doe", "age": 30}' | jq
	mux.HandleFunc("POST /user", routes.CreateUser)

	// curl -X GET http://localhost:8080/user/<id> -H "Accept: application/json" | jq
	mux.HandleFunc("GET /user/{id}", routes.GetUserById)

	// curl -X PUT http://localhost:8080/user/<id> -H "Content-Type: application/json" -d '{"name": "John Doe", "age": 26}' | jq
	mux.HandleFunc("PUT /user/{id}", routes.UpdateUserById)

	// curl -X DELETE http://localhost:8080/user/<id> -H "Accept: application/json" | jq
	mux.HandleFunc("DELETE /user/{id}", routes.DeleteUserById)

	// http server instance
	log.Println("websrv2 running on 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("server error %v", err)
	}
}
