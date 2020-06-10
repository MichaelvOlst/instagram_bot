package database

import (
	"fmt"
	"instagram_bot/config"
)

type User struct {
	Email    string
	Password string
}

func (db *Database) createAdminUser(cfg *config.Config) error {

	if db.rowExists("SELECT id FROM users WHERE email = ?", cfg.Admin.Email) {
		var query = `UPDATE users SET password = ? WHERE email = ?;`
		fmt.Println(query)

		_, err := db.Exec(query, cfg.Admin.Password, cfg.Admin.Email)
		if err != nil {
			return err
		}
	} else {
		var query = `INSERT INTO users (email,password) VALUES(?, ?);`
		fmt.Println(query)
		_, err := db.Exec(query, cfg.Admin.Email, cfg.Admin.Password)
		if err != nil {
			return err
		}
	}

	return nil

	// rows, err := db.Query("SELECT id FROM users WHERE email = ?", cfg.Admin.Email)
	// if err != nil && err != sql.ErrNoRows {
	// 	return err
	// }

	// u := &User{}
	// rows.Scan(u)

	// fmt.Printf("%+v\n", u)
	// fmt.Printf("%+v\n", err)

	// if err == sql.ErrNoRows {

	// }

	return nil
}
