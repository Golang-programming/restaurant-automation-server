package dto

type CreateTableInput struct {
	Number int `json:"number" validate:"required"`
}

type UpdateTableInput struct {
	Number int `json:"number" validate:"omitempty"`
}
