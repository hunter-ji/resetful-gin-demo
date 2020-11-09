package models

type Todo struct {
	TodoID int    `db:"todo_id"`
	Title  string `db:"title"`
	UserId int    `db:"user_id"`
	ModelTime
}
