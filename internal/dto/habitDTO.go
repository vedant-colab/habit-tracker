package dto

type CreateHabitRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required,min=3"`
}

type UpdateHabitRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
