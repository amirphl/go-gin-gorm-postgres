package dto

// BookUpdateDTO PUT model:book
type BookUpdateDTO struct {
	ID     uint64 `json:"-" form:"-"`
	Title  string `json:"title" form:"title" binding:"required" validate:"min=1,max=255"`
	Desc   string `json:"description" form:"description" binding:"required" validate:"min=1,max=1023"`
	UserID uint64 `json:"-" form:"-"`
}

// BookCreateDTO POST model:book
type BookCreateDTO struct {
	Title  string `json:"title" form:"title" binding:"required" validate:"min=1,max=255"`
	Desc   string `json:"description" form:"description" binding:"required" validate:"min=1,max=1023"`
	UserID uint64 `json:"-" form:"-"`
}
