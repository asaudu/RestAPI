package models

import (
	"database/sql"
	"errors"

	"addyCodes.com/RestAPI/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save(db *sql.DB) error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	statement, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	result, err := statement.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u User) ValidateCredentials(db *sql.DB) error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
