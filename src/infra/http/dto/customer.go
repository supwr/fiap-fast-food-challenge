package dto

type Customer struct {
	Name     string `json:"name" validate:"required"`
	Document string `json:"document" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
