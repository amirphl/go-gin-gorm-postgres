package dto

// LoginDTO POST /login
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$,max=255"`
	Password string `json:"password" form:"password" binding:"required" validate:"min=6,max=255"`
}
