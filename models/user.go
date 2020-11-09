package models

type User struct {
	UserID      int    `db:"user_id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	ModelTime
}
