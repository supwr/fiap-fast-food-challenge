package dto

type Item struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Type        string  `json:"type" validate:"required"`
	Price       float64 `json:"price" required:"required"`
}
