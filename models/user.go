package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"addyCodes.com/RestAPI/metrics"
	"addyCodes.com/RestAPI/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save(db *sql.DB) error {
	start := time.Now()
	defer func() {
		metrics.DbQueryDuration.WithLabelValues("GetAllEvents").Observe(time.Since(start).Seconds())
	}()

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
	start := time.Now()
	defer func() {
		metrics.DbQueryDuration.WithLabelValues("ValidateUserCredentials").Observe(time.Since(start).Seconds())
	}()

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.QueryRow(query, u.Email)
	fmt.Println("ValidateCredentials row check ", row)

	var retrievedPassword string
	//var theId int64
	err := row.Scan(&u.ID, &retrievedPassword)
	fmt.Println("ID Check", u.ID)
	fmt.Println("ValidateCredentials error check ", err)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	fmt.Println("ValidateCredentials passwordIsValid check ", passwordIsValid)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
