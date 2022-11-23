package models

type User struct {
	Uuid       int    `json:"-" db:"uuid"`
	Firstname  string `json:"firstname" db:"firstname"`
	Surname    string `json:"surname" db:"surname"`
	Middlename string `json:"middlename" db:"middlename"`
	FIO        string `json:"-" db:"fio"`
	Sex        string `json:"sex" db:"sex" binding:"oneof=man woman"`
	Age        uint   `json:"age" db:"age" binding:"min=0,max=100"`
}
