package database

// User ...
type User struct {
	Email    string
	Password string
}

func (db *Database) createAdminUser() error {

	// if cfg.Admin.Email == "" || cfg.Admin.Password == "" {
	// 	return nil
	// }

	// if db.rowExists("SELECT id FROM users WHERE email = ?", cfg.Admin.Email) {
	// 	var query = `UPDATE users SET password = ? WHERE email = ?;`
	// 	_, err := db.Exec(query, cfg.Admin.Password, cfg.Admin.Email)
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	// 	var query = `INSERT INTO users (email,password) VALUES(?, ?);`
	// 	_, err := db.Exec(query, cfg.Admin.Email, cfg.Admin.Password)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
