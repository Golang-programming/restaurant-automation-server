package dto

type CreateMenuInput struct {
	Name        string `json:"name" binding:"required" validate:"required,min=2,max=255"`
	Description string `json:"description" validate:"omitempty"`
	Status      string `json:"status" binding:"required" validate:"required,oneof='active' 'inactive'"`
}

type UpdateMenuInput struct {
	Name        string `json:"name" validate:"omitempty,min=2,max=255"`
	Description string `json:"description" validate:"omitempty"`
	Status      string `json:"status" validate:"omitempty,oneof='active' 'inactive'"`
}

type AddFoodToMenuInput struct {
	FoodID uint `json:"food_id" binding:"required" validate:"required"`
}

type RemoveFoodFromMenuInput struct {
	FoodID uint `json:"food_id" binding:"required" validate:"required"`
}

type AddDealToMenuInput struct {
	DealID uint `json:"deal_id" binding:"required" validate:"required"`
}

type RemoveDealFromMenuInput struct {
	DealID uint `json:"deal_id" binding:"required" validate:"required"`
}
