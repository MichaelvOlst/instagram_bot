package database

import "instagram_bot/models"

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
