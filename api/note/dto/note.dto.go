package dto

type CreateNoteInput struct {
	Description string `json:"description" binding:"required"`
	OrderItemID uint   `json:"order_item_id" binding:"required"`
}

type UpdateNoteInput struct {
	Description string `json:"description" binding:"required"`
	OrderItemID uint   `json:"order_item_id" binding:"required"`
}
