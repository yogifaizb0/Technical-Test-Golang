package main

import (
	"log"
	"net/http"
	"simplefileupload/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", controllers.UploadController).Methods("POST")
	r.HandleFunc("/files", controllers.GetAllUploadedFilesController).Methods("GET")
	r.HandleFunc("/files/{filename}", controllers.DownloadFileController).Methods("GET")

	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}