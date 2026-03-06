package dto

// TaskCreateDTO is used to validate the request body when creating a new task.
type TaskCreateDTO struct {
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Description string `json:"description" validate:"max=1000"`
	Status      string `json:"status" validate:"required,oneof=pending in_progress completed"`
}

// TaskUpdateDTO is used to validate the request body when updating an existing task.
type TaskUpdateDTO struct {
	Title       string `json:"title" validate:"omitempty,min=1,max=200"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	Status      string `json:"status" validate:"omitempty,oneof=pending in_progress completed"`
}
