package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email" gorm:"uniqueIndex"`
}
