package database

import (
	"database/sql"
	"errors"
	"instagram_bot/models"
)

// CreateUser creates the user in the database
func (db *Database) CreateUser(u *models.User) error {

	var query = `INSERT INTO users (email,password) VALUES(?, ?);`
	_, err := db.Exec(query, u.Email, u.Password)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByEmail deletes an user by Email
func (db *Database) DeleteUserByEmail(u *models.User) error {

	var query = `DELETE FROM users WHERE email = ?;`
	_, err := db.Exec(query, u.Email)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID get a user by ID
func (db *Database) GetUserByID(ID int64) (*models.User, error) {

	u := &models.User{}
	query := db.Rebind("SELECT * FROM users WHERE id = ?")
	err := db.Get(u, query, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}

		return nil, err
	}

	return u, nil
}
