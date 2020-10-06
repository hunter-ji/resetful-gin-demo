package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int    `gorm:"autoIncrement;primaryKey"`
	Title     string `gorm:"type:varchar(30)"`
	UserId    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
