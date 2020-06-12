package database

import (
	"database/sql"
	"errors"
	"instagram_bot/models"
)

// CheckToken checks if a give token is valid
func (db *Database) CheckToken(token string) (*models.User, error) {

	t := &models.Token{}
	query := db.Rebind("SELECT * FROM tokens WHERE token = ?")
	err := db.Get(t, query, token)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Api token is invalid")
		}

		return nil, err
	}

	u := &models.User{}
	query = db.Rebind("SELECT * FROM users WHERE id = ? LIMIT 1")
	err = db.Get(u, query, t.UserID)

	return u, err
}
