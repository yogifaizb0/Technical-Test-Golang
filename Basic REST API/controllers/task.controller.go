package controllers

import (
	"basicrestapi/dto"
	"basicrestapi/services"
	"basicrestapi/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateTaskController(w http.ResponseWriter, r *http.Request) {
	var taskDto dto.TaskCreate
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		utils.Response(w, false, "Bad Request", http.StatusBadRequest, nil)
		return
	}
	newTask := services.CreateTask(taskDto)
	utils.Response(w, true, "Success create new tasks", http.StatusCreated, newTask)
}

func GetAllTasksController(w http.ResponseWriter, r *http.Request){
	tasks := services.GetAllTasks()
	utils.Response(w, true, "Success", http.StatusOK, tasks)
}

func DeleteTaskController(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, false, "Invalid ID", http.StatusBadRequest, nil)
		return
	}

	if services.DeleteTask(id) {
		utils.Response(w, true, "Success Delete", http.StatusOK, nil)
	} else {
		utils.Response(w, false, "Task not found", http.StatusNotFound, nil)
	}
}