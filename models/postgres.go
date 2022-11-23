package models

type ConfigPostgres struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}
