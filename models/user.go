package models

import (
	"time"
)

type User struct {
	UserID      int       `db:"user_id"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	IsDeleted   bool      `db:"is_deleted"`
}
