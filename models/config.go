package models

type Config struct {
	Logger   `json:"logger"`
	Gin      `json:"gin"`
	Postgres `json:"postgres"`
}
