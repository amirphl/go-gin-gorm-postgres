package entity

// Book model
type Book struct {
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title  string `gorm:"type:varchar(255)"	json:"title"`
	Desc   string `gorm:"type:text"	json:"description"`
	UserID uint64 `gorm:"not null"	json:"-"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"user"`
}
