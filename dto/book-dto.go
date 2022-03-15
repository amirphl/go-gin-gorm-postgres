package dto

// BookUpdateDTO PUT model:book
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required" validate:"min=1,max=255"`
	Description string `json:"description" form:"description" binding:"required" validate:"min=1,max=1023"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty" binding:"required"`
}

// BookCreateDTO POST model:book
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required" validate:"min=1,max=255"`
	Description string `json:"description" form:"description" binding:"required" validate:"min=1,max=1023"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty" binding:"required"`
}
