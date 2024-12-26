package main

import (
	"log"
	"net/http"
	"userauthentication/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.RegisterController).Methods("POST")
	r.HandleFunc("/login", controllers.LoginController).Methods("POST")
	r.HandleFunc("/profile", controllers.ProfileController).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}