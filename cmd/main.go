package main

import (
  "net/http"

  "github.com/brunoMirand/rebellious-users_api/internal/factory"
  "github.com/brunoMirand/rebellious-users_api/internal/handlers"
  "github.com/gorilla/mux"
)

func main() {
	userFactory := factory.NewUserFactory()
  userHandler := handlers.NewUserHandler(userFactory)

  r := mux.NewRouter()
  r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
  r.HandleFunc("/users", userHandler.FindAllUsers).Methods("GET")

  http.ListenAndServe(":8080", r)
}
