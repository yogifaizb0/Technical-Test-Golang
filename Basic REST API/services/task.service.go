package services

import (
	"basicrestapi/dto"
	"basicrestapi/models"
)

var (
	tasks = []models.Task{}
	id = 1
)
func CreateTask(taskDto dto.TaskCreate) models.Task {
	newTask := models.Task{
		ID: id,
		Title: taskDto.Title,
		Description: taskDto.Description,
	}
	id++
	tasks = append(tasks, newTask)
	return newTask
}

func GetAllTasks() []models.Task {
	return tasks
}

func DeleteTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]... )
			return true
		}
	}
	return false
}