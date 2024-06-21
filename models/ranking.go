package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSkillRank struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int
	User      User
	SkillID   int
	Skill     Skill
	Rank      int `json:"rank"`
}

type Competence struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"-"`
	User      User           `json:"-"`
	SkillID   uint           `json:"-"`
	Skill     Skill          `json:"skill"`
	Rank      int            `json:"rank"`
}

func NewUserSkillRank() *UserSkillRank {
	return &UserSkillRank{}
}

func NewCompetence() *Competence {
	return &Competence{}
}
