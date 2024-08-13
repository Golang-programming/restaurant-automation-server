package dto

type UpdateCustomerInput struct {
	TotalGuests int    `json:"total_guests" validate:"required,min=1"`
	Name        string `json:"name" validate:"omitempty,min=2,max=100"`
}

type CreateCustomerInput struct {
	TotalGuests int    `json:"total_guests" binding:"required" validate:"required,min=1"`
	Name        string `json:"name" validate:"omitempty,min=2,max=100"`
}
