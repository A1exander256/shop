package storage

import (
	"fmt"
	"strings"

	"github.com/alexander256/shop/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserStorage struct {
	log *logrus.Logger
	db  *sqlx.DB
}

func NewUserStorage(db *sqlx.DB, log *logrus.Logger) *UserStorage {
	return &UserStorage{
		log: log,
		db:  db,
	}
}

func (s *UserStorage) Create(user *models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstname, surname, middlename, sex, age) values ($1, $2, $3, $4, $5) RETURNING id", tableUsers)

	row := s.db.QueryRow(query, user.Firstname, user.Surname, user.Middlename, user.Sex, user.Age)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (s *UserStorage) Update(user *models.User) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Firstname != "" {
		setValue = append(setValue, fmt.Sprintf("firstname=$%d", argId))
		args = append(args, user.Firstname)
		argId++
	}

	if user.Surname != "" {
		setValue = append(setValue, fmt.Sprintf("surname=$%d", argId))
		args = append(args, user.Surname)
		argId++
	}

	if user.Middlename != "" {
		setValue = append(setValue, fmt.Sprintf("middlename=$%d", argId))
		args = append(args, user.Middlename)
		argId++
	}
	if user.Sex != "" {
		setValue = append(setValue, fmt.Sprintf("sex=$%d", argId))
		args = append(args, user.Sex)
		argId++
	}

	if user.Age != 0 {
		setValue = append(setValue, fmt.Sprintf("age=$%d", argId))
		args = append(args, user.Age)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE %s users SET %s WHERE users.id = $%d", tableUsers, setQuery, argId)

	args = append(args, user.Id)

	_, err := s.db.Exec(query, args...)
	return err
}

func (s *UserStorage) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s users WHERE users.id=$1", tableUsers)
	_, err := s.db.Exec(query, userId)
	return err
}

func (s *UserStorage) GetById(userId int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s users WHERE users.id=$1", tableUsers)
	err := s.db.Get(&user, query, userId)
	return user, err
}

func (s *UserStorage) GetAll() ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf("SELECT * FROM %s users ORDER BY users.id", tableUsers)
	err := s.db.Select(&users, query)
	return users, err
}
