package dto

import "github.co/golang-programming/restaurant/api/entity"

type CreateTableInput struct {
	Number int                `json:"number" validate:"required"`
	Status entity.TableStatus `json:"status" validate:"omitempty,oneof=observed available"`
}

type UpdateTableInput struct {
	Number int                `json:"number" validate:"omitempty"`
	Status entity.TableStatus `json:"status" validate:"omitempty,oneof=observed available"`
}
