package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	FullName   string         `json:"fullName"`
	Email      string         `json:"email"`
	Competence []Competence   `gorm:"foreignKey:UserID" json:"competence"`
}

func NewUser() *User {
	return &User{}
}
