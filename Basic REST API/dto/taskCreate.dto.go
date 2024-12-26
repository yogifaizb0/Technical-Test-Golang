package dto

type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}