package models

type Config struct {
	Logger   `json:"logger"`
	Gin      `json:"gin"`
	Postgres `json:"postgres"`
}

type Gin struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type Logger struct {
	Level string `json:"level"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	SSLMode  string `json:"ssl_mode"`
}
