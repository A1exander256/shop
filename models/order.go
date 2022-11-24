package models

type Order struct {
	Uuid   int `json:"id" db:"uuid"`
	UserId int `json:"user_id" db:"user_id"`
}
