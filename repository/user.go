package repository

import (
	"database/sql"
	"errors"
	"ks/domain"
	"log"
)

var err error

type Repository struct {
	db *sql.DB
}

var repo Repository

func init() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	repo.db = db
}
func GetByName(user domain.SignUp) error {
	query := "SELECT COUNT(*) FROM sign_up WHERE user_name=$1"
	var count int
	err := repo.db.QueryRow(query, user.UserName).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	return nil
}

func CreateUser(user domain.SignUp) error {
	query := "INSERT INTO sign_up (user_name, password) VALUES ($1, $2)"
	_, err = repo.db.Exec(query, user.UserName, user.Password)
	if err != nil {
		return err
	} else {
		return nil
	}

}
func Login(user domain.SignUp) error {
	query := "SELECT COUNT(*) FROM sign_up WHERE user_name = $1 AND password = $2"
	var count int
	err = repo.db.QueryRow(query, user.UserName, user.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("invalid username or password")
	}

	return nil
}

func DeleteUser(user domain.SignUp) error {
	query := "DELETE FROM sign_up WHERE user_name = $1"
	_, err := repo.db.Exec(query, user.UserName)
	if err != nil {
		return err
	}
	return nil
}
