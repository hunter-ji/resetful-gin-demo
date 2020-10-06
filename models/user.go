package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int    `gorm:"autoIncrement;primaryKey"`
	Username    string `gorm:"type:varchar(10)"`
	Password    string `gorm:"type:varchar(30)"`
	PhoneNumber string `gorm:"type:varchar(12)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
