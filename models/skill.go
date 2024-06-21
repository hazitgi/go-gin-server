package models

import (
	"time"

	"gorm.io/gorm"
)
type SkillGroup struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	Skills    []Skill        `gorm:"foreignKey:SkillGroupID" json:"skills"`
}

type Skill struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	SkillGroupID int           `json:"-"`
	SkillGroup   SkillGroup     `json:"skillGroup"`
}

func NewSkill() *Skill {
	return &Skill{}
}

func NewSkillGroup() *SkillGroup {
	return &SkillGroup{}
}
