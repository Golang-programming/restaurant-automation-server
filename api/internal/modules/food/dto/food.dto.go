package dto

import "github.co/golang-programming/restaurant/api/internal/entity"

type CreateFoodInput struct {
	Name         string            `json:"name" binding:"required"`
	Price        float64           `json:"price" binding:"required"`
	Description  string            `json:"description" binding:"required"`
	Images       []string          `json:"images"`
	Status       entity.FoodStatus `json:"status" binding:"required,oneof=active inactive"`
	ParentFoodID *uint             `json:"parent_food_id"`
	MenuID       uint              `json:"menu_id" binding:"required"`
	Category     string            `json:"category" binding:"required"`
}

type UpdateFoodInput struct {
	Name         string            `json:"name" binding:"omitempty"`
	Price        float64           `json:"price" binding:"omitempty"`
	Description  string            `json:"description" binding:"omitempty"`
	Images       []string          `json:"images"`
	Status       entity.FoodStatus `json:"status" binding:"omitempty,oneof=active inactive"`
	ParentFoodID *uint             `json:"parent_food_id"`
	MenuID       uint              `json:"menu_id" binding:"omitempty"`
	Category     string            `json:"category" binding:"omitempty"`
}

type ChangeFoodStatusInput struct {
	Status entity.FoodStatus `json:"status" binding:"required,oneof=active inactive"`
}
