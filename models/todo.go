package models

import (
	"time"
)

type Todo struct {
	TodoID    int       `db:"todo_id"`
	Title     string    `db:"title"`
	UserId    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsDeleted bool      `db:"is_deleted"`
}
