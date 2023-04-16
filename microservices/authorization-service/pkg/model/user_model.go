package model

type User struct {
	ID       uint   `gorm:"primaryKey" json:"ID"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      uint8  `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
	Salt     []byte `json:"salt"`
}
