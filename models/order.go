package models

type Order struct {
	Id     int `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
}
