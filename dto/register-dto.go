package dto

// RegisterDTO POST /register
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min=1,max=255"`
	Email    string `json:"email" form:"email" binding:"required" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$,max=255"`
	Password string `json:"password" form:"password" binding:"required" validate:"min=6,max=255"`
}
