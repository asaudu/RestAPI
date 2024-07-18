package models

import "database/sql"

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

	result, err := statement.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}
