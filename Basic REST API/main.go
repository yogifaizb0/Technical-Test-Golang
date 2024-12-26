package main

import (
	"basicrestapi/controllers"
	"basicrestapi/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllTasksController(w, r)
		case http.MethodPost:
			controllers.CreateTaskController(w, r)
		default: 
			utils.Response(w, false, "Method not allowed", http.StatusMethodNotAllowed, nil)
		}
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			controllers.DeleteTaskController(w, r)
		default: 
			utils.Response(w, false, "Method not allowed", http.StatusMethodNotAllowed, nil)
	}})

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}