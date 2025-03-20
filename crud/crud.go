package main

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/crud/handlers"
	"github.com/AgustinCardozo/labo-so/utils/greetings"
	"github.com/AgustinCardozo/labo-so/utils/web/server"
	"net/http"
)

const Puerto = 8080

func main() {
	fmt.Println(greetings.Hello("Test"))

	http.HandleFunc("POST /users", handlers.CreateHandler())
	http.HandleFunc("DELETE /users/{id}", handlers.DeleteHandler())
	http.HandleFunc("GET /users", handlers.GetAllHandler())
	http.HandleFunc("GET /users/{id}", handlers.GetByIdHandler())
	http.HandleFunc("PUT /users/{id}", handlers.UpdateHandler())

	err := server.InitServer(Puerto)
	if err != nil {
		return
	}
}
