package main

import (
	"fmt"
	"net/http"

	"backend/database"
	"backend/handlers"
	"backend/middleware"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	database.InitDatabase()

	http.HandleFunc("/", middleware.WithCORS(WrongHandler))
	http.HandleFunc("/api/login", middleware.WithCORS(handlers.LoginHandler))
	http.HandleFunc("/api/register", middleware.WithCORS(handlers.RegisterHandler))

	fmt.Println("Servern körs på http://localhost:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func WrongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WrongHandler handled it")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Hej från backend! något gick fel")
}