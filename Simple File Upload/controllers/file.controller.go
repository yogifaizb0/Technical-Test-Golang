package controllers

import (
	"net/http"
	"simplefileupload/services"
	"simplefileupload/utils"

	"github.com/gorilla/mux"
)

func UploadController(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.Response(w, false, err.Error(), http.StatusBadRequest, nil)
		return
	}

	if err := services.UploadFile(file, header); err != nil {
		utils.Response(w, false, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	utils.Response(w, true, "Success upload file", http.StatusOK, header.Filename)
}

func GetAllUploadedFilesController(w http.ResponseWriter, r *http.Request) {
	files, err := services.GetAllUploadedFiles()
	if err != nil {
		utils.Response(w, false, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	utils.Response(w, true, "Success get all files", http.StatusOK, files)
}

func DownloadFileController(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]
	filepath, err := services.CheckFile(filename)
	if err != nil {
		utils.Response(w, false, "File not found", http.StatusNotFound, nil)
		return
	}
	http.ServeFile(w, r, filepath)
	utils.Response(w, true, "Success download file", http.StatusOK, nil)
}